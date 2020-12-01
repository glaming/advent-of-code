package main

import "testing"

var (
	check = func(a, b node) bool {
		if a.point == b.point {
			return true
		}
		return false
	}

	checkDepth = func(a, b node) bool {
		if a.point == b.point && a.depth == 0 {
			return true
		}
		return false
	}
)

func TestFindShortestPath(t *testing.T) {
	tt := []struct {
		filename string
		expected int
		check    func(a, b node) bool
	}{
		{"input_test_1.txt", 26, check},
		{"input_test_2.txt", -1, check},
		{"input_test_3.txt", 396, checkDepth},
	}

	for _, test := range tt {
		a, b, c, d, e, err := readDonut(test.filename)
		if err != nil {
			t.Fatal(err)
		}
		sp := findShortestPath(a, b, c, d, e, test.check)

		if sp != test.expected {
			t.Errorf("For %s: expected %d, got %d", test.filename, test.expected, sp)
		}
	}
}
