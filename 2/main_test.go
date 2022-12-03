package main

import "testing"

func TestCalculateScore(t *testing.T) {
	tests := []struct {
		opponent string
		me       string
		expected int
	}{
		{"A", "Y", 8},
		{"B", "X", 1},
		{"C", "Z", 6},
	}

	for _, test := range tests {
		result := calculateScore(opponentShapes[test.opponent], meShapes[test.me])
		if result != test.expected {
			t.Fatalf("%s%s should result in %v, but is %v", test.opponent, test.me, test.expected, result)
		}
	}
}
