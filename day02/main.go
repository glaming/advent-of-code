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

func restoreProgram(t tape, noun, verb intcode) tape {
	t[1] = noun
	t[2] = verb

	return t
}

func main() {
	originalTape, err := readTape("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	desiredOutput := 19690720

	// Try different nouns, verbs until we find the one where t[0] = desiredOutput
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			t := append(tape(nil), originalTape...)

			t = restoreProgram(t, intcode(noun), intcode(verb))
			t, err = executeIntcode(t)
			if err != nil {
				log.Fatal(err)
			}

			if t[0] == intcode(desiredOutput) {
				fmt.Printf("found noun, verb: %d, %d\n", noun, verb)
				return
			}
		}
	}

	fmt.Println("Uh-oh... Didn't find an appropriate noun, verb")
}
