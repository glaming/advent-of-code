package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput(filename string) (int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	timestamp, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return 0, nil, err
	}

	var buses []int
	scanner.Scan()
	split := strings.Split(scanner.Text(), ",")
	for _, s := range split {
		if s == "x" {
			buses = append(buses, -1)
			continue
		}
		busId, err := strconv.Atoi(s)
		if err != nil {
			return 0, nil, err
		}

		buses = append(buses, busId)
	}

	return timestamp, buses, nil
}

func main() {
	arrivalTs, buses, err := readInput("2020/day13/input.txt")
	if err != nil {
		log.Panic(err)
	}

	currTs := arrivalTs
	for {
		for _, b := range buses {
			if b == -1 {
				continue
			}
			if currTs % b == 0 {
				fmt.Println("BusId * wait time:", b * (currTs-arrivalTs))
				goto Part2
			}
		}
		currTs++
	}

Part2:
	time, interval := 0, 1
	for i, b := range buses {
		if b == -1 {
			continue
		}

		for {
			if (time + i) % b == 0 {
				break
			}
			time += interval
		}
		interval *= b
	}

	fmt.Println("Minimum time:", time)
}