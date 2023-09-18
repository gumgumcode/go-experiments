package main

import (
	"fmt"
	"math/rand"
	"slices"
)

var words = []string{
	"apple",
	"banana",
	"cherry",
	"orange",
	"grape",
	"strawberry",
	"watermelon",
}

func displayWord(word string, guessedLetters []string) string {
	display := ""
	for _, letter := range word {
		if slices.Contains(guessedLetters, string(letter)) {
			display += string(letter) + " "
		} else {
			display += "_ "
		}
	}
	return display
}

func main() {
	randomIndex := rand.Intn(len(words))
	chosenWord := words[randomIndex]
	guessedLetters := []string{"a"}

	for {
		fmt.Println("\n Word: ", displayWord(chosenWord, guessedLetters))
		break
	}
}
