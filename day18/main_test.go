package main

import "testing"

func TestFindShortestPath(t *testing.T) {
	tt := []struct {
		filename string
		expected int
	}{
		{"input_test_1.txt", 8},
		{"input_test_2.txt", 86},
		{"input_test_3.txt", 132},
		{"input_test_4.txt", 136},
		{"input_test_5.txt", 81},
	}

	for _, test := range tt {
		m, err := readMap(test.filename)
		if err != nil {
			t.Fatal(err)
		}

		sp := findShortestPath(m)
		if sp != test.expected {
			t.Errorf("For %s, expected %d, got %d", test.filename, test.expected, sp)
		}
	}
}
