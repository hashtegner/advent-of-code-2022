package main

import "testing"

func TestFullyContained(t *testing.T) {
	tests := []struct {
		pair     pair
		expected bool
	}{
		{pair{sequence{2, 4}, sequence{6, 8}}, false},
		{pair{sequence{2, 3}, sequence{4, 5}}, false},
		{pair{sequence{5, 7}, sequence{7, 9}}, false},
		{pair{sequence{2, 8}, sequence{3, 7}}, true},
		{pair{sequence{3, 7}, sequence{2, 8}}, true},
		{pair{sequence{6, 6}, sequence{4, 6}}, true},
		{pair{sequence{4, 6}, sequence{6, 6}}, true},
		{pair{sequence{2, 6}, sequence{4, 8}}, false},
	}

	for _, test := range tests {
		result := test.pair.FullyContained()
		if result != test.expected {
			t.Fatalf("%v expected %v, got %v", test.pair, test.expected, result)
		}
	}
}

func TestOverlaps(t *testing.T) {
	tests := []struct {
		pair     pair
		expected bool
	}{
		{pair{sequence{5, 7}, sequence{7, 9}}, true},
		{pair{sequence{2, 8}, sequence{3, 7}}, true},
		{pair{sequence{6, 6}, sequence{4, 6}}, true},
		{pair{sequence{2, 6}, sequence{4, 8}}, true},
		{pair{sequence{1, 9}, sequence{6, 6}}, true},
		{pair{sequence{1, 6}, sequence{2, 3}}, true},
		{pair{sequence{1, 8}, sequence{9, 9}}, false},
	}

	for _, test := range tests {
		result := test.pair.Overlaps()
		if result != test.expected {
			t.Fatalf("%v expected %v, got %v", test.pair, test.expected, result)
		}
	}
}
