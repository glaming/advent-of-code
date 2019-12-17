package main

import (
	"bufio"
	"github.com/glaming/advent-of-code-2019/intcode"
	"log"
	"strings"
	"sync"
	"time"
)

type point struct {
	x, y int
}

func readCamera(output string) (map[point]rune, error) {
	s := bufio.NewScanner(strings.NewReader(output))

	m := make(map[point]rune)

	row := 0
	for s.Scan() {
		line := s.Text()
		for i, r := range line {
			m[point{i, row}] = r
		}
		row++
	}

	return m, s.Err()
}

func isScaffolding(r rune) bool {
	if r == '#' || r == '^' {
		return true
	}
	return false
}

func sumAlignmentParameters(m map[point]rune) int {
	sum := 0

	for k := range m {
		if isScaffolding(m[k]) {
			around := []point{
				{k.x, k.y - 1},
				{k.x, k.y + 1},
				{k.x - 1, k.y},
				{k.x + 1, k.y},
			}

			valid := true
			for _, v := range around {
				if r, ok := m[v]; !ok || !isScaffolding(r) {
					valid = false
					break
				}
			}

			if valid {
				sum += k.x * k.y
			}
		}
	}

	return sum
}

func main() {
	t, err := intcode.ReadTape("day17/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var output string
	var wg sync.WaitGroup
	out := make(chan int)

	wg.Add(1)
	go func() {
		for {
			select {
			case v := <-out:
				output += string(rune(v))
			case <-time.After(100 * time.Millisecond):
				wg.Done()
			}
		}
	}()

	_, err = intcode.Execute(t, nil, out)
	if err != nil {
		log.Fatal(err)
	}

	wg.Wait()

	m, err := readCamera(output)
	if err != nil {
		log.Fatal(err)
	}

	sum := sumAlignmentParameters(m)
	println(sum)

	println(output)
}
