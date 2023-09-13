package main

import (
	"fmt"
	"os"

	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	fmt.Println("Press arrow keys (↑, ↓, ←, →) to see the code for the key pressed. Press 'q' to quit.")

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		if key == keyboard.KeyArrowUp {
			fmt.Println("Up Arrow Key Pressed")
		} else if key == keyboard.KeyArrowDown {
			fmt.Println("Down Arrow Key Pressed")
		} else if key == keyboard.KeyArrowLeft {
			fmt.Println("Left Arrow Key Pressed")
		} else if key == keyboard.KeyArrowRight {
			fmt.Println("Right Arrow Key Pressed")
		} else if key == keyboard.KeyEsc || char == 'q' {
			fmt.Println("Exiting...")
			os.Exit(0)
		}
	}
}
