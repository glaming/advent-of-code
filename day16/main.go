package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
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

func applyPhasesOffset(input, pattern []int, phases int) []int {
	offsetStr := ""
	for _, v := range input[:7] {
		offsetStr += strconv.Itoa(v)
	}
	offset, _ := strconv.Atoi(offsetStr)

	// Trim input to offset data length
	output := input[offset:]

	for phase := 0; phase < phases; phase++ {
		sum := 0
		for i := len(output) - 1; i >= 0; i-- {
			sum += output[i]
			output[i] = sum % 10
		}
	}

	return output[:8]
}

func main() {
	signal, err := readSignal("day16/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	outputSignal := applyPhases(signal, []int{0, 1, 0, -1}, 100)

	log.Printf("Output signal after 100 phases %v\n", outputSignal)

	// Re-read input
	signal, err = readSignal("day16/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	realSignal := make([]int, 0)
	for i := 0; i < 10000; i++ {
		realSignal = append(realSignal, signal...)
	}

	outputSignal = applyPhasesOffset(realSignal, []int{0, 1, 0, -1}, 100)

	log.Printf("Output signal after 100 phases %v\n", outputSignal)
}
