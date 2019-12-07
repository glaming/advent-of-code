package main

import (
	"bufio"
	"bytes"
	"github.com/glaming/advent-of-code-2019/intcode"
	"log"
	"strings"
)

func main() {
	tape, err := intcode.ReadTape("day05/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var output bytes.Buffer
	input := bufio.NewScanner(strings.NewReader("5"))
	tape, err = intcode.Execute(tape, input, &output)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("output:\n%s\n", output.String())
}
