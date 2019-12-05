package intcode

import (
	"testing"
)

type TapeTest struct {
	input, output Tape
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
		{Tape{2,3,0,3,99}, Tape{2,3,0,6,99}},
		{Tape{2,4,4,5,99,0}, Tape{2,4,4,5,99,9801}},
		{Tape{1,1,1,4,99,5,6,0,99}, Tape{30,1,1,4,2,5,6,0,99}},
	}...)

	for i, test := range tt {
		output, err := Execute(test.input)
		if err != nil {
			t.Errorf("test %d failed - error %s", i+1, err)
			continue
		}

		if output.String() != test.output.String() {
			t.Errorf("test %d failed - input %s, expected %s, got %s", i+1, test.input, test.output, output)
		}
	}

}
