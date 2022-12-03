package main

import "testing"

func TestDetermineMove(t *testing.T) {
	tests := []struct {
		opponent shape
		me       shape

		expected shape
	}{
		{shapes["rock"], meShapes["Y"], shapes["rock"]},         // should return in draw
		{shapes["paper"], meShapes["Y"], shapes["paper"]},       // should return in draw
		{shapes["scissors"], meShapes["Y"], shapes["scissors"]}, // should return in draw

		{shapes["paper"], meShapes["X"], shapes["rock"]},     // should return in lose
		{shapes["rock"], meShapes["X"], shapes["scissors"]},  // should return in lose
		{shapes["scissors"], meShapes["X"], shapes["paper"]}, // should return in lose

		{shapes["scissors"], meShapes["Z"], shapes["rock"]},  // should return in win
		{shapes["rock"], meShapes["Z"], shapes["paper"]},     // should return in win
		{shapes["paper"], meShapes["Z"], shapes["scissors"]}, // should return in win
	}

	for _, test := range tests {
		result := test.me.determineMove(test.opponent)
		if result.name != test.expected.name {
			t.Fatalf("opponent %v, me %v should => result in %v, but is %v", test.opponent, test.me, test.expected, result)
		}
	}

}

func TestCalculateScore(t *testing.T) {
	tests := []struct {
		opponent shape
		me       shape
		expected int
	}{
		{opponentShapes["A"], meShapes["Y"], 8},
		{opponentShapes["B"], meShapes["X"], 1},
		{opponentShapes["C"], meShapes["Z"], 6},
	}

	for _, test := range tests {
		result := test.me.roundPoints(test.opponent)
		if result != test.expected {
			t.Fatalf("opponent %v, me %v should result in %v, but is %v", test.opponent, test.me, test.expected, result)
		}
	}
}
