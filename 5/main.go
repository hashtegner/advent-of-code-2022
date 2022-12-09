package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

const columnSize = 3
const columnSizeWithSpace = columnSize + 1

var movePattern = regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

type move struct {
	times int
	from  int
	to    int
}

func parseInt(str string) int {
	n, _ := strconv.Atoi(str)
	return n
}

func ParseMove(str string) move {
	match := movePattern.FindStringSubmatch(str)

	return move{
		times: parseInt(match[1]),
		from:  parseInt(match[2]),
		to:    parseInt(match[3]),
	}
}

func ParseStacks(strs []string) []string {
	row := strs[0]
	size := int(math.Round(float64(len(row)) / float64(columnSizeWithSpace)))
	stacks := make([]string, size)

	for _, str := range strs[:len(strs)-1] {
		charPos := 1
		for i := 0; i < size; i += 1 {
			char := string(str[charPos])
			if char != " " {
				stacks[i] = char + stacks[i]
			}

			charPos += columnSizeWithSpace
		}
	}

	return stacks
}

func Move(stacks []string, m move) []string {
	newStack := stacks

	for i := 0; i < m.times; i++ {
		fromIndex := m.from - 1
		toIndex := m.to - 1
		fromStack := stacks[fromIndex]
		toStack := stacks[toIndex]

		fromLastIndex := len(fromStack) - 1
		fromLastChar := string(fromStack[fromLastIndex])

		newStack[toIndex] = toStack + fromLastChar
		newStack[fromIndex] = fromStack[:fromLastIndex]
	}

	return newStack
}

func Move9001(stacks []string, m move) []string {
	newStack := stacks

	fromIndex := m.from - 1
	toIndex := m.to - 1

	fromStack := stacks[fromIndex]

	fromStartIndex := len(fromStack) - m.times
	chars := fromStack[fromStartIndex:]

	newStack[toIndex] += chars
	newStack[fromIndex] = newStack[fromIndex][:fromStartIndex]

	return newStack
}

func ReadMessage(stacks []string) string {
	message := ""
	for _, stack := range stacks {
		size := len(stack)
		if size == 0 {
			continue
		}

		char := string(stack[size-1])
		message += char
	}

	return message
}

func main() {
	file, err := os.Open("./5/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	stacks := []string{}
	stacks9001 := []string{}
	moving := false

	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			moving = true
			stacks9001 = ParseStacks(stacks)
			stacks = ParseStacks(stacks)

			continue
		}

		if !moving {
			stacks = append(stacks, text)
			continue
		}

		move := ParseMove(text)
		stacks = Move(stacks, move)
		stacks9001 = Move9001(stacks9001, move)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Default crate Message:", ReadMessage(stacks))
	fmt.Println("9001 crate Message:", ReadMessage(stacks9001))
}
