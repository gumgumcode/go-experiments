package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	markDownFile string
	inParagraph  bool
	result       bytes.Buffer
)

func setMarkDownFilePath() {
	p, err := filepath.Abs("./main.go")
	if err != nil {
		fmt.Println("Error reading filepath: ", err)
		return
	}
	markDownFile = filepath.Dir(p) + "/text.md"
}

func maybeCloseParagraph() {
	if inParagraph {
		inParagraph = false
		result.WriteString("</p>")
	}
}

func textToHtml(text string) string {

	// Using the blackfriday package
	// html := textToHTML(inputText)
	// html := string(blackfriday.Run([]byte(inputText)))
	// fmt.Println(html)

	inParagraph = false

	headerRegex := regexp.MustCompile(`^#\s+(.*)$`)
	paragraphRegex := regexp.MustCompile(`^(.*)$`)

	lines := strings.Split(text, "\n")
	for _, line := range lines {
		if matches := headerRegex.FindStringSubmatch(line); matches != nil {
			maybeCloseParagraph()
			result.WriteString("<h1>")
			result.WriteString(matches[1])
			result.WriteString("</h1>")
		} else if matches := paragraphRegex.FindStringSubmatch(line); matches != nil {
			inParagraph = true
			result.WriteString("<p>")
			result.WriteString(matches[1])
		}
	}

	maybeCloseParagraph()
	return result.String()
}

func main() {
	setMarkDownFilePath()
	text, err := os.ReadFile(markDownFile)
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}
	fmt.Println(textToHtml(string(text)))
	// os.WriteFile("result.html", []byte(result.String()), 0644)
}
