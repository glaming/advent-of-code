package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/glaming/advent-of-code-2019/intcode"
	"strconv"
	"strings"
)

func runAmplifierSequence(tape intcode.Tape, phaseSettings []int) (int, error) {
	inputSignal := 0
	for amp := 0; amp < 5; amp++ {
		t := make(intcode.Tape, len(tape))
		copy(t, tape)

		input := strings.NewReader(fmt.Sprintf("%d\n%d\n", phaseSettings[amp], inputSignal))
		output := bytes.Buffer{}
		_, err := intcode.Execute(t, bufio.NewScanner(input), &output)
		if err != nil {
			return 0, fmt.Errorf("error running amplifer %d: %s", amp, err)
		}

		outputVal, err := strconv.Atoi(output.String())
		if err != nil {
			return 0, fmt.Errorf("error converting output to int for amp %d: %s", amp, err)

		}

		inputSignal = outputVal
	}

	return inputSignal, nil
}
