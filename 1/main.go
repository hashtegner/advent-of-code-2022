package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func sumElfCalories(scanner *bufio.Scanner) int {
	calories := 0

	for {
		text := scanner.Text()
		if len(text) == 0 {
			break
		}

		value, _ := strconv.ParseInt(text, 0, 64)
		calories += int(value)

		if !scanner.Scan() {
			break
		}
	}

	return calories
}

func main() {
	file, err := os.Open("./1/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	sums := make([]int, 0)
	for scanner.Scan() {
		elf_calories := sumElfCalories(scanner)
		sums = append(sums, elf_calories)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(sums)

	size := len(sums)
	max := sums[size-1]
	top_three := sums[size-1] + sums[size-2] + sums[size-3]

	fmt.Println("Max carried:", max)
	fmt.Println("Top three:", top_three)
}
