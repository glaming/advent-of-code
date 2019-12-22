package main

import (
	"bufio"
	"crypto/sha1"
	"fmt"
	"log"
	"os"
	"sort"
)

type (
	point struct {
		x, y int
	}

	state struct {
		point
		visited map[point]bool
		keys    map[rune]bool
		steps   int
	}
)

func (s state) hash() string {
	h := sha1.New()

	keys := make([]int, 0)
	for k := range s.keys {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)

	fmt.Fprintf(h, "%d,%d", s.point.x, s.point.y)
	fmt.Fprintf(h, "%v", keys)
	return fmt.Sprintf("%x", h.Sum(nil))
}

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

func newState(prev state, p point) state {
	s := state{p, make(map[point]bool), make(map[rune]bool), prev.steps + 1}
	for k := range prev.visited {
		s.visited[k] = true
	}
	for k := range prev.keys {
		s.keys[k] = true
	}
	return s
}

func findShortestPath(m [][]rune) int {
	// Find start + how many keys
	start, totalKeys := point{}, 0
	for y, line := range m {
		for x, r := range line {
			if r == '@' {
				start = point{x, y}
			}
			if isKey(r) {
				totalKeys++
			}
		}
	}

	queue := []state{{start, make(map[point]bool), make(map[rune]bool), 0}}
	visited := make(map[string]bool)

	for len(queue) > 0 {
		var head state
		head, queue = queue[0], queue[1:]

		if len(head.keys) == totalKeys {
			return head.steps
		}

		head.visited[head.point] = true
		visited[head.hash()] = true

		adjacent := []point{
			{head.x - 1, head.y},
			{head.x + 1, head.y},
			{head.x, head.y - 1},
			{head.x, head.y + 1},
		}

		for _, a := range adjacent {
			r := m[a.y][a.x]
			if head.visited[a] {
				continue
			}
			if r == '#' {
				continue
			}
			if _, ok := head.keys[r+('a'-'A')]; isDoor(r) && !ok {
				continue
			}

			s := newState(head, a)

			// If found new key, add key and clear path
			if _, ok := head.keys[r]; isKey(r) && !ok {
				s.keys[r] = true
				s.visited = make(map[point]bool)
			}

			if _, ok := visited[s.hash()]; !ok {
				queue = append(queue, s)
			}
		}
	}

	return -1
}

func main() {
	m, err := readMap("day18/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	sp := findShortestPath(m)
	log.Printf("Shortest path: %d\n", sp)
}
