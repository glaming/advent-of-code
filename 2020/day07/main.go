package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type rule struct {
	color string
	contains []string
}

func hasIntersection(as, bs []string) bool {
	for _, a := range as {
		for _, b := range bs {
			if a == b {
				return true
			}
		}
	}
	return false
}

func readRules(filename string) ([]rule, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	var rules []rule

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var r rule

		line := scanner.Text()
		split := strings.Split(line, " ")
		r.color = split[0] + " " + split[1]

		line = strings.Join(split[4:], " ")
		split = strings.Split(line, ", ")
		for _, s := range split {
			var ignoreInt int
			var ignoreStr string
			color := []string{"", ""}

			// "no other bags" will fail this scan
			_, err := fmt.Sscanf(s, "%d %s %s %s", &ignoreInt, &color[0], &color[1], &ignoreStr)
			if err != nil {
				continue
			}
			r.contains = append(r.contains, strings.Join(color, " "))
		}

		rules = append(rules, r)
	}

	return rules, nil
}

func main() {
	rules, err := readRules("2020/day07/input.txt")
	if err != nil {
		log.Panic(err)
	}

	containsColours := []string{"shiny gold"}
	for {
		setSize := len(containsColours)

		for _, r := range rules {
			// Is the current rule already known to contain
			alreadyPresent := hasIntersection(containsColours, []string{r.color})
			if alreadyPresent {
				continue
			}

			if hasIntersection(containsColours, r.contains) {
				containsColours = append(containsColours, r.color)
			}
		}

		// If result set remains unchanged
		if setSize == len(containsColours) {
			break
		}
	}

	fmt.Println("Number colours:", len(containsColours)-1)
}