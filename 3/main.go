package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type empty struct{}

func duplicated(a, b string) string {
	chars := map[rune]empty{}
	duplicated := map[rune]empty{}

	for _, char := range a {
		chars[char] = empty{}
	}

	for _, char := range b {
		_, exists := chars[char]
		if exists {
			duplicated[char] = empty{}
		}
	}

	str := strings.Builder{}
	for char := range duplicated {
		str.WriteRune(char)
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

func main() {
	file, err := os.Open("./3/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	score := 0
	for scanner.Scan() {
		text := scanner.Text()
		half := len(text) / 2

		first, second := text[0:half], text[half:]
		dupl := duplicated(first, second)

		score += calculatePriorities(dupl)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Total score:", score)
}
