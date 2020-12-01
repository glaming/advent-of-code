package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readIntList(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	var nums []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}

		nums = append(nums, num)
	}

	return nums, nil
}

func main() {
	nums, err := readIntList("2020/day01/input.txt")
	if err != nil {
		log.Panic(err)
	}

	targetNumber := 2020
	for _, m := range nums {
		for _, n := range nums {
			for _, o := range nums {
				if m + n + o == targetNumber {
					fmt.Printf("Solution: %d", m * n * o)
					os.Exit(0)
				}
			}
		}
	}
}