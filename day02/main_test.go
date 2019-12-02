package main

import (
	"testing"
)

func TestExecuteIntcode(t *testing.T) {
	tt := []struct {
		input, output tape
	}{
		{tape{99, 50, 51, 52}, tape{99, 50, 51, 52}},
		{tape{1, 5, 6, 7, 99, 2, 3, 0}, tape{1, 5, 6, 7, 99, 2, 3, 5}},
		{tape{2, 5, 6, 7, 99, 2, 3, 0}, tape{2, 5, 6, 7, 99, 2, 3, 6}},
		{tape{2, 5, 6, 7, 99, 2, 3, 0}, tape{2, 5, 6, 7, 99, 2, 3, 6}},
	}

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
