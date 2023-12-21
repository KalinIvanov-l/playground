package main

import (
	"bufio"
	"fmt"
	"os"
)

type Position struct {
	X, Y int
}

func main() {
	grid, start := loadGrid("2023/day21/src/input.txt")
	part1Result := part1(grid, start)
	fmt.Println("Part 1 Result:", part1Result)

	part2Result := part2(grid, start)
	fmt.Println("Part 2 Result:", part2Result)
}

func loadGrid(filePath string) ([][]rune, Position) {
	file, _ := os.Open(filePath)
	defer file.Close()

	var grid [][]rune
	var start Position
	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		var row []rune
		line := scanner.Text()
		for x, ch := range line {
			row = append(row, ch)
			if ch == 'S' {
				start = Position{X: x, Y: y}
			}
		}
		grid = append(grid, row)
		y++
	}
	return grid, start
}

func part1(grid [][]rune, start Position) int {
	elves := make(map[Position]bool)
	elves[start] = true

	for i := 0; i < 64; i++ {
		newElves := make(map[Position]bool)
		for e := range elves {
			for _, d := range []Position{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
				next := Position{X: e.X + d.X, Y: e.Y + d.Y}
				if isValidPosition(grid, next) && grid[next.Y][next.X] != '#' {
					newElves[next] = true
				}
			}
		}
		elves = newElves
	}
	return len(elves)
}

func part2(grid [][]rune, start Position) int {
	const LastStep = 26501365
	n := len(grid)
	m := len(grid[0])

	elves := make(map[Position]bool)
	elves[start] = true

	values := make([]int, LastStep+1)
	values[0] = 1

	for step := 1; step <= LastStep; step++ {
		newElves := make(map[Position]bool)
		for e := range elves {
			for _, d := range []Position{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
				next := Position{X: (e.X + d.X + m) % m, Y: (e.Y + d.Y + n) % n}
				if grid[next.Y][next.X] != '#' {
					newElves[next] = true
				}
			}
		}
		values[step] = len(newElves)
		elves = newElves
	}
	return values[LastStep]
}

func isValidPosition(grid [][]rune, p Position) bool {
	return p.X >= 0 && p.X < len(grid[0]) && p.Y >= 0 && p.Y < len(grid)
}
