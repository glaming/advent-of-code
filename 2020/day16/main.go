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
		label    string
		ranges   [2]validRange
		position int
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

		r := rule{position: -1}
		split := strings.Split(line, ": ")
		r.label = split[0]
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

func getRulesValidForPosition(rules []rule, validTickets [][]int) map[int][]rule {
	validForPosition := make(map[int][]rule, 0)

	for i := range validTickets[0] {
		var validRules []rule

		for _, r := range rules {
			validRule := true
			for _, ticket := range validTickets {
				if !r.isValueValid(ticket[i]) {
					validRule = false
					break
				}
			}
			if validRule {
				validRules = append(validRules, r)
			}
		}

		validForPosition[i] = validRules
	}

	return validForPosition
}

func main() {
	rules, yourTicket, nearbyTickets, err := readInput("2020/day16/input.txt")
	if err != nil {
		log.Panic(err)
	}

	var errorRate int
	var validTickets [][]int
	for _, ticket := range nearbyTickets {
		ticketValid := true
		for _, v := range ticket {
			valueValid := false

			for _, r := range rules {
				if r.isValueValid(v) {
					valueValid = true
					break
				}
			}

			if !valueValid {
				ticketValid = false
				errorRate += v
				break
			}
		}

		if ticketValid {
			validTickets = append(validTickets, ticket)
		}
	}

	fmt.Println("Ticket scanning error rate:", errorRate)

	rulesValidForPosition := getRulesValidForPosition(rules, validTickets)
	for {
		newRulesMatched := false

		// What positions only have 1 rule available
		for i, rs := range rulesValidForPosition {
			if len(rs) != 1 || rs[0].position != -1 {
				continue
			}

			rs[0].position = i
			newRulesMatched = true

			// remove this rule from all other positions...
			for j, rsj := range rulesValidForPosition {
				if i == j {
					continue
				}
				for k, r := range rsj {
					if r.label == rs[0].label {
						rulesValidForPosition[j] = append(rsj[0:k], rsj[k+1:]...)
						break
					}
				}
			}
		}

		if !newRulesMatched {
			break
		}
	}

	// Only 1 value now for each position
	multipedVal := 1
	for _, rs := range rulesValidForPosition {
		if strings.HasPrefix(rs[0].label, "departure") {
			multipedVal *= yourTicket[rs[0].position]
		}
	}

	fmt.Println("Multipled val", multipedVal)
}
