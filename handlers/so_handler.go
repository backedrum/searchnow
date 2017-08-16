package handlers

import (
	"fmt"
	"github.com/laktek/Stack-on-Go/stackongo"
)

/*
  Performs StackOverflow.com search via Stack-On-Go library.
*/
func searchStackOverflow(searchTerm string, numOfResults int) []*SearchResult {
	result := make([]*SearchResult, 0, numOfResults)

	session := stackongo.NewSession("stackoverflow")
	questions, err := session.Search(searchTerm, stackongo.Params{})
	if err != nil {
		fmt.Printf(err.Error())
		return result
	}

	for i, question := range questions.Items {
		if i == numOfResults {
			break
		}

		sr := SearchResult{}
		sr.Url = question.Link
		sr.Title = question.Title
		result = append(result, &sr)
	}

	return result
}
