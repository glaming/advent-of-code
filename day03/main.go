package main

import (
	"encoding/csv"
	"log"
	"os"
	"sort"
	"strconv"
)

type (
	point struct {
		x, y int
	}

	path []point

	direction string

	instruction struct {
		direction
		distance int
	}
)

const (
	directionUp    direction = "U"
	directionDown  direction = "D"
	directionLeft  direction = "L"
	directionRight direction = "R"
)

func readWirePathInstructions(filename string) ([][]instruction, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	paths, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, err
	}

	var instructionPaths [][]instruction

	for _, p := range paths {
		var is []instruction
		for _, iStr := range p {
			distance, err := strconv.Atoi(iStr[1:])
			if err != nil {
				return nil, err
			}

			i := instruction{direction(iStr[0]), distance}
			is = append(is, i)
		}

		instructionPaths = append(instructionPaths, is)
	}

	return instructionPaths, nil
}

func executePathInstructions(is []instruction) path {
	 var p path
	 var pos point

	for _, i := range is {
		for d := 0; d < i.distance; d++ {
			switch i.direction {
			case directionDown:
				pos.y = pos.y - 1
			case directionUp:
				pos.y = pos.y + 1
			case directionLeft:
				pos.x = pos.x - 1
			case directionRight:
				pos.x = pos.x + 1
			}

			p = append(p, pos)
		}
	}

	return p
}

func findMatchingPoints(a, b path) []point {
	var matches []point
	for _, aa := range a {
		for _, bb := range b {
			if aa.x == bb.x && aa.y == bb.y {
				matches = append(matches, aa)
				continue
			}
		}
	}

	return matches
}

func absInt(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func findClosestPointToOrigin(ps []point) point {
	sort.Slice(ps, func(a, b int) bool {
		aDist := absInt(ps[a].x) + absInt(ps[a].y)
		bDist := absInt(ps[b].x) + absInt(ps[b].y)
		return aDist < bDist
	})

	return ps[0]
}

func main() {
	pis, err := readWirePathInstructions("day03/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var paths []path
	for _, pi := range pis {
		p := executePathInstructions(pi)

		paths = append(paths, p)
	}

	ps := findMatchingPoints(paths[0], paths[1])

	p := findClosestPointToOrigin(ps)

	log.Printf("Closest point to origin: %v\n", p)
}
