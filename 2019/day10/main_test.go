package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestIsBetween(t *testing.T) {
	tt := []struct {
		a, b, p  point
		expected bool
	}{
		{point{0, 0}, point{2, 2}, point{1, 1}, true},
		{point{1, 1}, point{5, 3}, point{3, 2}, true},
		{point{1, 1}, point{5, 3}, point{3, 1}, false},
		{point{0, 0}, point{0, 5}, point{0, 2}, true},
		{point{0, 0}, point{0, 5}, point{1, 2}, false},
		{point{0, 0}, point{2, 2}, point{4, 4}, false},
	}

	for i, test := range tt {
		is := isBetween(test.a, test.b, test.p)
		if is != test.expected {
			t.Errorf("test %d failed - got %t, expected %t", i+1, is, test.expected)
		}
	}
}

func TestAsteroidsVisibleFrom(t *testing.T) {
	m := `
		.#..#
		.....
		#####
		....#
		...##
	`
	// Format map string... Remove tabs + extra spaces
	m = strings.Replace(strings.TrimSpace(m), "\t", "", -1)

	asteroids, err := getAsteroids(m)
	if err != nil {
		t.Fatal(err)
	}

	tt := []struct {
		from  point
		count int
	}{
		{point{1, 0}, 7},
		{point{4, 0}, 7},

		{point{0, 2}, 6},
		{point{1, 2}, 7},
		{point{2, 2}, 7},
		{point{3, 2}, 7},
		{point{4, 2}, 5},

		{point{4, 3}, 7},

		{point{3, 4}, 8},
		{point{4, 4}, 7},
	}

	for i, test := range tt {
		count := asteroidsVisibleFrom(test.from, asteroids)

		if count != test.count {
			t.Errorf("test %d failed - expected %d, got %d", i+1, test.count, count)
		}
	}
}

func TestVaporise(t *testing.T) {
	m := `
		.#..#
		.....
		#####
		....#
		...##
	`

	expected := []point{
		{3, 2},
		{4, 0},
		{4, 2},
		{4, 3},
		{4, 4},
		{0, 2},
		{1, 2},
		{2, 2},
		{1, 0},
	}

	// Format map string... Remove tabs + extra spaces
	m = strings.Replace(strings.TrimSpace(m), "\t", "", -1)
	as, _ := getAsteroids(m)
	p, _ := bestLocation(as)

	order := vaporise(p, as)

	if !reflect.DeepEqual(order, expected) {
		t.Error("test failed - out of order")
	}
}
