package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readInts(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	var is []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		i, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}

		is = append(is, i)
	}

	return is, nil
}

func isValidInt(i int, prev []int) bool {
	for _, x := range prev {
		for _, y := range prev {
			if x + y == i && x != y {
				return true
			}
		}
	}
	return false
}

func findFirstInvalidInt(preambleLen int, is []int) int {
	for i := preambleLen; i < len(is); i++ {
		if !isValidInt(is[i], is[i-preambleLen:preambleLen+i]) {
			return is[i]
		}
	}
	return -1
}

func main() {
	is, err := readInts("2020/day09/input.txt")
	if err != nil {
		log.Panic(err)
	}

	invalid := findFirstInvalidInt(25, is)

	fmt.Println("Invalid number:", invalid)
}
