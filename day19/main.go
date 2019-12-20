package main

import (
	"github.com/glaming/advent-of-code-2019/intcode"
	"log"
)

func main() {
	t, err := intcode.ReadTape("day19/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	in, out := make(chan int), make(chan int)

	go func() {
		for {
			tt := make(intcode.Tape, len(t))
			copy(tt, t)
			_, err = intcode.Execute(tt, in, out)
			if err != nil {
				log.Fatal(err)
			}
		}
	}()

	count := 0
	for y := 0; y < 50; y++ {
		for x := 0; x < 50; x++ {
			in <- x
			in <- y
			output := <-out
			count += output
		}
	}

	log.Printf("Points effeceted by beam: %d", count)

}
