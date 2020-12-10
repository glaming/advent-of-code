package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	opcode string
	value  int
}

func readInstructions(filename string) ([]instruction, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	var ins []instruction

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")

		value, err := strconv.Atoi(split[1])
		if err != nil {
			return nil, err
		}

		ins = append(ins, instruction{split[0], value})
	}

	return ins, nil
}

func containsInt(is []int, i int) bool {
	for _, v := range is {
		if v == i {
			return true
		}
	}
	return false
}

func accValueBeforeLoop(ins []instruction) (finished bool, acc int) {
	acc = 0

	var visitedInsIndex []int

	var i int
	for i = 0; i < len(ins); {
		if i == len(ins) {
			return true, acc
		}
		if i > len(ins) {
			return false, acc
		}
		if containsInt(visitedInsIndex, i) {
			return false, acc
		}

		visitedInsIndex = append(visitedInsIndex, i)

		switch ins[i].opcode {
		case "nop":
		case "acc":
			acc += ins[i].value
		case "jmp":
			i += ins[i].value
			continue
		}

		i++
	}

	return i == len(ins), acc
}

func main() {
	ins, err := readInstructions("2020/day08/input.txt")
	if err != nil {
		log.Panic(err)
	}

	_, acc := accValueBeforeLoop(ins)
	fmt.Println("Accumulator:", acc)

	for i := range ins {
		insClone := make([]instruction, len(ins))
		copy(insClone, ins)

		switch insClone[i].opcode {
		case "jmp":
			insClone[i].opcode = "nop"
		case "nop":
			insClone[i].opcode = "jmp"
		case "acc":
			continue
		}

		finished, acc := accValueBeforeLoop(insClone)
		if !finished {
			continue
		}

		fmt.Println(fmt.Println("Accumulator:", acc))
		os.Exit(0)
	}

}
