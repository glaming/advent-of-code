package main

import (
	"github.com/glaming/advent-of-code-2019/intcode"
	"testing"
)

func TestRunAmplifierSequence(t *testing.T) {
	tt := []struct {
		tape           intcode.Tape
		phaseSettings  []int
		expectedSignal int
	}{
		{
			intcode.Tape{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0},
			[]int{4, 3, 2, 1, 0},
			43210,
		}, {
			intcode.Tape{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0},
			[]int{0, 1, 2, 3, 4},
			54321,
		}, {
			intcode.Tape{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33, 1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0},
			[]int{1, 0, 4, 3, 2},
			65210,
		},
	}

	for i, test := range tt {
		signal, err := runAmplifierSequence(test.tape, test.phaseSettings)
		if err != nil {
			t.Errorf("test %d, encountered error %s", i+1, err)
		}

		if signal != test.expectedSignal {
			t.Errorf("expected signal %d, got %d", test.expectedSignal, signal)
		}
	}
}
