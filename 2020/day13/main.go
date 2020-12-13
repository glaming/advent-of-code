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
			if currTs % b == 0 {
				fmt.Println("BusId * wait time:", b * (currTs-arrivalTs))
				os.Exit(0)
			}
		}
		currTs++
	}
}
