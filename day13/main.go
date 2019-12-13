package main

import (
	"github.com/glaming/advent-of-code-2019/intcode"
	"log"
	"sync"
	"time"
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

	var wg sync.WaitGroup
	in, out := make(chan int), make(chan int)

	wg.Add(1)
	go func() {
		for {
			select {
			case x := <-out:
				y := <-out
				t := <-out

				grid[point{x, y}] = tile(t)
			case <-time.After(500 * time.Millisecond):
				wg.Done()
				return
			}
		}
	}()

	_, err = intcode.Execute(t, in, out)
	if err != nil {
		log.Panic(err)
	}

	wg.Wait()

	var countBlock int
	for k := range grid {
		if grid[k] == block {
			countBlock++
		}
	}

	log.Printf("block count %d\n", countBlock)
}
