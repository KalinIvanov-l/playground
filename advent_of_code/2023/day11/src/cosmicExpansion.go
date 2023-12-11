package main

import (
	"bufio"
	"fmt"
	"os"
	_ "strings"
)

type position struct {
	row, col int
}

func main() {
	galaxies, _ := readGalaxies("2023/day11/src/input.txt")
	expandedGalaxies := expandGalaxies(galaxies)
	totalDistance := sumOfDistances(expandedGalaxies)
	fmt.Println("Total distance:", totalDistance)
}

func readGalaxies(filePath string) ([]position, error) {
	file, _ := os.Open(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var galaxies []position
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		for col, char := range line {
			if char == '#' {
				galaxies = append(galaxies, position{row, col})
			}
		}
		row++
	}
	return galaxies, nil
}

func expandGalaxies(galaxies []position) []position {
	rows, cols := make(map[int]bool), make(map[int]bool)
	for _, g := range galaxies {
		rows[g.row] = true
		cols[g.col] = true
	}
	for i, g := range galaxies {
		for r := 0; r < g.row; r++ {
			if !rows[r] {
				galaxies[i].row++
			}
		}
		for c := 0; c < g.col; c++ {
			if !cols[c] {
				galaxies[i].col++
			}
		}
	}
	return galaxies
}

func sumOfDistances(galaxies []position) int {
	var sum int
	for i, g1 := range galaxies {
		for _, g2 := range galaxies[i+1:] {
			sum += manhattanDistance(g1, g2)
		}
	}
	return sum
}

func manhattanDistance(p1, p2 position) int {
	return abs(p1.row-p2.row) + abs(p1.col-p2.col)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
