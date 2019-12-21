package main

import (
	"bufio"
	"log"
	"os"
)

type (
	point struct {
		x, y, depth int
	}

	node struct {
		point
		steps int
	}
)

func isPortal(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func isPathway(r rune) bool {
	return r == '.'
}

func readDonut(filename string) ([]string, map[point]string, map[point]string, map[string][]point, point, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, nil, nil, point{}, err
	}

	var donut []string
	s := bufio.NewScanner(file)
	for s.Scan() {
		donut = append(donut, s.Text())
	}
	if s.Err() != nil {
		return nil, nil, nil, nil, point{}, s.Err()
	}

	var maxPoint point
	portals, portalSide := make(map[point]string), make(map[point]string)
	for i, line := range donut {
		for x, r := range line {
			if isPortal(r) {
				sideH, sideV := "", ""
				if x == 0 || len(line)-2 == x {
					sideH = "outer"
				} else {
					sideH = "inner"
				}
				if i == 0 || len(donut)-2 == i {
					sideV = "outer"
				} else {
					sideV = "inner"
				}

				if len(line) > x+1 && isPortal(rune(line[x+1])) {
					if len(line) > x+2 && isPathway(rune(line[x+2])) {
						portals[point{x + 2, i, 0}] = string(r) + string(rune(line[x+1]))
						portalSide[point{x + 2, i, 0}] = sideH
					} else {
						portals[point{x - 1, i, 0}] = string(r) + string(rune(line[x+1]))
						portalSide[point{x - 1, i, 0}] = sideH
					}
				} else if len(donut) > i+1 && len(donut[i+1]) > x && isPortal(rune(donut[i+1][x])) {
					if len(donut) > i+2 && isPathway(rune(donut[i+2][x])) {
						portals[point{x, i + 2, 0}] = string(r) + string(rune(donut[i+1][x]))
						portalSide[point{x, i + 2, 0}] = sideV
					} else {
						portals[point{x, i - 1, 0}] = string(r) + string(rune(donut[i+1][x]))
						portalSide[point{x, i - 1, 0}] = sideV
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

	return donut, portals, portalSide, portalLookup, maxPoint, nil

}

func findShortestPath(maze []string, portals, portalSide map[point]string, portalLookup map[string][]point, maxPoint point, check func(a, b node) bool) int {
	start, finish := node{portalLookup["AA"][0], 0}, node{portalLookup["ZZ"][0], 0}

	visited := make(map[point]bool)
	queue := []node{start}

	for len(queue) > 0 {
		var head node
		head, queue = queue[0], queue[1:]

		if check(head, finish) {
			return head.steps
		}

		visited[head.point] = true

		adjacent := []point{
			{head.point.x - 1, head.point.y, head.depth},
			{head.point.x + 1, head.point.y, head.depth},
			{head.point.x, head.point.y - 1, head.depth},
			{head.point.x, head.point.y + 1, head.depth},
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

			queue = append(queue, node{a, head.steps + 1})
		}

		if v, ok := portals[point{head.x, head.y, 0}]; ok {
			if portalSide[point{head.x, head.y, 0}] == "outer" && head.depth == 0 {
				continue
			}

			ps := portalLookup[v]
			for _, p := range ps {
				if portalSide[p] == "outer" {
					p.depth = head.depth + 1
				} else {
					p.depth = head.depth - 1
				}

				if p.x == head.point.x && p.y == head.point.y {
					continue
				}
				if visited[p] {
					continue
				}

				queue = append(queue, node{p, head.steps + 1})
			}
		}
	}

	return -1
}

func main() {
	maze, portals, portalSide, portalLookup, maxPoint, err := readDonut("day20/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	sp := findShortestPath(maze, portals, portalSide, portalLookup, maxPoint, func(a, b node) bool {
		if a.point == b.point {
			return true
		}
		return false
	})
	log.Printf("Shortest path: %d\n", sp)

	sp = findShortestPath(maze, portals, portalSide, portalLookup, maxPoint, func(a, b node) bool {
		if a.point == b.point && a.depth == 0 {
			return true
		}
		return false
	})
	log.Printf("Shortest path: %d\n", sp)

}
