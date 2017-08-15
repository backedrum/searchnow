package main

import (
	"fmt"
	"github.com/backedrum/searchnow/handlers"
	tm "github.com/buger/goterm"
	"os"
	"strconv"
)

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

	results := handlers.Search(engine, searchTerm, numOfResults)

	tm.Clear()
	resultTable := tm.NewTable(0, 20, 5, ' ', 0)
	fmt.Fprintf(resultTable, "Url\tTitle\n")

	for _, result := range results {
		fmt.Fprintf(resultTable, "%s\t%s\n", result.Url, result.Title)
	}
	tm.Println(resultTable)

	tm.Flush()

	return
}
