package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type instruction struct {
	t string
	memory int
	bitClear int
	bitSet int
}

func readProgram(filename string) ([]interface{}, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	return nil, nil
}

func main() {
	_, err := readProgram("2020/day14/input.txt")
	if err != nil {
		log.Panic(err)
	}
}
