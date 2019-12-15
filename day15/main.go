package main

import (
	"github.com/gdamore/tcell"
	"github.com/glaming/advent-of-code-2019/intcode"
	"log"
	"sync"
)

type (
	point struct {
		x, y int
	}
)

const (
	north = iota + 1
	south
	west
	east
	complete
)

const (
	wall = iota
	moved
	found
	robot
)

func render(t int) (rune, tcell.Style) {
	s := tcell.StyleDefault.Background(tcell.ColorLightSlateGrey)

	switch t {
	case wall:
		return '█', s.Foreground(tcell.ColorBlue)
	case moved:
		return '.', s.Foreground(tcell.ColorWhite)
	case found:
		return 'X', s.Foreground(tcell.ColorDarkGreen)
	case robot:
		return 'O', s.Foreground(tcell.ColorWhite)
	default:
		return 'f', s
	}
}

func draw(screen tcell.Screen, p point, m map[point]int) {
	// Offsets for drawing
	var minX, minY int
	for k := range m {
		if k.x < minX {
			minX = k.x
		}
		if k.y < minY {
			minY = k.y
		}
	}

	screen.Clear()

	for k := range m {
		r, s := render(m[k])
		screen.SetContent(k.x-minX, k.y-minY, r, nil, s)
	}
	r, s := render(robot)
	screen.SetContent(p.x-minX, p.y-minY, r, nil, s)

	screen.Show()
}

func nextDirection(p point, path []point, m map[point]int) (int, point) {
	n := point{p.x, p.y - 1}
	s := point{p.x, p.y + 1}
	e := point{p.x + 1, p.y}
	w := point{p.x - 1, p.y}

	if _, ok := m[w]; !ok {
		return west, w
	} else if _, ok := m[s]; !ok {
		return south, s
	} else if _, ok := m[e]; !ok {
		return east, e
	} else if _, ok := m[n]; !ok {
		return north, n
	}

	if len(path) == 2 {
		return complete, path[0]
	}

	// Need to reverse
	prev := path[len(path)-2]
	switch prev {
	case n:
		return north, n
	case s:
		return south, s
	case e:
		return east, e
	case w:
		return west, w
	}

	return -1, point{}
}

func traverse(in, out chan int, screen tcell.Screen) (map[point]int, int) {
	m := make(map[point]int)

	p := point{0, 0}
	pathLength := 0
	path := make([]point, 0)

	for {
		dir, attempted := nextDirection(p, path, m)

		if dir == complete {
			break
		}

		in <- dir
		status := <-out
		m[attempted] = status

		switch status {
		case moved:
			p = attempted

			// If reversing
			if len(path) > 1 && path[len(path)-2] == p {
				path = path[:len(path)-1]
			} else {
				path = append(path, p)
			}
		case found:
			p = attempted
			path = append(path, p)
			pathLength = len(path)
		}

		if screen != nil {
			draw(screen, p, m)
		}
	}

	return m, pathLength
}

func main() {
	t, err := intcode.ReadTape("day15/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatal(err)
	}
	if err = screen.Init(); err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	in, out := make(chan int), make(chan int)

	wg.Add(1)
	go func() {
		_, pathLen := traverse(in, out, screen)
		screen.Clear()

		log.Printf("Path length: %d\n", pathLen)
		wg.Done()
	}()

	_, err = intcode.Execute(t, in, out)
	if err != nil {
		log.Fatal(err)
	}

	wg.Wait()
}
