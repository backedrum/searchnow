/*
Copyright 2017 Andrii Zablodskyi (andrey.zablodskiy@gmail.com)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty"
	"strings"
)

const (
	MAX_RESULTS  = 25
	NEW_STORIES  = "newstories"
	TOP_STORIES  = "topstories"
	BEST_STORIES = "beststories"
	ASK_STORIES  = "askstories"
	JOB_STORIES  = "jobstories"
)

type storyContent struct {
	Url   string `json:"url"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

//fetchHackerNews fetches Hacker News stories data accordingly to the specified term.
func fetchHackerNews(term string, numOfResults int) []*SearchResult {

	if numOfResults > MAX_RESULTS {
		fmt.Printf("Max number of results to fetch is %d. Only first %d stories will be shown.", MAX_RESULTS, MAX_RESULTS)
		numOfResults = MAX_RESULTS
	}

	result := make([]*SearchResult, 0, numOfResults)

	switch term {
	case NEW_STORIES, TOP_STORIES, BEST_STORIES, ASK_STORIES, JOB_STORIES:
		for _, id := range fetchStoriesIds(term, numOfResults) {
			story := fetchStoryContent(strings.Replace(id, " ", "", -1))
			result = append(result, transformToSearchResult(story))
		}
		break
	default:
		fmt.Printf("Unsupported search term %s\n", term)
	}

	return result
}

func fetchStoriesIds(term string, maxCount int) []string {
	resp, err := resty.R().SetQueryParams(map[string]string{
		"print": "pretty",
	}).SetHeader("Accept", "application/json").
		Get("https://hacker-news.firebaseio.com/v0/" + term + ".json")

	if err != nil {
		println("Problem has been occurred during fetching stories ids list:" + err.Error())
		return make([]string, 0)
	}

	responseArray := strings.Replace(resp.String(), "[", "", 2)

	result := strings.Split(responseArray, ",")

	if len(result) > maxCount {
		return result[:maxCount]
	}

	return result
}

func fetchStoryContent(id string) *storyContent {
	storyContent := storyContent{}

	resp, err := resty.R().SetQueryParams(map[string]string{
		"print": "pretty",
	}).SetHeader("Accept", "application/json").
		Get("https://hacker-news.firebaseio.com/v0/item/" + id + ".json")

	if err != nil {
		println("Problem has been occurred during fetching of a story content:" + err.Error())
		return &storyContent
	}

	err = json.Unmarshal(resp.Body(), &storyContent)
	if err != nil {
		println("Problem has been occurred during unmarshalling of a story content:" + err.Error())
	}

	return &storyContent
}

func transformToSearchResult(storyContent *storyContent) *SearchResult {
	result := SearchResult{}
	result.Title = storyContent.Title
	result.Url = storyContent.Url
	result.Contents = storyContent.Text

	return &result
}
