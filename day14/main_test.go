package main

import "testing"

func TestFindOreRequired(t *testing.T) {
	tt := []struct {
		file     string
		expected int
	}{
		{"input_test_1.txt", 31},
		{"input_test_2.txt", 165},
		{"input_test_5.txt", 2210736},
	}

	for _, test := range tt {
		rs, err := readRecipes(test.file)
		if err != nil {
			t.Fatal(err)
		}

		var oreRequired int
		spare := make(map[string]int)
		oreRequired, spare = findOreRequired(ingredient{"FUEL", 1}, rs, spare)

		if oreRequired != test.expected {
			t.Errorf("ore required expected %d, got %d", test.expected, oreRequired)
		}
	}
}

func TestFindMaxFuelFrom(t *testing.T) {
	tt := []struct {
		file     string
		expected int
	}{
		{"input_test_3.txt", 82892753},
	}

	for _, test := range tt {
		rs, err := readRecipes(test.file)
		if err != nil {
			t.Fatal(err)
		}

		fuel := findMaxFuelFrom(1000000000000, rs)

		if fuel != test.expected {
			t.Errorf("fuel produced expected %d, got %d", test.expected, fuel)
		}
	}
}
