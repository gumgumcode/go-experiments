package main

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func main() {
	if res, err := readSingleCharacter(); err != nil {
		fmt.Print("Error:", err)
	} else {
		fmt.Printf("The char %q was hit and submitted without pressing enter.\n", res)
	}
}

func readSingleCharacter() (string, error) {
	var err error
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	b := make([]byte, 1)
	_, err = os.Stdin.Read(b)
	if err != nil {
		return "", err
	}
	return string(b[0]), err
}
