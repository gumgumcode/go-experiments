package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const (
	uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	digits           = "0123456789"
	specialChars     = "!@#$%^&*()_+{}[]:;<>,.?/~"
)

func generatePassword(length int) (string, error) {
	// Define the character set for the password
	characterSet := uppercaseLetters + lowercaseLetters + digits + specialChars

	// Create a random password of the specified length
	password := make([]byte, length)
	for i := 0; i < length; i++ {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(characterSet))))
		if err != nil {
			return "", err
		}
		password[i] = characterSet[randomIndex.Int64()]
	}

	return string(password), nil
}

func main() {
	password, err := generatePassword(12) // Change the length as needed
	if err != nil {
		fmt.Println("Error generating password:", err)
		return
	}
	fmt.Println("Random Password:", password)
}
