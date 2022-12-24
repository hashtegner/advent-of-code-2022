package main

import (
	"bufio"
	"log"
	"os"
)

type void struct{}

func allUniq(str string) bool {
	chars := map[rune]void{}

	for _, char := range str {
		if _, contains := chars[char]; contains {
			return false
		}

		chars[char] = void{}
	}

	return true
}

func FindMarker(str string, markerSize int) int {
	limit := len(str)
	for start := range str {
		final := start + markerSize
		if final > limit {
			break
		}

		sub := str[start:final]

		if allUniq(sub) {
			return final
		}
	}

	return 0
}

func main() {
	file, err := os.Open("./6/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	text := scanner.Text()
	marker4 := FindMarker(text, 4)
	marker14 := FindMarker(text, 14)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println("first marker of 4 after", marker4)
	log.Println("first marker of 14 after", marker14)
}
