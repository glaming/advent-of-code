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
	contains map[string]int
}

func getStrMapKeys(a map[string]int) []string {
	keys := make([]string, len(a))

	i := 0
	for k := range a {
		keys[i] = k
		i++
	}

	return keys
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
		r := rule{color: "", contains: make(map[string]int)}

		line := scanner.Text()
		split := strings.Split(line, " ")
		r.color = split[0] + " " + split[1]

		line = strings.Join(split[4:], " ")
		split = strings.Split(line, ", ")
		for _, s := range split {
			var ignoreStr string
			var quantity int
			color := []string{"", ""}

			// "no other bags" will fail this scan
			_, err := fmt.Sscanf(s, "%d %s %s %s", &quantity, &color[0], &color[1], &ignoreStr)
			if err != nil {
				continue
			}

			r.contains[strings.Join(color, " ")] = quantity
		}

		rules = append(rules, r)
	}

	return rules, nil
}

func requiredBags(colour string, rules []rule) int {
	var colorRule rule
	for _, r := range rules {
		if r.color == colour {
			colorRule = r
			break
		}
	}

	count := 0
	for c, q := range colorRule.contains {
		count += q * requiredBags(c, rules) + q
	}

	return count
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

			contains := getStrMapKeys(r.contains)
			if hasIntersection(containsColours, contains) {
				containsColours = append(containsColours, r.color)
			}
		}

		// If result set remains unchanged
		if setSize == len(containsColours) {
			break
		}
	}

	fmt.Println("Part 1")
	fmt.Println("Number colours:", len(containsColours)-1)


	numRequired := requiredBags("shiny gold", rules)
	fmt.Println("Part 2")
	fmt.Println("Required bags:", numRequired)


}