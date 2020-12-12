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

func rotateWaypoint(d rune, val int, wp point) point {
	for i := 0; i < val/90; i++ {
		oldWp := point{wp.x, wp.y}
		switch d {
		case 'L':
			wp.x = oldWp.y
			wp.y = oldWp.x * -1
		case 'R':
			wp.x = oldWp.y * -1
			wp.y = oldWp.x
		}
	}
	return wp
}

func followDirections(ds []direction) point {
	pos, waypoint := point{}, point{10, -1}

	for _, d := range ds {
		switch d.action {
		case 'N':
			waypoint.y -= d.value
		case 'S':
			waypoint.y += d.value
		case 'E':
			waypoint.x += d.value
		case 'W':
			waypoint.x -= d.value
		case 'L', 'R':
			waypoint = rotateWaypoint(d.action, d.value, waypoint)
		case 'F':
			pos.x += waypoint.x * d.value
			pos.y += waypoint.y * d.value
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
