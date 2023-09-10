package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	inputText := `
# This is a Header
This is a paragraph of text.

# Another Header
Here's another paragraph.
`

	// Using the blackfriday package
	// html := textToHTML(inputText)
	// html := string(blackfriday.Run([]byte(inputText)))
	// fmt.Println(html)

	headerRegex := regexp.MustCompile(`^#\s+(.*)$`)
	lines := strings.Split(inputText, "\n")
	for _, line := range lines {
		if matches := headerRegex.FindStringSubmatch(line); matches != nil {
			fmt.Printf("%v", matches[1])
		}
	}

	fmt.Println()
}
