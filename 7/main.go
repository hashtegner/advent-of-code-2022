package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Tree struct {
	currentPath []string
	dirSizes    map[string]int
}

func NewTree() *Tree {
	return &Tree{
		currentPath: make([]string, 0),
		dirSizes:    make(map[string]int),
	}
}

func (t *Tree) AddFile(_name string, size int) {
	for _, dir := range t.currentPath {
		t.dirSizes[dir] = t.dirSizes[dir] + size
	}
}

func (t *Tree) PopDir() {
	t.currentPath = t.currentPath[:len(t.currentPath)-1]
}

func (t *Tree) PushDir(dir string) {
	path := strings.Join(t.currentPath, "") + dir
	t.currentPath = append(t.currentPath, path)
}

func (t *Tree) Size(dir string) int {
	return t.dirSizes[dir]
}

func (t *Tree) PartOne() int {
	sum := 0
	threeshold := 100_000

	for _, size := range t.dirSizes {
		if size < threeshold {
			sum += size
		}
	}

	return sum
}

func (t *Tree) PartTwo() (string, int) {
	threeshold := 30_000_000
	candidateDir := "/"
	max := t.Size(candidateDir)
	candidateSize := max
	freeSpace := 70_000_000 - max

	for dir, size := range t.dirSizes {
		if (size + freeSpace) < threeshold {
			continue
		}

		if candidateSize > size {
			candidateDir = dir
			candidateSize = size
		}
	}

	return candidateDir, candidateSize
}

func main() {
	file, err := os.Open("./7/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	tree := NewTree()
	r := regexp.MustCompile(`\d+ `)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "$ cd .." {
			tree.PopDir()
		} else if strings.HasPrefix(line, "$ cd") {
			tree.PushDir(strings.TrimPrefix(line, "$ cd "))
		} else if r.MatchString(line) {

			splitted := strings.Split(line, " ")
			sizeStr, name := splitted[0], splitted[1]
			size, _ := strconv.Atoi(sizeStr)
			tree.AddFile(name, size)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sum of sizes:", tree.PartOne())

	name, size := tree.PartTwo()
	fmt.Println("Candidate:", name, size)
}
