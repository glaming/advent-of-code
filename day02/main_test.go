package main

import (
	"testing"
)

type tapeTest struct {
	input, output tape
}

func TestExecuteIntcode(t *testing.T) {
	tt := []tapeTest{
		{tape{99, 50, 51, 52}, tape{99, 50, 51, 52}},
		{tape{1, 5, 6, 7, 99, 2, 3, 0}, tape{1, 5, 6, 7, 99, 2, 3, 5}},
		{tape{2, 5, 6, 7, 99, 2, 3, 0}, tape{2, 5, 6, 7, 99, 2, 3, 6}},
		{tape{2, 5, 6, 7, 99, 2, 3, 0}, tape{2, 5, 6, 7, 99, 2, 3, 6}},
	}

	// Adding in test cases from exercise
	tt = append(tt, []tapeTest{
		{tape{1, 0, 0, 0, 99}, tape{2, 0, 0, 0, 99}},
		{tape{2,3,0,3,99}, tape{2,3,0,6,99}},
		{tape{2,4,4,5,99,0}, tape{2,4,4,5,99,9801}},
		{tape{1,1,1,4,99,5,6,0,99}, tape{30,1,1,4,2,5,6,0,99}},
	}...)

	for i, test := range tt {
		output, err := executeIntcode(test.input)
		if err != nil {
			t.Errorf("test %d failed - error %s", i+1, err)
			continue
		}

		if output.String() != test.output.String() {
			t.Errorf("test %d failed - input %s, expected %s, got %s", i+1, test.input, test.output, output)
		}
	}

}
