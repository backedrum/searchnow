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
	"github.com/go-resty/resty"
	"strconv"
)

type googleResult struct {
	Items []*googleResultItem `json:"items"`
}

type googleResultItem struct {
	Title   string `json:"title"`
	Url     string `json:"link"`
	Snippet string `json:"snippet"`
}

// searchViaGoogleApi performs search via Google Custom Search API.
func searchViaGoogleApi(searchTerm string, numOfResults int) []*SearchResult {
	result := make([]*SearchResult, 0, numOfResults)

	apiKey := tokens["google.apikey"]
	engine := tokens["google.engine"]

	resp, err := resty.R().SetQueryParams(map[string]string{
		"key":         apiKey,
		"cx":          engine,
		"prettyPrint": "true",
		"num":         strconv.Itoa(numOfResults),
		"q":           searchTerm,
	}).SetHeader("Accept", "application/json").
		Get("https://www.googleapis.com/customsearch/v1")

	if err != nil {
		println("Problem has been occurred during search:" + err.Error())
		return result
	}

	googleResult := googleResult{}
	err = json.Unmarshal(resp.Body(), &googleResult)
	if err != nil {
		println("Problem has been occurred during unmarshalling:" + err.Error())
		return result
	}

	for _, item := range googleResult.Items {
		sr := SearchResult{}
		sr.Title = item.Title
		sr.Url = item.Url
		sr.Contents = item.Snippet

		result = append(result, &sr)
	}

	return result
}
