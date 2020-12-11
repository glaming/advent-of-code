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

func isOccupied(x, y int, layout [][]rune) (isOccupied, isEmptySeat, isValid bool) {
	if x < 0 || y < 0 {
		return false, false,false
	}
	if x >= len(layout[0]) || y >= len(layout) {
		return false, false,false
	}
	return layout[y][x] == '#', layout[y][x] == 'L', true
}

func countVisibleOccupied(x, y int, layout [][]rune) int {
	count := 0
	directions := []struct{ x, y int }{
		{1, 0},
		{1, 1},
		{0, 1},
		{-1, 1},
		{-1, 0},
		{-1, -1},
		{0, -1},
		{1, -1},
	}

	for _, d := range directions {
		for i := 1; true; i++ {
			isOcc, isEmpty, valid := isOccupied(d.x * i + x, d.y * i + y, layout)
			if !valid || isEmpty {
				break
			}
			if isOcc {
				count++
				break
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

			count := countVisibleOccupied(x, y, layout)

			if layout[y][x] == 'L' && count == 0 {
				newLayout[y][x] = '#'
				changes++
			} else if layout[y][x] == '#' && count >= 5 {
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
