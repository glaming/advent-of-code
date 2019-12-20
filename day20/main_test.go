package main

import "testing"

func TestFindShortestPath(t *testing.T) {
	tt := []struct {
		filename string
		expected int
	}{
		{"input_test_1.txt", 23},
		{"input_test_2.txt", 58},
	}

	for _, test := range tt {
		a, b, c, d, err := readDonut(test.filename)
		if err != nil {
			t.Fatal(err)
		}
		sp := findShortestPath(a, b, c, d)

		if sp != test.expected {
			t.Errorf("For %s: expected %d, got %d", test.filename, test.expected, sp)
		}
	}
}
