package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type (
	intcode int

	tape []intcode
)

const (
	opcodeAdd  intcode = 1
	opcodeMul  intcode = 2
	opcodeHalt intcode = 99
)

func (t tape) String() string {
	var ts []string
	for _, v := range t {
		s := strconv.Itoa(int(v))
		ts = append(ts, s)
	}

	return strings.Join(ts, ",")
}

func readTape(filename string) (tape, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(file)
	line, err := r.Read()
	if err != nil {
		return nil, err
	}

	var t tape
	for _, v := range line {
		i, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}

		t = append(t, intcode(i))
	}

	return t, nil
}

// Assuming that the tape is valid
func executeIntcode(t tape) (tape, error) {
	headPos := 0

loop:
	for {
		switch t[headPos] {
		case opcodeHalt:
			break loop
		case opcodeAdd:
			t[t[headPos+3]] = t[t[headPos+1]] + t[t[headPos+2]]
		case opcodeMul:
			t[t[headPos+3]] = t[t[headPos+1]] * t[t[headPos+2]]
		default:
			return nil, fmt.Errorf("unknown opcode %d at tape position %d", t[headPos], headPos)
		}

		headPos = headPos + 4
	}

	return t, nil
}

func restoreProgram(t tape) tape {
	t[1] = 12
	t[2] = 2

	return t
}

func main() {
	t, err := readTape("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	t = restoreProgram(t)

	t, err = executeIntcode(t)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("resulting tape:\n%s\n", t)
}
