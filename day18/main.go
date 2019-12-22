package main

import (
	"bufio"
	"log"
	"os"
)

type (
	point struct {
		x, y int
	}

	state struct {
		point
		keys int
	}

	stateSteps struct {
		state
		steps int
	}
)

func readMap(filename string) ([][]rune, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	var m [][]rune

	s := bufio.NewScanner(file)
	for s.Scan() {
		line := s.Text()
		m = append(m, []rune(line))
	}

	return m, s.Err()
}

func isDoor(r rune) bool {
	if r >= 'A' && r <= 'Z' {
		return true
	}
	return false
}

func isKey(r rune) bool {
	if r >= 'a' && r <= 'z' {
		return true
	}
	return false
}

func findShortestPath(m [][]rune, start point, keys, totalKeys int) int {
	queue := []stateSteps{{state{start, keys}, 0}}
	visited := make(map[state]bool)

	for len(queue) > 0 {
		var head stateSteps
		head, queue = queue[0], queue[1:]

		if head.keys == totalKeys {
			return head.steps
		}

		visited[head.state] = true

		adjacent := []point{
			{head.x - 1, head.y},
			{head.x + 1, head.y},
			{head.x, head.y - 1},
			{head.x, head.y + 1},
		}

		for _, a := range adjacent {
			r := m[a.y][a.x]
			if visited[state{a, head.keys}] {
				continue
			}
			if r == '#' {
				continue
			}
			if isDoor(r) {
				// If not got a key for this door
				if head.keys&(1<<(uint(r)-'A')) != 1<<(uint(r)-'A') {
					continue
				}
			}

			s := stateSteps{state{a, head.keys}, head.steps + 1}

			if isKey(r) {
				// If found new key
				if head.keys&(1<<(uint(r)-'a')) != 1<<(uint(r)-'a') {
					s.keys |= 1 << (uint(r) - 'a')
				}
			}

			queue = append(queue, s)
		}
	}

	return -1
}

func main() {
	m, err := readMap("day18/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	//sp := findShortestPath(m, 0)
	//log.Printf("Shortest path: %d\n", sp)

	// Part 2 - modify map
	m[39][40] = '#'
	m[39][40] = '#'
	m[40][39] = '#'
	m[40][40] = '#'
	m[40][41] = '#'
	m[41][40] = '#'

	totalKeys := 0
	for _, line := range m {
		for _, r := range line {
			if isKey(r) {
				k := 1 << (uint(r) - 'a')
				totalKeys |= k
			}
		}
	}

	segments := []struct{ a, b, start point }{
		{point{0, 0}, point{40, 40}, point{39, 39}},
		{point{40, 0}, point{80, 40}, point{41, 39}},
		{point{0, 40}, point{40, 80}, point{39, 41}},
		{point{40, 40}, point{80, 80}, point{41, 41}},
	}

	totalSteps := 0
	for _, s := range segments {
		segmentKeys := 0
		for y := s.a.y; y <= s.b.y; y++ {
			for x := s.a.x; x <= s.b.x; x++ {
				r := m[y][x]
				if isKey(r) {
					k := 1 << (uint(r) - 'a')
					segmentKeys |= k
				}
			}
		}

		totalSteps += findShortestPath(m, s.start, totalKeys-segmentKeys, totalKeys)
	}

	log.Printf("Shortest path: %d\n", totalSteps)
}
