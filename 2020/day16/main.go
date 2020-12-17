package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type (
	// range is a keyword...
	validRange struct {
		min, max int
	}

	rule struct {
		label  string
		ranges [2]validRange
	}
)

func parseIntStr(s string) ([]int, error) {
	split := strings.Split(s, ",")

	var ns []int
	for _, x := range split {
		n, err := strconv.Atoi(x)
		if err != nil {
			return nil, err
		}
		ns = append(ns, n)
	}

	return ns, nil
}

func (r rule) isValueValid(v int) bool {
	for _, vr := range r.ranges {
		if vr.min <= v && vr.max >= v {
			return true
		}
	}
	return false
}

func readInput(filename string) ([]rule, []int, [][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, nil, err
	}

	scanner := bufio.NewScanner(file)

	// Read in rules
	var rules []rule
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		var r rule
		split := strings.Split(line, ": ")
		_, err := fmt.Sscanf(split[1], "%d-%d or %d-%d", &r.ranges[0].min, &r.ranges[0].max, &r.ranges[1].min, &r.ranges[1].max)
		if err != nil {
			return nil, nil, nil, err
		}

		rules = append(rules, r)
	}

	// Read "your ticket"
	scanner.Scan()
	scanner.Scan()
	line := scanner.Text()
	yourTicket, err := parseIntStr(line)
	if err != nil {
		return nil, nil, nil, err
	}

	// Read "nearby tickets"
	var nearbyTickets [][]int
	scanner.Scan()
	scanner.Scan()
	for scanner.Scan() {
		line := scanner.Text()
		ticket, err := parseIntStr(line)
		if err != nil {
			return nil, nil, nil, err
		}

		nearbyTickets = append(nearbyTickets, ticket)
	}

	return rules, yourTicket, nearbyTickets, nil
}

func main() {
	rules, _, nearbyTickets, err := readInput("2020/day16/input.txt")
	if err != nil {
		log.Panic(err)
	}

	var errorRate int
	for _, ticket := range nearbyTickets {
		for _, v := range ticket {
			valueValid := false

			for _, r := range rules {
				if r.isValueValid(v) {
					valueValid = true
					break
				}
			}

			if !valueValid {
				errorRate += v
			}
		}
	}

	fmt.Println("Ticket scanning error rate:", errorRate)
}
