package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readSeatLayout(filename string) ([][]rune, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	var layout [][]rune

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		layout = append(layout, []rune(line))
	}

	return layout, nil
}

func copyLayout(layout [][]rune) [][]rune {
	var newLayout [][]rune
	for _, row := range layout {
		newRow := make([]rune, len(row))
		copy(newRow, row)
		newLayout = append(newLayout, newRow)
	}
	return newLayout
}

func isOccupied(x, y int, layout [][]rune) bool {
	if x < 0 || y < 0 {
		return false
	}
	if x >= len(layout[0]) || y >= len(layout) {
		return false
	}
	return layout[y][x] == '#'
}

func countAdjacentOccupied(x, y int, layout [][]rune) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if isOccupied(x + i, y + j, layout) {
				count++
			}
		}
	}
	return count
}

func applyRules(layout [][]rune) ([][]rune, int) {
	width, height := len(layout[0]), len(layout)

	newLayout := copyLayout(layout)

	changes := 0

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if layout[y][x] == '.' {
				continue
			}

			count := countAdjacentOccupied(x, y, layout)

			if layout[y][x] == 'L' && count == 0 {
				newLayout[y][x] = '#'
				changes++
			} else if layout[y][x] == '#' && count >= 4 {
				newLayout[y][x] = 'L'
				changes++
			}
		}
	}

	return newLayout, changes
}

func main() {
	layout, err := readSeatLayout("2020/day11/input.txt")
	if err != nil {
		log.Panic(err)
	}

	for {
		var changes int
		layout, changes = applyRules(layout)
		if changes == 0 {
			break
		}
	}

	occupied := 0
	for _, row := range layout {
		for _, seat := range row {
			if seat == '#' {
				occupied++
			}
		}
	}

	fmt.Println("Occupied:", occupied)
}
