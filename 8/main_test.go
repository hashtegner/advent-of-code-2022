package main

import "testing"

func buildForest() [][]int {
	return [][]int{
		{3, 0, 3, 7, 3},
		{2, 5, 5, 1, 2},
		{6, 5, 3, 3, 2},
		{3, 3, 5, 4, 9},
		{3, 5, 3, 9, 0},
	}
}

func TestIsVisible(t *testing.T) {
	forest := buildForest()
	tests := []struct {
		row     int
		col     int
		visible bool
	}{
		{0, 0, true},
		{0, 1, true},
		{0, 4, true},
		{1, 0, true},
		{1, 4, true},
		{4, 0, true},
		{4, 1, true},
		{4, 4, true},

		{1, 1, true},  // 5
		{1, 2, true},  // 5
		{1, 3, false}, // 1

		{2, 1, true},  // 5
		{2, 2, false}, // 3
		{2, 3, true},  // 3

		{3, 1, false}, // 3
		{3, 2, true},  // 5
		{3, 3, false}, // 4
	}

	for _, test := range tests {
		result := IsVisible(forest, test.row, test.col)
		if result != test.visible {
			t.Errorf("coord %d %d expected %v, got %v", test.row, test.col, test.visible, result)
		}
	}
}

func TestCountVisible(t *testing.T) {
	forest := buildForest()

	result := CountVisible(forest)
	expected := 21
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestScenicScore(t *testing.T) {
	forest := buildForest()
	tests := []struct {
		description string
		expected    int
		result      int
	}{
		{"scoreTop(1,2) =>", 1, ScenicScoreTop(forest, 1, 2)},
		{"scoreBottom(1,2) =>", 2, ScenicScoreBottom(forest, 1, 2)},
		{"scoreLeft(1,2) =>", 1, ScenicScoreLeft(forest, 1, 2)},
		{"scoreRight(1,2) =>", 2, ScenicScoreRight(forest, 1, 2)},
		{"score(1,2) =>", 4, ScenicScore(forest, 1, 2)},

		{"scoreTop(3,2) =>", 2, ScenicScoreTop(forest, 3, 2)},
		{"scoreBottom(3,2) =>", 1, ScenicScoreBottom(forest, 3, 2)},
		{"scoreLeft(3,2) =>", 2, ScenicScoreLeft(forest, 3, 2)},
		{"scoreRight(3,2) =>", 2, ScenicScoreRight(forest, 3, 2)},
		{"score(3,2) =>", 8, ScenicScore(forest, 3, 2)},
	}

	for _, test := range tests {
		if test.result != test.expected {
			t.Errorf("%s expected %v, got %v", test.description, test.expected, test.result)
		}
	}
}

func TestMaxScenicScore(t *testing.T) {
	expected := 8
	result := MaxScenicScore(buildForest())
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
