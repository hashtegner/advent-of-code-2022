package main

import (
	"fmt"
	"testing"
)

func buildTree() *Tree {
	tree := NewTree()

	tree.PushDir("/")
	tree.AddFile("b.txt", 14848514)
	tree.AddFile("c.dat", 8504156)

	tree.PushDir("a")
	tree.AddFile("f", 29116)
	tree.AddFile("g", 2557)
	tree.AddFile("h.lst", 62596)

	tree.PushDir("e")
	tree.AddFile("i", 584)

	tree.PopDir()
	tree.PopDir()

	tree.PushDir("d")
	tree.AddFile("j", 4060174)
	tree.AddFile("d.log", 8033020)
	tree.AddFile("d.ext", 5626152)
	tree.AddFile("k", 7214296)

	return tree
}

func TestTree(t *testing.T) {
	tree := buildTree()

	fmt.Println(tree.dirSizes)

	tests := []struct {
		dir          string
		expectedSize int
	}{
		{"//ae", 584},
		{"/a", 94853},
		{"/d", 24933642},
		{"/", 48381165},
	}

	for _, test := range tests {
		result := tree.Size(test.dir)
		if result != test.expectedSize {
			t.Errorf("expected size %d for dir %s, got %d", test.expectedSize, test.dir, result)
		}
	}
}

func TestPartOne(t *testing.T) {
	tree := buildTree()

	result := tree.PartOne()
	expected := 95437
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func TestPartTwo(t *testing.T) {
	tree := buildTree()

	resultDir, resultSize := tree.PartTwo()
	expectedDir := "/d"
	expectedSize := 24_933_642
	if resultDir != expectedDir || resultSize != expectedSize {
		t.Errorf("expected dir %s=%d, got %s=%d", expectedDir, expectedSize, resultDir, resultSize)
	}
}
