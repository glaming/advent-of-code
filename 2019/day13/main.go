package main

import (
	"github.com/gdamore/tcell"
	"github.com/glaming/advent-of-code-2019/intcode"
	"log"
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

func (t tile) char() rune {
	switch t {
	case empty:
		return ' '
	case wall:
		return '|'
	case block:
		return '#'
	case horizontal:
		return '-'
	case ball:
		return 'O'
	default:
		return 'X'
	}
}

func draw(screen tcell.Screen, p point, t tile) {
	screen.SetContent(p.x, p.y, t.char(), nil, 0)
	screen.Show()
}

func main() {
	t, err := intcode.ReadTape("day13/input.txt")
	if err != nil {
		log.Panic(err)
	}

	t[0] = 2

	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatal(err)
	}
	if err = screen.Init(); err != nil {
		log.Fatal(err)
	}

	var score, paddleX, joystick int

	grid := make(map[point]tile)
	in, out := make(chan int), make(chan int)

	go func() {
		for {
			select {
			case x := <-out:
				y := <-out
				t := <-out

				if x == -1 && y == 0 {
					score = t
				} else {
					grid[point{x, y}] = tile(t)

					if tile(t) == horizontal {
						paddleX = x
					}
					if tile(t) == ball {
						if paddleX < x {
							joystick = 1
						} else if paddleX > x {
							joystick = -1
						} else {
							joystick = 0
						}
					}

					draw(screen, point{x, y}, tile(t))
				}

			case <-time.After(10 * time.Millisecond):
				in <- joystick
				joystick = 0
			}
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

	screen.Fini()

	log.Printf("block count %d\n", countBlock)
	log.Printf("score %d\n", score)
}
