package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type sequence struct {
	from int
	to   int
}

func (s sequence) Inside(o sequence) bool {
	return s.from >= o.from && s.to <= o.to
}

func (s sequence) Overlaps(o sequence) bool {
	return s.from >= o.from && s.from <= o.to ||
		s.to >= o.from && s.to <= o.to
}

type pair struct {
	first  sequence
	second sequence
}

func (p pair) FullyContained() bool {
	first := p.first
	second := p.second

	return first.Inside(second) || second.Inside(first)
}

func (p pair) Overlaps() bool {
	first := p.first
	second := p.second

	return first.Overlaps(second) || second.Overlaps(first)
}

func parseSequence(str string) sequence {
	splitted := strings.Split(str, "-")
	from, _ := strconv.Atoi(splitted[0])
	to, _ := strconv.Atoi(splitted[1])

	return sequence{from, to}
}

func parsePair(str string) pair {
	splitted := strings.Split(str, ",")
	first := parseSequence(splitted[0])
	second := parseSequence(splitted[1])

	return pair{first, second}
}

func main() {
	file, err := os.Open("./4/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	fullyContainedCount := 0
	overlapsCount := 0
	for scanner.Scan() {
		text := scanner.Text()
		p := parsePair(text)
		if p.FullyContained() {
			fullyContainedCount += 1
		}

		if p.Overlaps() {
			overlapsCount += 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Total Fully Contained:", fullyContainedCount)
	fmt.Println("Total Overlaps:", overlapsCount)
}
