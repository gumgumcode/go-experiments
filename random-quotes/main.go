package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	quotes := []string{
		"Life is 10% what happens to us and 90% how we react to it. - Charles R. Swindoll",
		"The only way to do great work is to love what you do. - Steve Jobs",
		"Success is not final, failure is not fatal: It is the courage to continue that counts. - Winston Churchill",
		"Believe you can and you're halfway there. - Theodore Roosevelt",
		"Strive not to be a success, but rather to be of value. - Albert Einstein",
	}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(quotes))

	// line above ^ generates a random integer between 0 (inclusive) and the length of the quotes slice (exclusive)
	// Intn(n int) function is used to generate a random non-negative integer less than n

	fmt.Println("Today's Random Quote: ")
	fmt.Println(quotes[randomIndex])
}
