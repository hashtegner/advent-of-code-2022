package main

import "testing"

func TestDuplicated(t *testing.T) {
	tests := []struct {
		first    string
		second   string
		expected string
	}{
		{"vJrwpWtwJgWr", "hcsFMMfFFhFp", "p"},
		{"jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL", "L"},
		{"PmmdzqPrV", "vPwwTWBwg", "P"},
		{"AAAAA", "AAAAA", "A"},
		{"Ab", "Abb", "Ab"},
	}

	for _, test := range tests {
		result := duplicated(test.first, test.second)

		if result != test.expected {
			t.Fatalf("%s %s expected %s, got %s", test.first, test.second, test.expected, result)
		}
	}
}

func TestCalculatePriorities(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"a", 1},
		{"z", 26},
		{"A", 27},
		{"Z", 52},
		{"p", 16},
		{"L", 38},
		{"P", 42},
		{"v", 22},
		{"t", 20},
		{"pLPvts", 157},
	}

	for _, test := range tests {
		result := calculatePriorities(test.input)

		if result != test.expected {
			t.Fatalf("%s expected %v, got %v", test.input, test.expected, result)
		}
	}
}
