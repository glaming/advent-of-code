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
	output                  bytes.Buffer
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
		{Tape{3, 0, 4, 0, 99}, Tape{52, 0, 4, 0, 99}, strings.NewReader("52"), bytes.Buffer{}, "52"},
	}

	for i, test := range tt {
		output, err := Execute(test.tapeInput, test.input, &test.output)
		if err != nil {
			t.Errorf("test %d failed - error %s", i+1, err)
			continue
		}

		if output.String() != test.tapeExpected.String() {
			t.Errorf("test %d failed - tapeInput %s, expected %s, got %s", i+1, test.tapeInput, test.tapeExpected, output)
		}

		if test.output.String() != test.outputExpected {
			t.Errorf("test %d failed - expected %s, got %s", i+1, test.outputExpected, test.output.String())
		}
	}
}

func TestExecute_ParameterModes(t *testing.T) {
	tt := []TapeTest{
		{Tape{1002, 4, 3, 4, 33}, Tape{1002, 4, 3, 4, 99}},
		{Tape{1101, 100, -1, 4, 0}, Tape{1101, 100, -1, 4, 99}},
	}

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

func TestExecute_JumpAndConditionals(t *testing.T) {
	tt := []TapeTestIO{
		{Tape{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, Tape{}, strings.NewReader("52"), bytes.Buffer{}, "0"},
		{Tape{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, Tape{}, strings.NewReader("8"), bytes.Buffer{}, "1"},

		{Tape{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, Tape{}, strings.NewReader("9"), bytes.Buffer{}, "0"},
		{Tape{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, Tape{}, strings.NewReader("8"), bytes.Buffer{}, "0"},
		{Tape{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, Tape{}, strings.NewReader("7"), bytes.Buffer{}, "1"},

		{Tape{3, 3, 1108, -1, 8, 3, 4, 3, 99}, Tape{}, strings.NewReader("52"), bytes.Buffer{}, "0"},
		{Tape{3, 3, 1108, -1, 8, 3, 4, 3, 99}, Tape{}, strings.NewReader("8"), bytes.Buffer{}, "1"},

		{Tape{3, 3, 1107, -1, 8, 3, 4, 3, 99}, Tape{}, strings.NewReader("9"), bytes.Buffer{}, "0"},
		{Tape{3, 3, 1107, -1, 8, 3, 4, 3, 99}, Tape{}, strings.NewReader("8"), bytes.Buffer{}, "0"},
		{Tape{3, 3, 1107, -1, 8, 3, 4, 3, 99}, Tape{}, strings.NewReader("7"), bytes.Buffer{}, "1"},

		{Tape{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}, Tape{}, strings.NewReader("0"), bytes.Buffer{}, "0"},
		{Tape{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}, Tape{}, strings.NewReader("-1"), bytes.Buffer{}, "1"},
		{Tape{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}, Tape{}, strings.NewReader("1"), bytes.Buffer{}, "1"},

		{Tape{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, Tape{}, strings.NewReader("0"), bytes.Buffer{}, "0"},
		{Tape{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, Tape{}, strings.NewReader("-1"), bytes.Buffer{}, "1"},
		{Tape{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, Tape{}, strings.NewReader("1"), bytes.Buffer{}, "1"},

		{Tape{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
			1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
			999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}, Tape{}, strings.NewReader("7"), bytes.Buffer{}, "999"},
		{Tape{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
			1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
			999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}, Tape{}, strings.NewReader("8"), bytes.Buffer{}, "1000"},
		{Tape{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
			1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
			999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}, Tape{}, strings.NewReader("9"), bytes.Buffer{}, "1001"},
	}

	for i, test := range tt {
		_, err := Execute(test.tapeInput, test.input, &test.output)
		if err != nil {
			t.Errorf("test %d failed - error %s", i+1, err)
			continue
		}

		if test.output.String() != test.outputExpected {
			t.Errorf("test %d failed - expected %s, got %s", i+1, test.outputExpected, test.output.String())
		}
	}
}
