package main

import (
	"github.com/glaming/advent-of-code-2019/intcode"
	"log"
	"sync"
)

func main() {
	tape, err := intcode.ReadTape("day05/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	var outputVal int
	in, out := make(chan int), make(chan int)

	wg.Add(1)
	go func() {
		in <- 5
		outputVal = <-out
		wg.Done()
	}()

	tape, err = intcode.Execute(tape, in, out)

	wg.Wait()

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("output:\n%d\n", outputVal)
}
