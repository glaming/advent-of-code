package main

import (
	"github.com/glaming/advent-of-code-2019/intcode"
	"log"
)

type (
	point struct {
		x, y int
	}

	tile int
)

const (
	empty tile = iota
	wall
	block
	horizontal
	ball
)

func main() {
	t, err := intcode.ReadTape("day13/input.txt")
	if err != nil {
		log.Panic(err)
	}

	grid := make(map[point]tile)
	in, out := make(chan int), make(chan int)

	go func() {
		for {
			x := <-out
			y := <-out
			t := <-out

			grid[point{x, y}] = tile(t)
		}
	}()

	_, err = intcode.Execute(t, in, out)
	if err != nil {
		log.Panic(err)
	}

	var countBlock int
	for k := range grid {
		if grid[k] == block {
			countBlock++
		}
	}

	log.Printf("block count %d\n", countBlock)
}
