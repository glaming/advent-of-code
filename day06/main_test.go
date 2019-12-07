package main

import "testing"

func TestSumOrbits(t *testing.T) {
	orbits := orbits{
		"COM": []object{"B"},
		"B":   []object{"C", "G"},
		"C":   []object{"D"},
		"D":   []object{"E", "I"},
		"E":   []object{"F", "J"},
		"G":   []object{"H"},
		"J":   []object{"K"},
		"K":   []object{"L"},
	}

	total := sumOrbits(orbits, COM, 0)
	if total != 42 {
		t.Errorf("expected 42, got %d", total)
	}
}
