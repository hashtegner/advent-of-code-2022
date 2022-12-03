package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type shape struct {
	name     string
	me       string
	opponent string

	losesTo  string
	winsFrom string

	points int
}

func (s shape) isDrawn(b shape) bool {
	return s.name == b.name
}

func (s shape) win(b shape) bool {
	return s.winsFrom == b.name
}

func (s shape) loses(b shape) bool {
	return !s.isDrawn(b) && !s.win(b)
}

func (s shape) roundPoints(b shape) int {
	if s.isDrawn(b) {
		return 3
	}

	if s.win(b) {
		return 6
	}

	return 0
}

var shapes = map[string]shape{
	"rock":     {"rock", "X", "A", "paper", "scissors", 1},
	"paper":    {"paper", "Y", "B", "scissors", "rock", 2},
	"scissors": {"scissors", "Z", "C", "rock", "paper", 3},
}

var meShapes = map[string]shape{
	"X": shapes["rock"],
	"Y": shapes["paper"],
	"Z": shapes["scissors"],
}

var opponentShapes = map[string]shape{
	"A": shapes["rock"],
	"B": shapes["paper"],
	"C": shapes["scissors"],
}

func calculateScore(opponent, me shape) int {
	return me.points + me.roundPoints(opponent)
}

func main() {
	file, err := os.Open("./2/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	score := 0
	for scanner.Scan() {
		text := scanner.Text()
		opponent, me := text[0:1], text[2:3]
		opponentShape, meShape := opponentShapes[opponent], meShapes[me]

		score += calculateScore(opponentShape, meShape)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Total score:", score)
}
