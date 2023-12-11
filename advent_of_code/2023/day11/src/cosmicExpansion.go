package main

import (
	"bufio"
	"fmt"
	"os"
)

type position struct {
	row, col int
}

func main() {
	galaxies, _ := readGalaxies("2023/day11/src/input.txt")
	emptyRows, emptyCols := findEmptyRowsAndColumns(galaxies)
	totalDistance := sumOfExpandedDistances(galaxies, emptyRows, emptyCols, 1000000)
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

func findEmptyRowsAndColumns(galaxies []position) (map[int]bool, map[int]bool) {
	rows, cols := make(map[int]bool), make(map[int]bool)
	for _, g := range galaxies {
		rows[g.row] = true
		cols[g.col] = true
	}
	emptyRows, emptyCols := make(map[int]bool), make(map[int]bool)
	for r := 0; r <= getMaxRow(galaxies); r++ {
		if !rows[r] {
			emptyRows[r] = true
		}
	}
	for c := 0; c <= getMaxCol(galaxies); c++ {
		if !cols[c] {
			emptyCols[c] = true
		}
	}
	return emptyRows, emptyCols
}

func getMaxRow(galaxies []position) int {
	maxRow := 0
	for _, g := range galaxies {
		if g.row > maxRow {
			maxRow = g.row
		}
	}
	return maxRow
}

func getMaxCol(galaxies []position) int {
	maxCol := 0
	for _, g := range galaxies {
		if g.col > maxCol {
			maxCol = g.col
		}
	}
	return maxCol
}

func sumOfExpandedDistances(galaxies []position, emptyRows, emptyCols map[int]bool, factor int) int64 {
	var sum int64
	for i, g1 := range galaxies {
		for _, g2 := range galaxies[i+1:] {
			sum += expandedManhattanDistance(g1, g2, emptyRows, emptyCols, factor)
		}
	}
	return sum
}

func expandedManhattanDistance(p1, p2 position, emptyRows, emptyCols map[int]bool, factor int) int64 {
	rowDistance := int64(abs(p1.row - p2.row))
	colDistance := int64(abs(p1.col - p2.col))

	for r := min(p1.row, p2.row); r < max(p1.row, p2.row); r++ {
		if emptyRows[r] {
			rowDistance += int64(factor - 1)
		}
	}
	for c := min(p1.col, p2.col); c < max(p1.col, p2.col); c++ {
		if emptyCols[c] {
			colDistance += int64(factor - 1)
		}
	}
	return rowDistance + colDistance
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
