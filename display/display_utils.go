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
package display

import (
	"fmt"
	tm "github.com/buger/goterm"
	"github.com/jaytaylor/html2text"
	"html"
)

// PutLine puts a formatted line into the output.
// Params are: name - name of the field, value - value of the field, nameColor,valueColor - text color,
// allowed - if true then line will be printed.
func PutLine(name, value string, nameColor, valueColor int, allowed bool) {
	if allowed {
		if nameColor > -1 {
			fmt.Print(tm.Color(name, nameColor) + "\t")
		} else {
			fmt.Print(name + "\t")
		}

		unescapedValue := html.UnescapeString(value)
		if valueColor > -1 {
			fmt.Print(tm.Color(unescapedValue, valueColor) + "\n")
		} else {
			fmt.Print(unescapedValue + "\n")
		}
	}
}

// ConvertHtmlToText performs conversion from html string to txt string
func ConvertHtmlToText(html string) string {
	txt, err := html2text.FromString(html, html2text.Options{PrettyTables: true})
	if err != nil {
		// fallback to raw html
		txt = html
	}

	return txt
}
