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

	node struct {
		point
		depth int
	}
)

func isPortal(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func isPathway(r rune) bool {
	return r == '.'
}

func readDonut(filename string) ([]string, map[point]string, map[string][]point, point, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, nil, point{}, err
	}

	var donut []string
	s := bufio.NewScanner(file)
	for s.Scan() {
		donut = append(donut, s.Text())
	}
	if s.Err() != nil {
		return nil, nil, nil, point{}, s.Err()
	}

	var maxPoint point
	portals := make(map[point]string)
	for i, line := range donut {
		for x, r := range line {
			if isPortal(r) {
				if len(line) > x+1 && isPortal(rune(line[x+1])) {
					if len(line) > x+2 && isPathway(rune(line[x+2])) {
						portals[point{x + 2, i}] = string(r) + string(rune(line[x+1]))
					} else {
						portals[point{x - 1, i}] = string(r) + string(rune(line[x+1]))
					}
				} else if len(donut) > i+1 && len(donut[i+1]) > x && isPortal(rune(donut[i+1][x])) {
					if len(donut) > i+2 && isPathway(rune(donut[i+2][x])) {
						portals[point{x, i + 2}] = string(r) + string(rune(donut[i+1][x]))
					} else {
						portals[point{x, i - 1}] = string(r) + string(rune(donut[i+1][x]))
					}
				}
			}

			if x > maxPoint.x {
				maxPoint.x = x
			}
		}

		maxPoint.y = i
	}

	portalLookup := make(map[string][]point)
	for k := range portals {
		portalLookup[portals[k]] = append(portalLookup[portals[k]], k)
	}

	return donut, portals, portalLookup, maxPoint, nil

}

func findShortestPath(maze []string, portals map[point]string, portalLookup map[string][]point, maxPoint point) int {
	start, finish := portalLookup["AA"][0], portalLookup["ZZ"][0]

	visited := make(map[point]bool)
	queue := []node{{start, 0}}

	for len(queue) > 0 {
		var head node
		head, queue = queue[0], queue[1:]

		if head.point == finish {
			return head.depth
		}

		visited[head.point] = true

		adjacent := []point{
			{head.point.x - 1, head.point.y},
			{head.point.x + 1, head.point.y},
			{head.point.x, head.point.y - 1},
			{head.point.x, head.point.y + 1},
		}

		for _, a := range adjacent {
			if visited[a] {
				continue
			}
			if a.x < 0 || a.x > maxPoint.x || a.y < 0 || a.y > maxPoint.y {
				continue
			}
			if !isPathway(rune(maze[a.y][a.x])) {
				continue
			}

			queue = append(queue, node{a, head.depth + 1})
		}

		if v, ok := portals[head.point]; ok {
			ps := portalLookup[v]
			for _, p := range ps {
				if p == head.point {
					continue
				}
				if visited[p] {
					continue
				}
				queue = append(queue, node{p, head.depth + 1})
			}
		}
	}

	return -1
}

func main() {
	maze, portals, portalLookup, maxPoint, err := readDonut("day20/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	sp := findShortestPath(maze, portals, portalLookup, maxPoint)
	log.Printf("Shortest path: %d\n", sp)
}
