package intcode

import (
	"encoding/csv"
	"fmt"
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
	opcodeRelAdj opcode = 9
	opcodeHalt   opcode = 99

	paramModePosition  parameterMode = 0
	paramModeImmediate parameterMode = 1
	paramModeRelative  parameterMode = 2
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

func (t Tape) value(p intcode, pm parameterMode, relativeBase int) (intcode, error) {
	switch pm {
	case paramModePosition:
		return t.getIntcode(int(p)), nil
	case paramModeImmediate:
		return p, nil
	case paramModeRelative:
		return t.getIntcode(int(p) + relativeBase), nil
	default:
		return -1, fmt.Errorf("unknown parameter mode %d", p)
	}
}

// This is a hack to get the length to be satisfactory
func (t Tape) set(pm parameterMode, rb int, i, p intcode) Tape {
	currentLen := len(t)
	index := int(i)
	if index > currentLen-1 {
		zeros := make([]intcode, index-currentLen+1)
		t = append(t, zeros...)
	}
	switch pm {
	case paramModePosition:
		t[index] = p
	case paramModeImmediate:
		panic("setting tape value with immediate positioning - this shouldn't have happened")
	case paramModeRelative:
		t[index+rb] = p
	}
	return t
}

func runeToInt(r rune) int {
	if r >= '0' && r <= '9' {
		return int(r - '0')
	}
	return -1
}

func (i intcode) parse() (opcode, []parameterMode) {
	var pms [3]parameterMode

	iStr := fmt.Sprintf("%05d", int(i))
	for index, r := range iStr[:3] {
		pms[index] = parameterMode(runeToInt(r))
	}

	opVal := int(i) % 100
	return opcode(opVal), []parameterMode{pms[2], pms[1], pms[0]}
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
// TODO: Should really turn Tape into a struct at this point
func Execute(t Tape, input chan int, output chan int) (Tape, error) {
	headPos := 0
	relativeBase := 0

loop:
	for {
		opcode, pms := t[headPos].parse()

		var params []intcode
		for i, pm := range pms {
			ic, err := t.value(t.getIntcode(headPos+1+i), pm, relativeBase)
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
			t = t.set(pms[2], relativeBase, t[headPos+3], params[0]+params[1])
			headPos = headPos + 4
		case opcodeMul:
			t = t.set(pms[2], relativeBase, t[headPos+3], params[0]*params[1])
			headPos = headPos + 4
		case opcodeInput:
			val := <-input
			t = t.set(pms[0], relativeBase, t[headPos+1], intcode(val))
			headPos = headPos + 2
		case opcodeOutput:
			val := int(params[0])
			output <- val
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
				t = t.set(pms[2], relativeBase, t[headPos+3], 1)
			} else {
				t = t.set(pms[2], relativeBase, t[headPos+3], 0)
			}

			headPos = headPos + 4
		case opcodeEq:
			if params[0] == params[1] {
				t = t.set(pms[2], relativeBase, t[headPos+3], 1)
			} else {
				t = t.set(pms[2], relativeBase, t[headPos+3], 0)
			}

			headPos = headPos + 4
		case opcodeRelAdj:
			relativeBase += int(params[0])
			headPos = headPos + 2
		default:
			return nil, fmt.Errorf("unknown opcode %d at tape position %d", t[headPos], headPos)
		}
	}

	return t, nil
}
