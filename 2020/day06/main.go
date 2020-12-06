package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type form struct {
	people int
	ans [26]int
}

func (f form) countYeses() int {
	count := 0
	for _, v := range f.ans {
		if v == f.people {
			count++
		}
	}
	return count
}

func readForms(filename string) ([]form, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	var forms []form
	var f form

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			forms = append(forms, f)
			f = form{}
			continue
		}

		f.people++

		for _, c := range line {
			f.ans[c - 'a'] += 1
		}
	}
	forms = append(forms, f)

	return forms, nil
}

func main() {
	forms, err := readForms("2020/day06/input.txt")
	if err != nil {
		log.Panic(err)
	}

	count := 0
	for _, f := range forms {
		count += f.countYeses()
	}

	fmt.Println("Sum of yes counts", count)
}
