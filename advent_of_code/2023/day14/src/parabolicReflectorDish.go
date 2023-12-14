package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Point struct {
	X, Y int
}

func (p Point) Add(direction Point) Point {
	return Point{X: p.X + direction.X, Y: p.Y + direction.Y}
}

var (
	North = Point{0, -1}
	East  = Point{1, 0}
	South = Point{0, 1}
	West  = Point{-1, 0}
)

func main() {
	grid, _ := readInput("2023/day14/src/input.txt")
	grid = runCycles(grid, 1_000_000_000)
	totalLoad := calculateTotalLoad(grid)
	fmt.Println("Total Load:", totalLoad)
}

func readInput(filePath string) (map[Point]rune, error) {
	file, _ := os.Open(filePath)
	defer file.Close()

	grid := make(map[Point]rune)
	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		for x, c := range scanner.Text() {
			grid[Point{X: x, Y: y}] = c
		}
		y++
	}
	return grid, scanner.Err()
}

func tilt(grid map[Point]rune, dir Point) {
	var points []Point
	for p := range grid {
		points = append(points, p)
	}

	sort.Slice(points, func(i, j int) bool {
		if dir.Y != 0 {
			return points[i].Y*dir.Y > points[j].Y*dir.Y
		}
		return points[i].X*dir.X > points[j].X*dir.X
	})

	for _, p := range points {
		if grid[p] == 'O' {
			nextP := p
			for {
				nextP = nextP.Add(dir)
				if grid[nextP] != '.' {
					break
				}
				grid[p], grid[nextP] = '.', 'O'
				p = nextP
			}
		}
	}
}

func runCycles(grid map[Point]rune, n int) map[Point]rune {
	seen := make(map[string]int)
	var remainingCycles int

	for cycle := 1; cycle <= n; cycle++ {
		gridStr := gridToString(grid)
		if firstSeenIndex, exists := seen[gridStr]; exists {
			cycleLength := cycle - firstSeenIndex
			remainingCycles = (n - cycle) % cycleLength
			break
		} else {
			seen[gridStr] = cycle
			tilt(grid, North)
			tilt(grid, West)
			tilt(grid, South)
			tilt(grid, East)
		}
	}
	for i := 0; i < remainingCycles; i++ {
		tilt(grid, North)
		tilt(grid, West)
		tilt(grid, South)
		tilt(grid, East)
	}
	return grid
}

func calculateTotalLoad(grid map[Point]rune) int {
	totalLoad := 0
	maxY := findMaxY(grid)
	for p, v := range grid {
		if v == 'O' {
			totalLoad += maxY - p.Y + 1
		}
	}
	return totalLoad
}

func findMaxY(grid map[Point]rune) int {
	maxY := 0
	for p := range grid {
		if p.Y > maxY {
			maxY = p.Y
		}
	}
	return maxY
}

func gridToString(grid map[Point]rune) string {
	maxX, maxY := findMaxXY(grid)
	var result string
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			result += string(grid[Point{X: x, Y: y}])
		}
		result += "\n"
	}
	return result
}

func findMaxXY(grid map[Point]rune) (int, int) {
	maxX, maxY := 0, 0
	for p := range grid {
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}
	return maxX, maxY
}
