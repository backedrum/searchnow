package handlers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type SearchResult struct {
	Source   string
	Url      string
	Title    string
	Contents string
}

type searchFn func(searchTerm string, numOfResults int) []*SearchResult

var tokens = make(map[string]string)

var engines = map[string]searchFn{
	"google": searchViaGoogleApi,
	"so":     searchStackOverflow,
}

func HasEngineSupport(engine string) bool {
	return engines[engine] != nil
}

func Search(engine, searchTerm string, numOfResults int) []*SearchResult {
	return engines[engine](searchTerm, numOfResults)
}

//TODO make .tokens files an optional command line arg
func init() {
	file, err := os.Open(".tokens")
	if err != nil {
		fmt.Printf("Cannot read .tokens file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(bufio.NewReader(file))
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "#") {
			continue
		}

		values := strings.Split(line, "=")
		tokens[values[0]] = values[1]
	}
}
