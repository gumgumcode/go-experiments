package main

import (
	"fmt"
	"math/rand"
	"slices"
	"strings"
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
	guessedLetters := []string{}

	maxIncorrectGuesses := 6
	incorrectGuesses := 0

	for {
		// Display the current state of the word
		fmt.Println("\nWord:", displayWord(chosenWord, guessedLetters))

		// Prompt the player for a letter guess
		fmt.Print("Guess a letter: ")
		var guess string
		fmt.Scanln(&guess)

		// Validate the guess (ensure it's a single letter)
		if len(guess) != 1 || !isLetter(guess) {
			fmt.Println("Invalid input. Please enter a single letter.")
			continue
		}

		// Check if the guess is correct
		if strings.Contains(chosenWord, guess) {
			// Update the guessedLetters slice with the correct guess
			guessedLetters = append(guessedLetters, guess)
			// Check if the player has won
			if strings.EqualFold(chosenWord, displayWord(chosenWord, guessedLetters)) {
				fmt.Println("Congratulations! You've won. The word was:", chosenWord)
				break
			}
		} else {
			fmt.Println("Incorrect guess.")
			incorrectGuesses++
			// Check if the player has lost
			if incorrectGuesses >= maxIncorrectGuesses {
				fmt.Println("Sorry, you've run out of guesses. The word was:", chosenWord)
				break
			}
		}
	}
	fmt.Println("\nThank you for playing Hangman!")
}

func isLetter(s string) bool {
	if len(s) != 1 {
		return false
	}
	return ('a' <= s[0] && s[0] <= 'z') || ('A' <= s[0] && s[0] <= 'Z')
}
