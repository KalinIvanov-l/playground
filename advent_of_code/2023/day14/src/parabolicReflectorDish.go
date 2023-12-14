package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	grid, _ := readInput("2023/day14/src/input.txt")
	tiltNorth(grid)
	totalLoad := calculateTotalLoad(grid)
	fmt.Println("Total Load:", totalLoad)
}

func readInput(filePath string) ([][]rune, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}
	return grid, scanner.Err()
}

func tiltNorth(grid [][]rune) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == 'O' {
				newY := y
				for k := y - 1; k >= 0; k-- {
					if grid[k][x] == '.' {
						newY = k
					} else {
						break
					}
				}
				if newY != y {
					grid[newY][x], grid[y][x] = 'O', '.'
				}
			}
		}
	}
}

func calculateTotalLoad(grid [][]rune) int {
	totalLoad := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == 'O' {
				totalLoad += len(grid) - y
			}
		}
	}
	return totalLoad
}
