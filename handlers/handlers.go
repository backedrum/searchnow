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
	"bufio"
	"fmt"
	"os"
	"strings"
)

type SearchResult struct {
	Source      string
	Url         string
	Title       string
	Contents    string
	Extras      map[string]string
	ExtrasOrder []string
}

type searchFn func(searchTerm string, numOfResults int) []*SearchResult

var tokens = make(map[string]string)

var engines = map[string]searchFn{
	"google": searchViaGoogleApi,
	"so":     searchStackOverflow,
	"hn":     fetchHackerNews,
	"ip_loc": searchIpLocation,
}

var enginesAliases = map[string]string{
	"g":              "google",
	"stack":          "so",
	"stack-overflow": "so",
	"hacker":         "hn",
	"ip":             "ip_loc",
}

// ResolveTargetEngineName Identify target engine by the given input.
// Input can be either exact engine name or alias
func ResolveTargetEngineName(input string) string {
	if val, exists := enginesAliases[input]; exists {
		return val
	}
	return input
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
		return
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
