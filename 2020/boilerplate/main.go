package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readX(filename string) ([]interface{}, error) {
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
	_, err := readX("2020/XXX_DATE_XXX/input.txt")
	if err != nil {
		log.Panic(err)
	}
}
