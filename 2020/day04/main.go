package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type passport struct {
	Byr, Iyr, Eyr, Cid int
	Hgt, Hcl, Ecl, Pid string
}

// All fields required apart from Cid
func (p passport) isValid() bool {
	if p.Byr < 1920 || p.Byr > 2002 { return false }
	if p.Iyr < 2010 || p.Iyr > 2020 { return false }
	if p.Eyr < 2020 || p.Eyr > 2030 { return false }

	var hgtVal int
	var hgtMeasure string
	_, err := fmt.Sscanf(p.Hgt, "%d%s", &hgtVal, &hgtMeasure)
	if err != nil {
		return false
	}
	if hgtMeasure != "cm" && hgtMeasure != "in" { return false }
	if hgtMeasure == "cm" && (hgtVal < 150 || hgtVal > 193) { return false }
	if hgtMeasure == "in" && (hgtVal < 59 || hgtVal > 76) { return false }

	if len(p.Hcl) != 7 { return false }
	for i, c := range p.Hcl {
		if i == 0 && c != '#' { return false }
		if i != 0 && !((c >= 'a' && c <= 'f') || (c >= '0' && c <= '9')) { return false }
	}

	switch p.Ecl {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
	default:
		return false
	}

	if len(p.Pid) != 9 { return false }
	for _, c := range p.Pid {
		if c < '0' || c > '9' { return false }
	}

	return true
}

func readPassports(filename string) ([]passport, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	var ps []passport
	var p passport

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			ps = append(ps, p)
			p = passport{}
			continue
		}

		parts := strings.Split(line, " ")
		for _, part := range parts {
			fieldval := strings.Split(part, ":")
			if len(fieldval) != 2 {
				return nil, fmt.Errorf("invalid field value pair %s", fieldval)
			}

			// Titling field as reflect can only set exported fields
			field, val := strings.Title(fieldval[0]), fieldval[1]

			switch field {
			case "Byr", "Iyr", "Eyr", "Cid":
				valInt, err := strconv.Atoi(val)
				if err != nil {
					continue
				}

				reflect.ValueOf(&p).Elem().FieldByName(field).SetInt(int64(valInt))
			case "Hgt", "Hcl", "Ecl", "Pid":
				reflect.ValueOf(&p).Elem().FieldByName(field).SetString(val)
			default:
				return nil, fmt.Errorf("unrecognised field %s", field)
			}
		}
	}

	ps = append(ps, p)

	return ps, nil
}

func main() {
	ps, err := readPassports("2020/day04/input.txt")
	if err != nil {
		log.Panic(err)
	}

	validCount := 0
	for _, p := range ps {
		if p.isValid() {
			validCount++
		}
	}

	fmt.Println("Valid passports:", validCount)
}
