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

loop:
	for y := 0; ; y++ {
		validInRow := 0
		foundValid := false
		for x := 0; ; x++ {
			in <- x
			in <- y
			output := <-out

			if output == 1 {
				if !foundValid {
					foundValid = true
				}
				validInRow++
			} else {
				// If previously found valid and isn't, means it's at end of row
				if foundValid {
					if validInRow >= 100 {
						valid := true
						for y2 := 0; y2 < 100; y2++ {
							in <- x - 99
							in <- y2 + y
							output := <-out
							if output == 0 {
								valid = false
								break
							}
						}
						if valid {
							log.Printf("100 x 100 starting from: (%d, %d)", x-99, y)

							in <- x
							in <- y + 99
							output := <-out
							print(output)

							break loop
						}
					}
					continue loop
				}
			}
		}
	}

}
