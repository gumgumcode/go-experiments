package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/eiannone/keyboard"
)

const (
	ROWS  = 20
	COLS  = 40
	SPEED = 200
)

type Direction int

const (
	Right Direction = iota
	Down
	Left
	Up
)

var directionNames = map[Direction]string{
	Right: "Right",
	Down:  "Down",
	Left:  "Left",
	Up:    "Up",
}

type Point struct {
	x, y int
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func updatePosition(p *Point, dir *Direction) {
	switch *dir {
	case Right:
		p.x++
	case Down:
		p.y++
	case Left:
		p.x--
	case Up:
		p.y--
	}

	if p.x >= COLS {
		p.x = 0
	} else if p.x < 0 {
		p.x = COLS - 1
	}

	if p.y >= ROWS {
		p.y = 0
	} else if p.y < 0 {
		p.y = ROWS - 1
	}
}

func maybeUpdateDirection(dir *Direction) {
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		if key == keyboard.KeyArrowRight {
			*dir = Right
		} else if key == keyboard.KeyArrowDown {
			*dir = Down
		} else if key == keyboard.KeyArrowLeft {
			*dir = Left
		} else if key == keyboard.KeyArrowUp {
			*dir = Up
		} else if key == keyboard.KeyEsc || char == 'q' {
			os.Exit(0)
		}
	}
}

func main() {
	p := Point{x: 0, y: 0}
	dir := Right

	go maybeUpdateDirection(&dir)

	for {
		clearScreen()

		for i := 0; i < ROWS; i++ {
			for j := 0; j < COLS; j++ {
				if p.x == j && p.y == i {
					fmt.Print("0")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}

		updatePosition(&p, &dir)
		fmt.Println(directionNames[dir])

		time.Sleep(time.Millisecond * SPEED)
	}
}
