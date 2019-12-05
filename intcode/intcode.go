package intcode

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type (
	intcode int

	Tape []intcode
)

const (
	opcodeAdd    intcode = 1
	opcodeMul    intcode = 2
	opcodeInput  intcode = 3
	opcodeOutput intcode = 4
	opcodeHalt   intcode = 99
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
func Execute(t Tape, input io.Reader, output io.Writer) (Tape, error) {
	headPos := 0

loop:
	for {
		switch t[headPos] {
		case opcodeHalt:
			headPos = headPos + 1
			break loop
		case opcodeAdd:
			t[t[headPos+3]] = t[t[headPos+1]] + t[t[headPos+2]]
			headPos = headPos + 4
		case opcodeMul:
			t[t[headPos+3]] = t[t[headPos+1]] * t[t[headPos+2]]
			headPos = headPos + 4
		case opcodeInput:
			in := make([]byte, 100)
			n, err := input.Read(in)
			if err != nil {
				return nil, fmt.Errorf("got error on reading from tapeInput: %s", err)
			}
			val, err := strconv.Atoi(string(in[:n]))
			if err != nil {
				return nil, fmt.Errorf("failed to convert tapeInput to int: %s", err)
			}

			t[t[headPos+1]] = intcode(val)

			headPos = headPos + 2
		case opcodeOutput:
			val := int(t[t[headPos+1]])
			out := strconv.Itoa(val)
			output.Write([]byte(out))

			headPos = headPos + 2
		default:
			return nil, fmt.Errorf("unknown opcode %d at tape position %d", t[headPos], headPos)
		}
	}

	return t, nil
}
