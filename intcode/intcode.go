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
	intcode       int
	opcode        int
	parameterMode int

	Tape []intcode
)

const (
	opcodeAdd    opcode = 1
	opcodeMul    opcode = 2
	opcodeInput  opcode = 3
	opcodeOutput opcode = 4
	opcodeJmpIfT opcode = 5
	opcodeJmpIfF opcode = 6
	opcodeLT     opcode = 7
	opcodeEq     opcode = 8
	opcodeHalt   opcode = 99

	paramModePosition  parameterMode = 0
	paramModeImmediate parameterMode = 1
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

func (t Tape) getIntcode(index int) intcode {
	if int(index) >= len(t) || index < 0 {
		return -1
	}
	return t[index]
}

func (t Tape) value(p intcode, pm parameterMode) (intcode, error) {
	switch pm {
	case paramModePosition:
		return t.getIntcode(int(p)), nil
	case paramModeImmediate:
		return p, nil
	default:
		return -1, fmt.Errorf("unknown parameter mode %d", p)
	}
}

func (i intcode) parse() (opcode, []parameterMode) {
	iVal := int(i)

	opVal := iVal % 100
	pm1Val := ((iVal - opVal) / 100) % 2
	pm2Val := ((iVal - pm1Val - opVal) / 1000) % 2
	pm3Val := ((iVal - pm2Val - pm1Val - opVal) / 10000) % 2

	return opcode(opVal), []parameterMode{parameterMode(pm1Val), parameterMode(pm2Val), parameterMode(pm3Val)}
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
		opcode, pms := t[headPos].parse()

		var params []intcode
		for i, pm := range pms {
			ic, err := t.value(t.getIntcode(headPos+1+i), pm)
			if err != nil {
				return nil, err
			}
			params = append(params, ic)
		}

		switch opcode {
		case opcodeHalt:
			headPos = headPos + 1
			break loop
		case opcodeAdd:
			t[t[headPos+3]] = params[0] + params[1]
			headPos = headPos + 4
		case opcodeMul:
			t[t[headPos+3]] = params[0] * params[1]
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
			val := int(params[0])
			out := strconv.Itoa(val)
			output.Write([]byte(out))

			headPos = headPos + 2
		case opcodeJmpIfT:
			if params[0] != 0 {
				headPos = int(params[1])
			} else {
				headPos = headPos + 3
			}
		case opcodeJmpIfF:
			if params[0] == 0 {
				headPos = int(params[1])
			} else {
				headPos = headPos + 3

			}
		case opcodeLT:
			if params[0] < params[1] {
				t[t[headPos+3]] = 1
			} else {
				t[t[headPos+3]] = 0
			}

			headPos = headPos + 4
		case opcodeEq:
			if params[0] == params[1] {
				t[t[headPos+3]] = 1
			} else {
				t[t[headPos+3]] = 0
			}

			headPos = headPos + 4
		default:
			return nil, fmt.Errorf("unknown opcode %d at tape position %d", t[headPos], headPos)
		}
	}

	return t, nil
}
