package main

import (
	"fmt"
	"github.com/glaming/advent-of-code-2019/intcode"
	"log"
	"sync"
)

func runAmplifierSequence(tape intcode.Tape, phaseSettings []int) (int, error) {
	inputSignal := 0

	var wg sync.WaitGroup

	var chs []chan int
	for i := 0; i < 5; i++ {
		chs = append(chs, make(chan int))
	}
	chs = append(chs, chs[0])

	go func() {
		for i := 0; i < 5; i++ {
			chs[i] <- phaseSettings[i]
		}
		chs[0] <- 0
	}()

	// Only adding 4 as want last amplifier to be active for final read
	// If included in above go func then the feedback loop won't be read correctly
	wg.Add(4)
	for amp := 0; amp < 5; amp++ {
		t := make(intcode.Tape, len(tape))
		copy(t, tape)

		go func(i int) {
			// Not handling error...
			_, _ = intcode.Execute(t, chs[i], chs[i+1])
			wg.Done()
		}(amp)
	}

	wg.Wait()

	// Adding so WaitGroup doesn't panic when count < 0
	wg.Add(1)

	inputSignal = <-chs[5]

	return inputSignal, nil
}

func runeToInt(r rune) int {
	if r >= '0' && r <= '9' {
		return int(r - '0')
	}
	return -1
}

// Looking to find the settings for maximal signal output
func findPhaseSettings(tape intcode.Tape, psStart, psEnd rune) ([]int, int, error) {
	// Find valid phases to test
	var phasesToTest [][]int
loop:
	for i := 0; i <= 99999; i++ {
		var ps []int
		phaseStr := fmt.Sprintf("%05d", i)
		for _, p := range phaseStr {
			if p < psStart || p > psEnd {
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

	ps, signal, err := findPhaseSettings(tape, '0', '4')
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Phase settings: %v\nMax signal: %d", ps, signal)

	psFbl, signalFbl, err := findPhaseSettings(tape, '5', '9')
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Phase settings: %v\nMax signal: %d", psFbl, signalFbl)
}
