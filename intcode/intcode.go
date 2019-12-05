package intcode

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type (
	intcode int

	Tape []intcode
)


const (
	opcodeAdd  intcode = 1
	opcodeMul  intcode = 2
	opcodeHalt intcode = 99
)

func (t Tape) String() string {
	var ts []string
	for _, v := range t {
		s := strconv.Itoa(int(v))
		ts = append(ts, s)
	}

	return strings.Join(ts, ",")
}

func (t Tape) Get(index int) int {
	return int(t[index])
}

func ReadTape(filename string) (Tape, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(file)
	line, err := r.Read()
	if err != nil {
		return nil, err
	}

	var t Tape
	for _, v := range line {
		i, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}

		t = append(t, intcode(i))
	}

	return t, nil
}

func RestoreProgram(t Tape, noun, verb int) Tape {
	t[1] = intcode(noun)
	t[2] = intcode(verb)

	return t
}

// Assuming that the tape is valid
func Execute(t Tape) (Tape, error) {
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
