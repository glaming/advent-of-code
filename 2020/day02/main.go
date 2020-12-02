package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type passwordPolicyTest struct {
	lowerBound, upperBound int
	testChar rune
	password string
}

func (p passwordPolicyTest) isValid() bool {
	occurrences := 0
	for _, c := range p.password {
		if c == p.testChar {
			occurrences++
		}
	}

	if occurrences >= p.lowerBound && occurrences <= p.upperBound {
		return true
	}
	return false
}


func readPasswordPolicyTests(filename string) ([]passwordPolicyTest, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	var ppts []passwordPolicyTest

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		var ppt passwordPolicyTest
		_, err := fmt.Sscanf(line, "%d-%d %c: %s", &ppt.lowerBound, &ppt.upperBound, &ppt.testChar, &ppt.password)
		if err != nil {
			return nil, err
		}

		ppts = append(ppts, ppt)
	}

	return ppts, nil
}

func main() {
	ppts, err := readPasswordPolicyTests("2020/day02/input.txt")
	if err != nil {
		log.Panic(err)
	}

	numValid := 0
	for _, ppt := range ppts {
		if ppt.isValid() {
			numValid++
		}
	}

	fmt.Println("Valid solutions:", numValid)
}
