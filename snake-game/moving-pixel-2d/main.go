package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
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

func main() {
	rows := 10
	cols := 20
	p := Point{x: 0, y: 0}
	dir := Right

	for {
		clearScreen()

		for i := 0; i < rows; i++ { // i is row, i.e. y axis
			for j := 0; j < cols; j++ { // j is column, i.e. x axis

				if p.x == j && p.y == i {
					fmt.Print("O")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}

		adjustPosition(&p, &dir, rows, cols)
		fmt.Println(directionNames[dir])

		// Speed of pixel's movement
		time.Sleep(time.Millisecond * 100)
	}

}

func adjustPosition(p *Point, dir *Direction, rows, cols int) {
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

	if p.x >= cols || p.x < 0 || p.y >= rows || p.y < 0 {
		switch *dir {
		case Right:
			p.x--
		case Down:
			p.y--
		case Left:
			p.x++
		case Up:
			p.y++
		}
		*dir = (*dir + 1) % 4
	}
}
