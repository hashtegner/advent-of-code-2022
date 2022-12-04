package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type empty struct{}

func uniqChars(s string) string {
	uniq := map[rune]empty{}

	for _, char := range s {
		uniq[char] = empty{}
	}

	str := strings.Builder{}
	for char := range uniq {
		str.WriteRune(char)
	}

	return str.String()
}

func duplicated(group []string) string {
	counts := map[rune]int{}

	size := len(group)
	for _, g := range group {
		uniq := uniqChars(g)
		for _, char := range uniq {
			counts[char] += 1
		}
	}

	str := strings.Builder{}
	for char, count := range counts {
		in_all := count == size
		if in_all {
			str.WriteRune(char)
		}
	}

	return str.String()
}

func isLower(a rune) bool {
	return a >= 'a' && a <= 'z'
}

func charNumber(a rune) int {
	if isLower(a) {
		return (int(a) - 'a') + 1
	}

	return (int(a) - 'A') + 27
}

func calculatePriorities(a string) int {
	sum := 0
	for _, char := range a {
		sum += charNumber(char)
	}

	return sum
}

func readGroup(scanner *bufio.Scanner) []string {
	group := []string{}

	for {
		group = append(group, scanner.Text())

		if len(group) == 3 || !scanner.Scan() {
			break
		}
	}

	return group
}

func main() {
	file, err := os.Open("./3-2/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	score := 0
	for scanner.Scan() {
		group := readGroup(scanner)
		dupl := duplicated(group)
		score += calculatePriorities(dupl)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Total score:", score)
}
