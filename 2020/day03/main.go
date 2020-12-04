package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type point struct {
	x, y int
}

func readMap(filename string) ([][]rune, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	var rss [][]rune

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rs := []rune(scanner.Text())
		rss = append(rss, rs)
	}

	return rss, nil
}

func traverse(rrs [][]rune, down, across int) int {
	pos := point{0,0 }

	trees := 0
	for ; pos.y < len(rrs) - 1; {
		pos.x = (pos.x + across) % len(rrs[0])
		pos.y = pos.y + down

		if rrs[pos.y][pos.x] == '#' {
			trees++
		}
	}
	return trees
}

func main() {
	rrs, err := readMap("2020/day03/input.txt")
	if err != nil {
		log.Panic(err)
	}

	trees := traverse(rrs, 1, 3)

	fmt.Println("Trees encountered:", trees)

	runs := []int{
		traverse(rrs, 1, 1),
		traverse(rrs, 1, 3),
		traverse(rrs, 1, 5),
		traverse(rrs, 1, 7),
		traverse(rrs, 2, 1),
	}

	acc := 1
	for _, r := range runs {
		acc = acc * r
	}

	fmt.Println("Trees multiplied:", acc)
}
