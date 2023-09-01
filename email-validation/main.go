package main

import (
	"fmt"
	"regexp"
)

func isValidEmail(email string) bool {
	emailPattern := "^[A-Za-z0-9._%-]+@[A-Za-z0-9.-]+\\.[A-Za-z]{2,}"
	emailRegex := regexp.MustCompile(emailPattern)
	return emailRegex.MatchString(email)
}

func main() {
	emails := []string{"omkar@google.com", "omkar@@google.com"}
	for _, email := range emails {
		fmt.Println(isValidEmail(email))
	}
}
