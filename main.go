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
package main

import (
	"encoding/json"
	"fmt"
	"github.com/backedrum/searchnow/display"
	"github.com/backedrum/searchnow/handlers"
	tm "github.com/buger/goterm"
	"github.com/manifoldco/promptui"
	"io/ioutil"
	"os"
	"strconv"
)

type Config struct {
	ShowTitle    bool `json:"title"`
	ShowURL      bool `json:"url"`
	ShowContents bool `json:"contents"`
	ShowOthers   bool `json:"others"`
}

func main() {
	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Println("Usage: searchnow <searh term> <optional engine> <optional max number of results>")
		os.Exit(1)
	}

	searchTerm := os.Args[1]

	engine := "google"
	if len(os.Args) > 2 {
		engine = os.Args[2]
	}

	numOfResults := 5
	if len(os.Args) == 4 {
		numOfResults, _ = strconv.Atoi(os.Args[3])
	}

	if !handlers.HasEngineSupport(engine) {
		fmt.Printf("Sorry, but engine %s is not supported.\n", engine)
		os.Exit(1)
	}

	// init config
	bytes, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Printf("Cannot read application config. Error:%s", err.Error())
		os.Exit(1)
	}

	config := Config{}
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		fmt.Printf("Cannot init application config. Error:%s", err.Error())
		os.Exit(1)
	}

	results := handlers.Search(engine, searchTerm, numOfResults)

	titles := []string{}
	for _, result := range results {
		titles = append(titles, result.Title)
	}

	console := promptui.Select{
		Label: "Select result",
		Items: titles,
	}

	index, _, err := console.Run()
	if err != nil {
		fmt.Printf("Cannot display results. Error:%s", err.Error())
		os.Exit(1)
	}

	result := results[index]

	display.PutLine("URL:", result.Url, tm.RED, tm.BLUE, config.ShowURL && result.Url != "")
	display.PutLine("Title:", result.Title, tm.RED, -1, config.ShowTitle && result.Title != "")
	display.PutLine("Snippet:", display.ConvertHtmlToText(result.Contents), tm.RED, tm.GREEN, config.ShowContents && result.Contents != "")

	if config.ShowOthers && len(result.Extras) > 0 {
		for _, extra := range result.ExtrasOrder {
			display.PutLine(extra, display.ConvertHtmlToText(result.Extras[extra]), tm.RED, tm.YELLOW, true)
		}
	}
	tm.Flush()

	return
}
