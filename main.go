package main

import (
	"encoding/json"
	"fmt"
	"github.com/backedrum/searchnow/display"
	"github.com/backedrum/searchnow/handlers"
	tm "github.com/buger/goterm"
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

	for _, result := range results {
		for i := 0; i < tm.Width(); i++ {
			fmt.Print(tm.Color("_", tm.CYAN))
		}
		fmt.Print("\n")

		display.PutLine("URL:", result.Url, tm.RED, tm.BLUE, config.ShowURL)
		display.PutLine("Title:", result.Title, tm.RED, -1, config.ShowTitle)

		display.PutLine("Snippet:", display.ConvertHtmlToText(result.Contents), tm.RED, tm.GREEN, config.ShowContents)

		if config.ShowContents && len(result.Extras) > 0 {
			for _, extra := range result.ExtrasOrder {
				display.PutLine(extra, display.ConvertHtmlToText(result.Extras[extra]), tm.RED, tm.YELLOW, true)
			}
		}
	}

	tm.Flush()

	return
}
