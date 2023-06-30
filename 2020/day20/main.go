package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type (
	point struct {
		x, y int
	}

	tile struct {
		id                  int
		data                []string
		edges, edgesReverse [4]string
	}
)

func (t *tile) makeEdges() {
	t.edges[0] = t.data[0]
	t.edgesReverse[0] = reverseString(t.data[0])

	t.edgesReverse[2] = t.data[9]
	t.edges[2] = reverseString(t.data[9])

	for i := 0; i < 10; i++ {
		t.edges[1] += string(t.data[i][9])
		t.edgesReverse[3] += string(t.data[i][0])
	}

	t.edgesReverse[1] = reverseString(t.edges[1])
	t.edges[3] = reverseString(t.edgesReverse[3])
}

// https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func readTiles(filename string) ([]tile, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	var tiles []tile
	var t tile

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			t.makeEdges()
			tiles = append(tiles, t)
			t = tile{}
			continue
		}

		if strings.HasPrefix(line, "Tile") {
			_, err := fmt.Sscanf(line, "Tile %d:", &t.id)
			if err != nil {
				return nil, err
			}
			continue
		}

		t.data = append(t.data, line)
	}

	t.makeEdges()
	tiles = append(tiles, t)

	return tiles, nil
}

// Assumes input is square...
func rotate(s []string) []string {
	size := len(s)
	out := make([]string, size)

	for j := 0; j < size; j++ {
		for i := size-1; i >= 0; i-- {
			out[j] += string(s[i][j])
		}
	}
	return out
}

func flip(s []string) []string {
	size := len(s)
	out := make([]string, size)

	for j := 0; j < size; j++ {
		out[size-j] = s[j]
	}
	return out
}

func findMatch(id int, edge string, ts []tile) tile {
	for _, t := range ts {
		// Don't match with itself...
		if t.id == id {
			continue
		}

		for i := 0; i < 4; i++ {
			if edge != t.edgesReverse[i] && edge != t.edges[i]{
				continue
			}

			rotationsRequired := (3 - i) % 4
			for r:=0; r < rotationsRequired; r++ {
				t.data = rotate(t.data)
			}
			if edge == t.edges[i] {
				t.data = flip(t.data)
			}

			t.makeEdges()

			return t
		}
	}

	return tile{}
}

func main() {
	tiles, err := readTiles("2020/day20/example.txt")
	if err != nil {
		log.Panic(err)
	}

	tileConnections := make(map[int][]int, 0)
	for _, t := range tiles {
		for _, tt := range tiles {
			if t.id == tt.id {
				continue
			}

			connects := false
			for _, e := range t.edges {
				for i := 0; i < 4; i++ {
					if e == tt.edges[i] || e == tt.edgesReverse[i] {
						connects = true
					}
				}
			}

			if connects {
				tileConnections[t.id] = append(tileConnections[t.id], tt.id)
			}
		}
	}

	var corners []tile
	for _, t := range tiles {
		if len(tileConnections[t.id]) == 2 {
			corners = append(corners, t)
		}
	}

	cornerIdsMultiplied := 1
	for _, c := range corners {
		cornerIdsMultiplied *= c.id
	}

	fmt.Println("Corner Ids multipled:", cornerIdsMultiplied)

	tileMap := make(map[point]tile, 0)
	tileMap[point{0, 0}] = corners[0]

	size := int(math.Sqrt(float64(len(tiles))))

	// TODO ORIENT THE CORNER BEFORE CONTINUING...
	// AND THEN REMAKE EDGES

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			// if already present
			if _, ok := tileMap[point{x, y}]; ok {
				continue
			}

			// if tile to left
			if t, ok := tileMap[point{x - 1, y}]; ok {
				matchingTile := findMatch(t.id, t.edges[1], tiles)
				tileMap[point{x, y}] = matchingTile
				continue
			}

			// tile above
			t := tileMap[point{x, y - 1}]
			matchingTile := findMatch(t.id, t.edges[2], tiles)
			tileMap[point{x, y}] = matchingTile
		}
	}

	fmt.Println("A")

}
