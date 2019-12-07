package main

import (
	"bufio"
	"log"
	"os"
)

type (
	object string

	orbits map[object][]object
)

const COM = object("COM")

func readOrbits(filename string) (orbits, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	orbits := make(orbits)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		obj := object(line[0:3])
		orb := object(line[4:7])

		orbits[obj] = append(orbits[obj], orb)
	}
	if scanner.Err() != nil {
		return nil, err
	}

	return orbits, nil
}

func sumOrbits(orbits orbits, from object, depth int) int {
	if _, ok := orbits[from]; !ok {
		return depth
	}

	total := depth
	for _, obj := range orbits[from] {
		total = total + sumOrbits(orbits, obj, depth+1)
	}
	return total
}

func main() {
	orbits, err := readOrbits("day06/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	totalOrbits := sumOrbits(orbits, COM, 0)

	log.Printf("total number orbits: %d\n", totalOrbits)
}
