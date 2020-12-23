package main

import (
	"bufio"
	"fmt"
	"log"
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

func main() {
	tiles, err := readTiles("2020/day20/input.txt")
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
}
