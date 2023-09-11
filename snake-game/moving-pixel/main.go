package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	width := 40
	x := 0

	for {
		clearScreen()

		// Draw the canvas
		for i := 0; i < width; i++ {
			if i == x {
				fmt.Print("◼︎")
			} else {
				fmt.Print(" ")
			}
		}

		// Move the pixel to the right
		x++

		// Wrap the pixel around when it reaches the right edge
		if x >= width {
			x = 0
		}

		// Speed of pixel's movement
		time.Sleep(time.Millisecond * 100)
	}
}
