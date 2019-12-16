package main

import (
	"io/ioutil"
	"log"
	"os"
)

func runeToInt(r rune) int {
	if r >= '0' && r <= '9' {
		return int(r - '0')
	}
	return -1
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func readSignal(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var s []int
	for _, r := range string(b) {
		s = append(s, runeToInt(r))
	}

	return s, nil
}

func applyPhases(input, pattern []int, phases int) []int {
	output := make([]int, len(input))

	for phase := 0; phase < phases; phase++ {
		for i := 0; i < len(input); i++ {
			sum := 0
			for j, v := range input {
				index := ((j + 1) / (i + 1)) % 4
				sum += v * pattern[index]
			}
			output[i] = abs(sum % 10)
		}
		copy(input, output)
	}
	return output
}

func main() {
	signal, err := readSignal("day16/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	outputSignal := applyPhases(signal, []int{0, 1, 0, -1}, 100)

	log.Printf("Output signal after 100 phases %v\n", outputSignal)
}
