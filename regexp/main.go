package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Using MustCompile (panics if invalid pattern)
	pattern := regexp.MustCompile(`\d{3}-\d{2}-\d{4}`)

	// Using Compile (returns error if invalid pattern)
	pattern, err := regexp.Compile(`\d{3}-\d{2}-\d{4}`)
	if err != nil {
		// Handle error
	}

	input := "123-45-6789"
	if pattern.MatchString(input) {
		fmt.Println("Pattern matched")
	}

	input = "Phone numbers: 123-45-6789, 987-65-4321"

	match := pattern.FindString(input)
	fmt.Println("First match:", match) // First match: 123-45-6789

	matches := pattern.FindAllString(input, -1)
	fmt.Println("All matches:", matches) // All matches: [123-45-6789 987-65-4321]

	allMatches := pattern.FindAllStringSubmatch(input, -1)
	fmt.Printf("All sub matches: %q \n", allMatches) // All sub matches: [[123-45-6789] [987-65-4321]]

	replacement := "XXX-XX-XXXX"
	result := pattern.ReplaceAllString(input, replacement)
	fmt.Println("Result:", result) // Result: Phone numbers: XXX-XX-XXXX, XXX-XX-XXXX

	pattern = regexp.MustCompile(`\s+`)
	input = "Hello     World    Golang"
	// tokens := pattern.Split(input, 1) // Tokens: ["Hello     World    Golang"]
	// tokens := pattern.Split(input, 2) // Tokens: ["Hello" "World    Golang"]
	// tokens := pattern.Split(input, 3) // Tokens: ["Hello" "World" "Golang"]
	tokens := pattern.Split(input, -1) // Tokens: ["Hello" "World" "Golang"]
	fmt.Printf("Tokens: %q\n", tokens)

	pattern = regexp.MustCompile(`(?i)go`)
	input = "Golang is great! Go is also awesome! go go go"
	matches = pattern.FindAllString(input, -1)
	fmt.Println("Case-insensitive matches:", matches) // Case-insensitive matches: [Go Go go go go]
}
