package main

import "testing"

func TestFindMarker(t *testing.T) {
	tests := []struct {
		input    string
		size     int
		expected int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 4, 7},
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 14, 19},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 4, 5},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 14, 23},
		{"nppdvjthqldpwncqszvftbrmjlhg", 4, 6},
		{"nppdvjthqldpwncqszvftbrmjlhg", 14, 23},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 4, 10},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 14, 29},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 4, 11},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 14, 26},
	}

	for _, test := range tests {
		result := FindMarker(test.input, test.size)

		if result != test.expected {
			t.Errorf("expected %v should result %v, got %v", test.input, test.expected, result)
		}
	}
}
