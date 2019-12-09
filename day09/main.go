package main

import (
	"github.com/glaming/advent-of-code-2019/intcode"
	"log"
	"sync"
	"time"
)

func main() {
	t, err := intcode.ReadTape("day09/input.txt")
	if err != nil {
		log.Panic(err)
	}

	var wg sync.WaitGroup
	var outputs []int
	in, out := make(chan int), make(chan int)

	wg.Add(1)
	go func() {
		in <- 1

		for {
			select {
			case v := <-out:
				outputs = append(outputs, v)
			case <-time.After(2 * time.Second):
				wg.Done()
				return
			}
		}

		wg.Done()
	}()

	_, err = intcode.Execute(t, in, out)

	wg.Wait()

	if err != nil {
		log.Panic(err)
	}

	log.Printf("BOOST keycode: %+v\n", outputs)
}
