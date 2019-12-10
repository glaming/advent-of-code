package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"math"
	"os"
	"sort"
	"strings"
)

type point struct {
	x, y int
}

func readMapFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	b, err := ioutil.ReadAll(file)
	return string(b), err
}

func getAsteroids(m string) ([]point, error) {
	var asteroids []point

	s := bufio.NewScanner(strings.NewReader(m))

	line := 0
	for s.Scan() {
		for i, r := range s.Text() {
			if r == '#' {
				asteroids = append(asteroids, point{i, line})
			}
		}
		line++
	}

	return asteroids, s.Err()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Is point p between a and b
func isBetween(a, b, p point) bool {
	crossProduct := (p.y-a.y)*(b.x-a.x) - (p.x-a.x)*(b.y-a.y)
	sameLine := crossProduct == 0
	if !sameLine {
		return false
	}

	betweenX := min(a.x, b.x) <= p.x && p.x <= max(a.x, b.x)
	betweenY := min(a.y, b.y) <= p.y && p.y <= max(a.y, b.y)
	return betweenX && betweenY
}

func asteroidsVisibleFrom(p point, asteroids []point) int {
	count := 0
loop:
	for _, a := range asteroids {
		// Don't count itself
		if a == p {
			continue
		}

		for _, b := range asteroids {
			if a == b || p == b {
				continue
			}

			is := isBetween(p, a, b)
			if is {
				continue loop
			}
		}
		count++
	}
	return count
}

func bestLocation(asteroids []point) (point, int) {
	var best point
	count := -1

	for _, a := range asteroids {
		c := asteroidsVisibleFrom(a, asteroids)

		if c > count {
			best = a
			count = c
		}
	}
	return best, count
}

func angle(a, b point) float64 {
	return math.Mod(math.Atan2(float64(b.x-a.x), float64(a.y-b.y)), 2*math.Pi)
}

// Return the order in which they're vaporised
func vaporise(p point, asteroids []point) []point {
	lines := map[float64][]point{}

	for _, aster := range asteroids {
		if aster == p {
			continue
		}

		a := angle(p, aster)

		if a < 0 {
			a += 2 * math.Pi
		}

		lines[a] = append(lines[a], aster)
	}

	var angles []float64
	for a := range lines {
		angles = append(angles, a)
		sort.Slice(lines[a], func(i, j int) bool {
			return isBetween(p, lines[a][j], lines[a][i])
		})
	}
	sort.Float64s(angles)

	var vaporised []point
	for {
		moreToVaporise := false

		for _, a := range angles {
			if len(lines[a]) > 0 {
				var x point
				x, lines[a] = lines[a][0], lines[a][1:]
				vaporised = append(vaporised, x)

				moreToVaporise = true
			}
		}

		if !moreToVaporise {
			break
		}
	}

	return vaporised
}

func main() {
	m, err := readMapFile("day10/input.txt")
	if err != nil {
		log.Panic(err)
	}

	as, err := getAsteroids(m)
	if err != nil {
		log.Panic(err)
	}

	a, count := bestLocation(as)

	log.Printf("Best position: %v\nBest count: %d\n", a, count)

	as = vaporise(a, as)

	log.Printf("200th to be vaporised: %v\n", as[199])
}
