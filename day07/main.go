package main

import (
	"fmt"
	"github.com/glaming/advent-of-code-2019/intcode"
	"log"
	"sync"
)

func runAmplifierSequence(tape intcode.Tape, phaseSettings []int) (int, error) {
	inputSignal := 0
	for amp := 0; amp < 5; amp++ {
		t := make(intcode.Tape, len(tape))
		copy(t, tape)

		var wg sync.WaitGroup
		in, out := make(chan int), make(chan int)

		wg.Add(1)
		go func() {
			in <- phaseSettings[amp]
			in <- inputSignal
			inputSignal = <-out
			wg.Done()
		}()

		_, err := intcode.Execute(t, in, out)
		if err != nil {
			return 0, fmt.Errorf("error running amplifer %d: %s", amp, err)
		}

		wg.Wait()
	}

	return inputSignal, nil
}

func runeToInt(r rune) int {
	if r >= '0' && r <= '9' {
		return int(r - '0')
	}
	return -1
}

// Looking to find the settings for maximal signal output
func findPhaseSettings(tape intcode.Tape) ([]int, int, error) {
	// Find valid phases to test
	var phasesToTest [][]int
loop:
	for i := 0; i <= 55555; i++ {
		var ps []int
		phaseStr := fmt.Sprintf("%05d", i)
		for _, p := range phaseStr {
			if p < '0' || p > '4' {
				continue loop
			}

			for _, p2 := range ps {
				if p2 == runeToInt(p) {
					continue loop
				}
			}
			ps = append(ps, runeToInt(p))
		}

		phasesToTest = append(phasesToTest, ps)
	}

	maxSignal := -1
	maxPhaseSetting := []int{0}

	for _, ps := range phasesToTest {
		signal, err := runAmplifierSequence(tape, ps)
		if err != nil {
			return nil, 0, err
		}

		if signal > maxSignal {
			maxSignal = signal
			maxPhaseSetting = ps
		}
	}

	return maxPhaseSetting, maxSignal, nil
}

func main() {
	tape, err := intcode.ReadTape("day07/input.txt")
	if err != nil {
		log.Panic(err)
	}

	ps, signal, err := findPhaseSettings(tape)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Phase settings: %v\nMax signal: %d", ps, signal)

}
