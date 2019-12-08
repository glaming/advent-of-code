package main

import (
	"fmt"
	"github.com/glaming/advent-of-code-2019/intcode"
	"log"
)

func main() {
	originalTape, err := intcode.ReadTape("day02/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	desiredOutput := 19690720

	// Try different nouns, verbs until we find the one where t[0] = desiredOutput
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			t := append(intcode.Tape(nil), originalTape...)

			t = intcode.RestoreProgram(t, noun, verb)
			t, err = intcode.Execute(t, nil, nil)
			if err != nil {
				log.Fatal(err)
			}

			if t.Get(0) == desiredOutput {
				fmt.Printf("found noun, verb: %d, %d\n", noun, verb)
				return
			}
		}
	}

	fmt.Println("Uh-oh... Didn't find an appropriate noun, verb")
}
