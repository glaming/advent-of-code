package intcode

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

type TapeTest struct {
	tapeInput, tapeExpected Tape
}

type TapeTestIO struct {
	tapeInput, tapeExpected Tape
	input                   io.Reader
	output                  io.ReadWriter
	outputExpected          string
}

func TestExecute(t *testing.T) {
	tt := []TapeTest{
		{Tape{99, 50, 51, 52}, Tape{99, 50, 51, 52}},
		{Tape{1, 5, 6, 7, 99, 2, 3, 0}, Tape{1, 5, 6, 7, 99, 2, 3, 5}},
		{Tape{2, 5, 6, 7, 99, 2, 3, 0}, Tape{2, 5, 6, 7, 99, 2, 3, 6}},
		{Tape{2, 5, 6, 7, 99, 2, 3, 0}, Tape{2, 5, 6, 7, 99, 2, 3, 6}},
	}

	// Adding in test cases from exercise
	tt = append(tt, []TapeTest{
		{Tape{1, 0, 0, 0, 99}, Tape{2, 0, 0, 0, 99}},
		{Tape{2, 3, 0, 3, 99}, Tape{2, 3, 0, 6, 99}},
		{Tape{2, 4, 4, 5, 99, 0}, Tape{2, 4, 4, 5, 99, 9801}},
		{Tape{1, 1, 1, 4, 99, 5, 6, 0, 99}, Tape{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}...)

	for i, test := range tt {
		output, err := Execute(test.tapeInput, &bytes.Buffer{}, &bytes.Buffer{})
		if err != nil {
			t.Errorf("test %d failed - error %s", i+1, err)
			continue
		}

		if output.String() != test.tapeExpected.String() {
			t.Errorf("test %d failed - tapeInput %s, expected %s, got %s", i+1, test.tapeInput, test.tapeExpected, output)
		}
	}

}

func TestExecute_InputOutput(t *testing.T) {
	tt := []TapeTestIO{
		{Tape{3, 0, 4, 0, 99}, Tape{52, 0, 4, 0, 99}, strings.NewReader("52"), &bytes.Buffer{}, "52"},
	}

	for i, test := range tt {
		output, err := Execute(test.tapeInput, test.input, test.output)
		if err != nil {
			t.Errorf("test %d failed - error %s", i+1, err)
			continue
		}

		if output.String() != test.tapeExpected.String() {
			t.Errorf("test %d failed - tapeInput %s, expected %s, got %s", i+1, test.tapeInput, test.tapeExpected, output)
		}

		givenOutput := make([]byte, 100)
		 n, err := test.output.Read(givenOutput)
		 if err != nil {
			t.Errorf("test %d failed - failed to read from output", i+1)
		 }
		 if string(givenOutput[:n]) != test.outputExpected {
			 t.Errorf("test %d failed - expected %s, got %s", i+1, test.outputExpected, string(givenOutput[:n]))
		 }
	}

}
