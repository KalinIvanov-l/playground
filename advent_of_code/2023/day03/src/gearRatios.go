package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func isAdjacentToSymbol(grid [][]rune, startX, startY, endY int) bool {
	for x := startX - 1; x <= startX+1; x++ {
		for y := startY - 1; y <= endY+1; y++ {
			if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[x]) {
				if grid[x][y] != '.' && !unicode.IsDigit(grid[x][y]) {
					return true
				}
			}
		}
	}
	return false
}

func sumPartNumbers(grid [][]rune) int {
	sum := 0
	for x, row := range grid {
		for y := 0; y < len(row); {
			if unicode.IsDigit(row[y]) {
				start := y
				for y < len(row) && unicode.IsDigit(row[y]) {
					y++
				}
				number, _ := strconv.Atoi(string(row[start:y]))
				if isAdjacentToSymbol(grid, x, start, y-1) {
					sum += number
				}
			} else {
				y++
			}
		}
	}
	return sum
}

func main() {
	file, _ := os.Open("2023/day03/src/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var grid [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}
	sum := sumPartNumbers(grid)
	fmt.Println("Sum of all part numbers:", sum)
}
