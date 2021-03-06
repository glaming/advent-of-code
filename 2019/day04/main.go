package main

import (
	"log"
	"strconv"
)

func numStrWithinRange(to, from int) []string {
	var numStrs []string
	for i := to; i <= from; i++ {
		s := strconv.Itoa(i)
		numStrs = append(numStrs, s)
	}
	return numStrs
}

func isNumStrValid(ns string) bool {
	// Are two adjacent digits the same
	for i := range ns {
		if i == len(ns)-1 {
			return false
		}
		if ns[i] == ns[i+1] {
			break
		}
	}

	// Do digits only increase
	for i := range ns {
		if i == len(ns)-1 {
			break
		}
		if ns[i] > ns[i+1] {
			return false
		}
	}

	// Are two digits not part of a larger group of digits
	var prevSeen uint8
	for i := range ns {
		if i == len(ns)-1 {
			return false
		}
		if ns[i] == ns[i+1] && ns[i] != prevSeen {
			if i+1 == len(ns)-1 {
				break
			}
			if ns[i] == ns[i+2] {
				prevSeen = ns[i]
				continue
			} else {
				break
			}
		}
	}

	return true
}

func main() {
	numStrs := numStrWithinRange(236491, 713787)

	var validNumStrs []string
	for _, ns := range numStrs {
		valid := isNumStrValid(ns)
		if valid {
			validNumStrs = append(validNumStrs, ns)
		}
	}

	log.Printf("Number of valid passwords: %d", len(validNumStrs))
}
