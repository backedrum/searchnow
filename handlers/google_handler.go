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
