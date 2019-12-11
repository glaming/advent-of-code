package main

import (
	"github.com/glaming/advent-of-code-2019/intcode"
	"log"
)

type (
	point struct {
		x, y int
	}

	direction int
)

const (
	up    direction = 0
	right direction = 1
	down  direction = 2
	left  direction = 3
)

func draw(m map[point]int) {
	println()
	var maxX, maxY int
	for k := range m {
		if k.x < 0 {
			panic("x < 0")
		}
		if k.y < 0 {
			panic("y < 0")
		}

		if k.x > maxX {
			maxX = k.x
		}
		if k.y > maxY {
			maxY = k.y
		}
	}

	for x := 0; x <= maxX; x++ {
		for y := 0; y <= maxY; y++ {
			v, ok := m[point{x, y}]
			if ok && v == 1 {
				print("#")
			} else {
				print(" ")
			}
		}
		println()
	}
}

func main() {
	t, err := intcode.ReadTape("day11/input.txt")
	if err != nil {
		log.Panic(err)
	}

	in, out := make(chan int), make(chan int)

	m := make(map[point]int)
	currPos := point{50, 50}
	currDirection := up

	m[currPos] = 1

	go func() {
		for {
			in <- m[currPos]
			m[currPos] = <-out
			dirTurn := <-out

			currDirection = direction((int(currDirection) - 1 + (dirTurn * 2) + 4) % 4)

			switch currDirection {
			case up:
				currPos.y--
			case right:
				currPos.x++
			case down:
				currPos.y++
			case left:
				currPos.x--
			}
		}
	}()

	_, err = intcode.Execute(t, in, out)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Panels painted: %d\n", len(m))

	draw(m)
}
