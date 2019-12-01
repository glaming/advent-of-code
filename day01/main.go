package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type module struct {
	mass int64
}

func (m module) fuelRequiredToLaunch() int64 {
	return (m.mass / 3) - 2
}

func readModuleMasses(filename string) (ms []module, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, err
		}

		ms = append(ms, module{mass: i})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return ms, nil
}

func main() {
	modules, err := readModuleMasses("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var totalFuelRequried int64
	for _, m := range modules {
		totalFuelRequried = totalFuelRequried + m.fuelRequiredToLaunch()
	}

	log.Printf("Fuel required to launch: %d\n", totalFuelRequried)
}
