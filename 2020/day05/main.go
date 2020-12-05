package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type (
	boardingpass struct {
		row, col int
	}

	boardingpasses []boardingpass
)

func (b boardingpass) seatID() int {
	return b.row * 8 + b.col
}

func (b boardingpasses) Len() int           { return len(b) }
func (b boardingpasses) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b boardingpasses) Less(i, j int) bool { return b[i].seatID() < b[j].seatID() }

func readBoardingPasses(filename string) ([]boardingpass, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	var bps []boardingpass

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var bp boardingpass
		line := scanner.Text()

		for _, c := range line[0:7] {
			if c == 'B' {
				bp.row = bp.row * 2 + 1
			} else {
				bp.row *= 2
			}
		}
		for _, c := range line[7:10] {
			if c == 'R' {
				bp.col = bp.col * 2 + 1
			} else {
				bp.col *= 2
			}
		}

		bps = append(bps, bp)
	}

	return bps, nil
}

func main() {
	bps, err := readBoardingPasses("2020/day05/input.txt")
	if err != nil {
		log.Panic(err)
	}

	maxSeatId := 0
	for _, bp := range bps {
		if bp.seatID() > maxSeatId {
			maxSeatId = bp.seatID()
		}
	}

	fmt.Println("Max seatID:", maxSeatId)

	sort.Sort(boardingpasses(bps))

	i := 1;
	for ; i < len(bps); i++ {
		if bps[i].seatID() != bps[i - 1].seatID() + 1 {
			break
		}
	}
	fmt.Println("Missing seatID", bps[i - 1].seatID() + 1)
}
