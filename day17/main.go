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
	in, out := make(chan int), make(chan int)

	wg.Add(1)
	go func() {
		for {
			select {
			case v := <-out:
				output += string(rune(v))
			case <-time.After(100 * time.Millisecond):
				wg.Done()
				return
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

	// Reset the program... Part 2
	t, err = intcode.ReadTape("day17/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var collected int
	t[0] = 2
	input := "A,A,B,B,C,B,C,B,C,A\n"

	// A,A,B,B,C,B,C,B,C,A\nL,10,L,10,R,6\nR,12,L,12,L,12\nL,6,L,10,R,12,R,12\ny

	wg.Add(1)
	go func() {
	loop:
		for {
			select {
			case <-out:
			case <-time.After(500 * time.Millisecond):
				break loop
			}
		}

		for _, r := range input {
			in <- int(r)
		}

	loop2:
		for {
			select {
			case <-out:
			case <-time.After(500 * time.Millisecond):
				break loop2
			}
		}

		for _, r := range "L,10,L,10,R,6\n" {
			in <- int(r)
		}

	loop3:
		for {
			select {
			case <-out:
			case <-time.After(500 * time.Millisecond):
				break loop3
			}
		}

		for _, r := range "R,12,L,12,L,12\n" {
			in <- int(r)
		}

	loop4:
		for {
			select {
			case <-out:
			case <-time.After(500 * time.Millisecond):
				break loop4
			}
		}

		for _, r := range "L,6,L,10,R,12,R,12\n" {
			in <- int(r)
		}

	loop5:
		for {
			select {
			case <-out:
			case <-time.After(500 * time.Millisecond):
				break loop5
			}
		}

		in <- int('y')
		in <- int('\n')

	loop6:
		for {
			select {
			case collected = <-out:
			case <-time.After(500 * time.Millisecond):
				break loop6
			}
		}

		wg.Done()
	}()

	_, err = intcode.Execute(t, in, out)
	if err != nil {
		log.Fatal(err)
	}

	wg.Wait()
	println(collected)

}
