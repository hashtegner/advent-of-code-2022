package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func IsLeftVisible(forest [][]int, currentRow, currentCol int) bool {
	currentHeight := forest[currentRow][currentCol]

	for col := 0; col < currentCol; col++ {
		height := forest[currentRow][col]

		if height >= currentHeight {
			return false
		}
	}

	return true
}

func IsRightVisible(forest [][]int, currentRow, currentCol int) bool {
	currentHeight := forest[currentRow][currentCol]

	for col := currentCol + 1; col < len(forest); col++ {
		height := forest[currentRow][col]

		if height >= currentHeight {
			return false
		}
	}

	return true
}

func IsTopVisible(forest [][]int, currentRow, currentCol int) bool {
	currentHeight := forest[currentRow][currentCol]

	for row := 0; row < currentRow; row++ {
		height := forest[row][currentCol]

		if height >= currentHeight {
			return false
		}
	}

	return true
}

func IsBottomVisible(forest [][]int, currentRow, currentCol int) bool {
	currentHeight := forest[currentRow][currentCol]

	for row := currentRow + 1; row < len(forest); row++ {
		height := forest[row][currentCol]

		if height >= currentHeight {
			return false
		}
	}

	return true
}

func IsVisible(forest [][]int, row, col int) bool {
	return IsLeftVisible(forest, row, col) ||
		IsRightVisible(forest, row, col) ||
		IsTopVisible(forest, row, col) ||
		IsBottomVisible(forest, row, col)
}

func ScenicScoreLeft(forest [][]int, currentRow, currentCol int) int {
	sum := 0
	currentHeight := forest[currentRow][currentCol]

	for col := currentCol - 1; col >= 0; col-- {
		height := forest[currentRow][col]

		sum += 1
		if height >= currentHeight {
			break
		}
	}

	return sum
}

func ScenicScoreRight(forest [][]int, currentRow, currentCol int) int {
	sum := 0
	currentHeight := forest[currentRow][currentCol]

	for col := currentCol + 1; col < len(forest); col++ {
		height := forest[currentRow][col]
		sum += 1

		if height >= currentHeight {
			break
		}
	}

	return sum
}

func ScenicScoreTop(forest [][]int, currentRow, currentCol int) int {
	sum := 0
	currentHeight := forest[currentRow][currentCol]

	for row := currentRow - 1; row >= 0; row-- {
		height := forest[row][currentCol]
		sum += 1
		if height >= currentHeight {
			break
		}
	}

	return sum
}

func ScenicScoreBottom(forest [][]int, currentRow, currentCol int) int {
	sum := 0
	currentHeight := forest[currentRow][currentCol]

	for row := currentRow + 1; row < len(forest); row++ {
		height := forest[row][currentCol]
		sum += 1
		if height >= currentHeight {
			break
		}
	}

	return sum
}

func ScenicScore(forest [][]int, row, col int) int {
	return ScenicScoreTop(forest, row, col) *
		ScenicScoreBottom(forest, row, col) *
		ScenicScoreLeft(forest, row, col) *
		ScenicScoreRight(forest, row, col)
}

func CountVisible(forest [][]int) int {
	visible := 0

	for row := range forest {
		for col := range forest[row] {
			if IsVisible(forest, row, col) {
				visible += 1
			}
		}
	}

	return visible
}

func MaxScenicScore(forest [][]int) int {
	max := 0

	for row := range forest {
		for col := range forest[row] {
			scenic := ScenicScore(forest, row, col)
			if scenic > max {
				max = scenic
			}
		}
	}

	return max
}

func parseRow(str string) []int {
	row := []int{}

	for _, char := range str {
		height, _ := strconv.Atoi(string(char))
		row = append(row, height)
	}

	return row
}

func main() {
	file, err := os.Open("./8/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	forest := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()
		forest = append(forest, parseRow(line))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Visible:", CountVisible(forest))
	fmt.Println("Max Scenic Score:", MaxScenicScore(forest))
}
