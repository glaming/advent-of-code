package intcode

import (
	"sync"
	"testing"
)

type TapeTest struct {
	tapeInput, tapeExpected Tape
}

type TapeTestIO struct {
	tapeInput, tapeExpected Tape
	input, outputExpected   int
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
		output, err := Execute(test.tapeInput, nil, nil)
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
		{Tape{3, 0, 4, 0, 99}, Tape{52, 0, 4, 0, 99}, 52, 52},
	}

	for i, test := range tt {
		var wg sync.WaitGroup
		var outputVal int
		in, out := make(chan int), make(chan int)

		wg.Add(1)
		go func() {
			in <- 52
			outputVal = <-out
			wg.Done()
		}()

		outputTape, err := Execute(test.tapeInput, in, out)

		wg.Wait()

		if err != nil {
			t.Errorf("test %d failed - error %s", i+1, err)
			continue
		}

		if outputTape.String() != test.tapeExpected.String() {
			t.Errorf("test %d failed - tapeInput %s, expected %s, got %s", i+1, test.tapeInput, test.tapeExpected, outputTape)
		}

		if outputVal != test.outputExpected {
			t.Errorf("test %d failed - expected %d, got %d", i+1, test.outputExpected, outputVal)
		}
	}
}

func TestExecute_ParameterModes(t *testing.T) {
	tt := []TapeTest{
		{Tape{1002, 4, 3, 4, 33}, Tape{1002, 4, 3, 4, 99}},
		{Tape{1101, 100, -1, 4, 0}, Tape{1101, 100, -1, 4, 99}},
	}

	for i, test := range tt {
		output, err := Execute(test.tapeInput, nil, nil)
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
		{Tape{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, Tape{}, 52, 0},
		{Tape{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, Tape{}, 8, 1},

		{Tape{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, Tape{}, 9, 0},
		{Tape{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, Tape{}, 8, 0},
		{Tape{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, Tape{}, 7, 1},

		{Tape{3, 3, 1108, -1, 8, 3, 4, 3, 99}, Tape{}, 52, 0},
		{Tape{3, 3, 1108, -1, 8, 3, 4, 3, 99}, Tape{}, 8, 1},

		{Tape{3, 3, 1107, -1, 8, 3, 4, 3, 99}, Tape{}, 9, 0},
		{Tape{3, 3, 1107, -1, 8, 3, 4, 3, 99}, Tape{}, 8, 0},
		{Tape{3, 3, 1107, -1, 8, 3, 4, 3, 99}, Tape{}, 7, 1},

		{Tape{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}, Tape{}, 0, 0},
		{Tape{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}, Tape{}, -1, 1},
		{Tape{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}, Tape{}, 1, 1},

		{Tape{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, Tape{}, 0, 0},
		{Tape{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, Tape{}, -1, 1},
		{Tape{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, Tape{}, 1, 1},

		{Tape{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
			1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
			999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}, Tape{}, 7, 999},
		{Tape{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
			1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
			999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}, Tape{}, 8, 1000},
		{Tape{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
			1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
			999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}, Tape{}, 9, 1001},
	}

	for i, test := range tt {
		var wg sync.WaitGroup
		var outputVal int
		in, out := make(chan int), make(chan int)

		wg.Add(1)
		go func() {
			in <- test.input
			outputVal = <-out
			wg.Done()
		}()

		_, err := Execute(test.tapeInput, in, out)

		wg.Wait()

		if err != nil {
			t.Errorf("test %d failed - error %s", i+1, err)
			continue
		}

		if outputVal != test.outputExpected {
			t.Errorf("test %d failed - expected %d, got %d", i+1, test.outputExpected, outputVal)
		}
	}
}
