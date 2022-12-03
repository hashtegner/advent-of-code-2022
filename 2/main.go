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

	mustWin   bool
	mustDrawn bool
	mustLose  bool
}

func (s shape) isDrawn(b shape) bool {
	return s.name == b.name
}

func (s shape) win(b shape) bool {
	return s.winsFrom == b.name
}

func (s shape) determineMove(opponent shape) shape {
	if s.mustWin {
		return shapes[opponent.losesTo]
	}

	if s.mustLose {
		return shapes[opponent.winsFrom]
	}

	return shapes[opponent.name]
}

func (s shape) determineRoundPoints(opponent shape) int {
	if s.isDrawn(opponent) {
		return 3
	}

	if s.win(opponent) {
		return 6
	}

	return 0
}

func (s shape) roundPoints(opponent shape) int {
	return s.points + s.determineRoundPoints(opponent)
}

var shapes = map[string]shape{
	"rock":     {"rock", "X", "A", "paper", "scissors", 1, false, false, true},
	"paper":    {"paper", "Y", "B", "scissors", "rock", 2, false, true, false},
	"scissors": {"scissors", "Z", "C", "rock", "paper", 3, true, false, false},
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

func main() {
	file, err := os.Open("./2/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	score := 0
	determined_score := 0
	for scanner.Scan() {
		text := scanner.Text()
		opponent, me := text[0:1], text[2:3]
		opponentShape, meShape := opponentShapes[opponent], meShapes[me]
		score += meShape.roundPoints(opponentShape)
		determined_score += meShape.determineMove(opponentShape).roundPoints(opponentShape)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Total score:", score)
	fmt.Println("Determined Move Total score:", determined_score)
}
