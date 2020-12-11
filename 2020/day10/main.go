package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func readInts(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	var ints []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		i, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}

		ints = append(ints, i)
	}

	return ints, nil
}

func main() {
	adapters, err := readInts("2020/day10/input.txt")
	if err != nil {
		log.Panic(err)
	}

	adapters = append(adapters, 0)
	sort.Ints(adapters)
	adapters = append(adapters, adapters[len(adapters)-1]+3)

	var counts [4]int
	for i := 0; i < len(adapters)-1; i++ {
		joltDifference := adapters[i+1] - adapters[i]
		counts[joltDifference]++
	}

	fmt.Println("Jolt differences:", counts[1]*counts[3])

	arrangements := make([]int, len(adapters))
	arrangements[0] = 1
	for i := 1; i < len(adapters); i++ {
		arrangements[i] = arrangements[i-1]
		for _, j := range []int{2,3} {
			if i-j < 0 {
				continue
			}

			if adapters[i] - adapters[i-j] <= 3 {
				arrangements[i] += arrangements[i-j]
			}
		}
	}

	fmt.Println("Arrangements:", arrangements[len(arrangements)-1])
}
