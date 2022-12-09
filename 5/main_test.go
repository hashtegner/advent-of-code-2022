package main

import (
	"reflect"
	"testing"
)

func TestParseMove(t *testing.T) {
	tests := []struct {
		input    string
		expected move
	}{
		{"move 1 from 2 to 1", move{times: 1, from: 2, to: 1}},
		{"move 2 from 2 to 1", move{times: 2, from: 2, to: 1}},
		{"move 1 from 1 to 2", move{times: 1, from: 1, to: 2}},
	}

	for _, test := range tests {
		result := ParseMove(test.input)
		if !reflect.DeepEqual(test.expected, result) {
			t.Errorf("expected %v should result %v, got %v", test.input, test.expected, result)
		}
	}
}

func TestParseStacks(t *testing.T) {

	tests := []struct {
		input    []string
		expected []string
	}{
		{
			input: []string{
				"    [D]    ",
				"[N] [C]    ",
				"[Z] [M] [P]",
				" 1   2   3 ",
			},
			expected: []string{
				"ZN",
				"MCD",
				"P",
			},
		},

		{
			input: []string{
				"                [B]     [L]     [S]",
				"        [Q] [J] [C]     [W]     [F]",
				"    [F] [T] [B] [D]     [P]     [P]",
				"    [S] [J] [Z] [T]     [B] [C] [H]",
				"    [L] [H] [H] [Z] [G] [Z] [G] [R]",
				"[R] [H] [D] [R] [F] [C] [V] [Q] [T]",
				"[C] [J] [M] [G] [P] [H] [N] [J] [D]",
				"[H] [B] [R] [S] [R] [T] [S] [R] [L]",
				" 1   2   3   4   5   6   7   8   9 ",
			},
			expected: []string{
				"HCR",
				"BJHLSF",
				"RMDHJTQ",
				"SGRHZBJ",
				"RPFZTDCB",
				"THCG",
				"SNVZBPWL",
				"RJQGC",
				"LDTRHPFS",
			},
		},
	}

	for _, test := range tests {
		result := ParseStacks(test.input)

		if !reflect.DeepEqual(test.expected, result) {
			t.Errorf("expected %v should result %v, got %v", test.input, test.expected, result)
		}
	}
}

func TestMove(t *testing.T) {
	tests := []struct {
		input    []string
		move     move
		expected []string
	}{
		{
			[]string{
				"ZN",
				"MCD",
				"P",
			},
			move{times: 1, from: 2, to: 1},
			[]string{
				"ZND",
				"MC",
				"P",
			},
		},

		{
			[]string{
				"ZND",
				"MC",
				"P",
			},
			move{times: 3, from: 1, to: 3},
			[]string{
				"",
				"MC",
				"PDNZ",
			},
		},

		{
			[]string{
				"",
				"MC",
				"PDNZ",
			},
			move{times: 2, from: 2, to: 1},
			[]string{
				"CM",
				"",
				"PDNZ",
			},
		},

		{
			[]string{
				"CM",
				"",
				"PDNZ",
			},
			move{times: 1, from: 1, to: 2},
			[]string{
				"C",
				"M",
				"PDNZ",
			},
		},
	}

	for _, test := range tests {
		result := Move(test.input, test.move)

		if !reflect.DeepEqual(test.expected, result) {
			t.Errorf("expected %v should result %v, got %v", test.input, test.expected, result)
		}
	}
}

func TestMove9001(t *testing.T) {
	tests := []struct {
		input    []string
		move     move
		expected []string
	}{
		{
			input: []string{
				"ZN",
				"MCD",
				"P",
			},
			move: move{times: 1, from: 2, to: 1},
			expected: []string{
				"ZND",
				"MC",
				"P",
			},
		},

		{
			input: []string{
				"ZND",
				"MC",
				"P",
			},
			move: move{times: 3, from: 1, to: 3},
			expected: []string{
				"",
				"MC",
				"PZND",
			},
		},

		{
			[]string{
				"",
				"MC",
				"PZND",
			},
			move{times: 2, from: 2, to: 1},
			[]string{
				"MC",
				"",
				"PZND",
			},
		},

		{
			[]string{
				"MC",
				"",
				"PZND",
			},
			move{times: 1, from: 1, to: 2},
			[]string{
				"M",
				"C",
				"PZND",
			},
		},
	}

	for _, test := range tests {
		result := Move9001(test.input, test.move)

		if !reflect.DeepEqual(test.expected, result) {
			t.Errorf("expected %v should result %v, got %v", test.input, test.expected, result)
		}
	}
}

func TestReadMessage(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{
		{
			[]string{
				"ZN",
				"MCD",
				"P",
			},
			"NDP",
		},

		{
			[]string{
				"",
				"MC",
				"PDNZ",
			},
			"CZ",
		},
		{
			[]string{
				"C",
				"M",
				"PDNZ",
			},
			"CMZ",
		},

		{
			[]string{
				"",
				"",
				"",
			},
			"",
		},
		{
			[]string{},
			"",
		},
	}

	for _, test := range tests {
		result := ReadMessage(test.input)

		if !reflect.DeepEqual(test.expected, result) {
			t.Errorf("expected %v should result %v, got %v", test.input, test.expected, result)
		}
	}
}
