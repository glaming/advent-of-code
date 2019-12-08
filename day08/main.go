package main

import (
	"io/ioutil"
	"log"
	"sort"
)

func runeToInt(r rune) int {
	if r >= '0' && r <= '9' {
		return int(r - '0')
	}
	return -1
}

func readSpecialImageFormat(filename string, pxWidth, pxHeight int) ([][]int, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var image [][]int
	var layer []int

	for _, r := range string(b) {
		i := runeToInt(r)

		layer = append(layer, i)

		if len(layer) == pxHeight*pxWidth {
			image = append(image, layer)
			layer = []int{}
		}
	}

	return image, nil
}

func countInt(i int, is []int) int {
	count := 0
	for _, v := range is {
		if v == i {
			count++
		}
	}
	return count
}

func layerWithLeast(image [][]int, i int) []int {
	type layerCount struct {
		layer []int
		count int
	}

	var lcs []layerCount
	for _, l := range image {
		lcs = append(lcs, layerCount{l, countInt(i, l)})
	}

	sort.Slice(lcs, func(i, j int) bool {
		return lcs[i].count < lcs[j].count
	})

	return lcs[0].layer
}

func main() {
	image, err := readSpecialImageFormat("day08/input.txt", 25, 6)
	if err != nil {
		log.Panic(err)
	}

	l := layerWithLeast(image, 0)
	ones := countInt(1, l)
	twos := countInt(2, l)

	log.Printf("1s * 2s: %d", ones*twos)
}
