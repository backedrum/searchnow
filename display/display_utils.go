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
