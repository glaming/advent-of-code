package main

import (
	"github.com/glaming/advent-of-code-2019/intcode"
	"log"
	"sync"
	"time"
)

func printOutput(out chan int) (f int) {
	for {
		select {
		case f = <-out:
			print(string(rune(f)))
		case <-time.After(100 * time.Millisecond):
			return f
		}
	}
}

func main() {
	t, err := intcode.ReadTape("day21/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	in, out := make(chan int), make(chan int)

	wg.Add(1)
	go func() {
		_, err = intcode.Execute(t, in, out)

		wg.Done()

		if err != nil {
			log.Fatal(err)
		}
	}()

	printOutput(out)

	inputProgram := []string{
		"NOT A T",
		"NOT B J",
		"OR T J",
		"NOT C T",
		"OR T J",
		"AND D J",
		"WALK",
	}

	for _, line := range inputProgram {
		for _, r := range line {
			in <- int(r)
		}
		in <- int('\n')
	}

	v := printOutput(out)
	log.Printf("Output %d\n", v)

	wg.Wait()

}
