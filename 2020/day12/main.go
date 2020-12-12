package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type (
	direction struct {
		action rune
		value  int
	}

	point struct {
		x, y int
	}
)

func readDirections(filename string) ([]direction, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	var directions []direction

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		var d direction
		_, err := fmt.Sscanf(line, "%c%d", &d.action, &d.value)
		if err != nil {
			return nil, err
		}

		directions = append(directions, d)
	}

	return directions, nil
}

func followDirections(ds []direction) point {
	pos := point{}
	orientations := []point{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}
	orientIndex := 0

	for _, d := range ds {
		switch d.action {
		case 'N':
			pos.y -= d.value
		case 'S':
			pos.y += d.value
		case 'E':
			pos.x += d.value
		case 'W':
			pos.x -= d.value
		case 'L':
			orientIndex = orientIndex - (d.value / 90) + 4
			orientIndex = orientIndex % 4
		case 'R':
			orientIndex = orientIndex + (d.value / 90)
			orientIndex = orientIndex % 4
		case 'F':
			pos.x += orientations[orientIndex].x * d.value
			pos.y += orientations[orientIndex].y * d.value
		}
	}

	return pos
}

func main() {
	directions, err := readDirections("2020/day12/input.txt")
	if err != nil {
		log.Panic(err)
	}

	pos := followDirections(directions)
	fmt.Println("Pos x, pos y:", pos.x, pos.y)
}
