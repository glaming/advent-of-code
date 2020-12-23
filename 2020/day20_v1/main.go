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
		id, orientation     int
		flipped             bool
		data                []string
		edges, edgesReverse [4]string
		neighbours          [4]int
		connected           bool
		coords              point
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

func groupTiles(tm map[int]tile) ([]tile, []tile) {
	var connected, unconnected []tile
	for _, t := range tm {
		if t.connected {
			connected = append(connected, t)
		} else {
			unconnected = append(unconnected, t)
		}
	}
	return connected, unconnected
}

func getTile(p point, tm map[int]tile) tile {
	for _, t := range tm {
		if p.x == t.coords.x && p.y == t.coords.y {
			return t
		}
	}
	return tile{}
}

func main() {
	tiles, err := readTiles("2020/day20_v1/example.txt")
	if err != nil {
		log.Panic(err)
	}

	tiles[0].connected = true

	tileMap := make(map[int]tile, 0)
	for _, t := range tiles {
		tileMap[t.id] = t
	}

	for {
		connectedTiles, unconnectedTiles := groupTiles(tileMap)
		if len(unconnectedTiles) == 0 {
			break
		}

		for _, uct := range unconnectedTiles {
			matched := false
			for _, ct := range connectedTiles {
				for j, ctn := range ct.neighbours {
					// If neighbour already found
					if ctn != 0 {
						continue
					}

					for i, er := range uct.edgesReverse {
						if er != ct.edges[j] && er != ct.edgesReverse[j] {
							continue
						}

						if er == ct.edgesReverse[j] {
							uct.flipped = true
						}

						uct.connected = true
						uct.orientation = (i + 2) % 4
						uct.neighbours[i] = ct.id
						ct.neighbours[j] = uct.id

						resultingOrientation := (j + ct.orientation) % 4
						if ct.flipped {
							resultingOrientation = (resultingOrientation + 2) % 4
						}

						switch resultingOrientation {
						case 0:
							uct.coords = point{ct.coords.x, ct.coords.y + 1}
						case 1:
							uct.coords = point{ct.coords.x + 1, ct.coords.y}
						case 2:
							uct.coords = point{ct.coords.x, ct.coords.y - 1}
						case 3:
							uct.coords = point{ct.coords.x - 1, ct.coords.y}
						}

						matched = true

						tileMap[uct.id] = uct
						tileMap[ct.id] = ct

						break
					}

					if matched {
						break
					}
				}

				if matched {
					break
				}
			}
		}
	}

	minPoint := point{math.MaxInt64, math.MaxInt64}
	for _, t := range tileMap {
		if t.coords.x < minPoint.x {
			minPoint.x = t.coords.x
		}
		if t.coords.y < minPoint.y {
			minPoint.y = t.coords.y
		}
	}

	size := int(math.Sqrt(float64(len(tiles)))) - 1

	corners := []tile{
		getTile(minPoint, tileMap),
		getTile(point{minPoint.x + size, minPoint.y}, tileMap),
		getTile(point{minPoint.x, minPoint.y + size}, tileMap),
		getTile(point{minPoint.x + size, minPoint.y + size}, tileMap),
	}

	cornerIdsMultiplied := 1
	for _, c := range corners {
		cornerIdsMultiplied *= c.id
	}

	fmt.Println("Corner Ids multipled:", cornerIdsMultiplied)
}
