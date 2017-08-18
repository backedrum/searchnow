package handlers

import (
	"fmt"
	"github.com/laktek/Stack-on-Go/stackongo"
	"strconv"
)

const MAX_ANSWERS = 4

//searchStackOverflow performs StackOverflow.com search via Stack-On-Go library.
func searchStackOverflow(searchTerm string, numOfResults int) []*SearchResult {
	result := make([]*SearchResult, 0, numOfResults)

	params := make(stackongo.Params)
	params.Add("filter", "withbody")

	session := stackongo.NewSession("stackoverflow")
	questions, err := session.Search(searchTerm, params)
	if err != nil {
		fmt.Printf("Error while retrieving questions:%s\n", err.Error())
		return result
	}

	for i, question := range questions.Items {
		if i == numOfResults {
			break
		}

		sr := SearchResult{}
		sr.Url = question.Link
		sr.Title = question.Title
		sr.Contents = question.Body

		if question.Answer_count > 0 {

			sr.Extras = make(map[string]string)
			sr.ExtrasOrder = make([]string, 0, MAX_ANSWERS)

			// sort answers by score
			params := make(stackongo.Params)
			params.Add("filter", "withbody")
			params.Sort("votes")

			// is answered?
			answersCount := 0
			if question.Is_answered && question.Accepted_answer_id > 0 {
				answer, err := session.GetAnswers([]int{question.Accepted_answer_id}, params)
				if err != nil {
					fmt.Printf("Cannot retrieve a correct answer:%s", err.Error())
				} else {
					vote := "(vote:" + strconv.Itoa(answer.Items[0].Score) + ") "
					extra := "Answer 1(\u2713):"
					sr.Extras[extra] = vote + answer.Items[0].Body
					sr.ExtrasOrder = append(sr.ExtrasOrder, extra)

					answersCount++
				}
			}

			answers, _ := session.AnswersForQuestions([]int{question.Question_id}, params)
			for _, answer := range answers.Items {
				if answersCount > MAX_ANSWERS {
					break
				}

				if answer.Is_accepted {
					continue
				}

				vote := "(vote:" + strconv.Itoa(answer.Score) + ") "
				extra := "Answer " + strconv.Itoa(answersCount+1) + ":"
				sr.Extras["Answer "+strconv.Itoa(answersCount+1)+":"] = vote + answer.Body
				sr.ExtrasOrder = append(sr.ExtrasOrder, extra)

				answersCount++
			}
		}

		result = append(result, &sr)
	}

	return result
}
