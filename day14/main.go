package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"time"
)

type (
	ingredient struct {
		name     string
		quantity int
	}

	recipe struct {
		produces ingredient
		requires []ingredient
	}
)

func parseIngredient(s string) (ingredient, error) {
	var i ingredient
	n, err := fmt.Fscanf(strings.NewReader(s), "%d %s", &i.quantity, &i.name)
	if n != 2 {
		return i, errors.New("less than 2 values scanned when reading ingredient")
	}
	return i, err
}

func parseRecipe(s string) (recipe, error) {
	var r recipe

	split := strings.Split(s, "=>")

	i, err := parseIngredient(split[1])
	if err != nil {
		return r, err
	}
	r.produces = i

	split = strings.Split(split[0], ",")
	for _, s := range split {
		i, err = parseIngredient(s)
		if err != nil {
			return r, err
		}
		r.requires = append(r.requires, i)
	}

	return r, nil
}

func readRecipes(filename string) ([]recipe, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	var rs []recipe
	s := bufio.NewScanner(file)
	for s.Scan() {
		r, err := parseRecipe(s.Text())
		if err != nil {
			return nil, err
		}

		rs = append(rs, r)
	}

	return rs, s.Err()
}

// spare is used as accumulator throughout the recursion
func findOreRequired(toProduce ingredient, rs []recipe, spare map[string]int) (int, map[string]int) {
	if toProduce.name == "ORE" {
		return toProduce.quantity, spare
	}

	var r recipe
	for _, r = range rs {
		if r.produces.name == toProduce.name {
			break
		}
	}

	var oreUsed, quantity int

	if s, ok := spare[toProduce.name]; ok {
		quantity += s
		spare[toProduce.name] = 0
	}

	num := int(math.Ceil(float64(toProduce.quantity-quantity) / float64(r.produces.quantity)))
	if num > 0 {
		for _, i := range r.requires {
			var o int
			i.quantity *= num
			o, spare = findOreRequired(i, rs, spare)
			oreUsed += o
		}
		quantity += r.produces.quantity * num
	}

	spare[toProduce.name] = quantity - toProduce.quantity

	return oreUsed, spare
}

func main() {
	rs, err := readRecipes("day14/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	t := time.Now()

	spare := make(map[string]int)
	ore, _ := findOreRequired(ingredient{"FUEL", 5000}, rs, spare)

	d := time.Since(t)
	log.Printf("Time taken: %s\n", d)

	log.Printf("Ore required: %d\n", ore)
}
