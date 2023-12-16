package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	X, Y int
}

type Direction struct {
	DX, DY int
}

var (
	Right = Direction{DX: 1, DY: 0}
	Left  = Direction{DX: -1, DY: 0}
	Up    = Direction{DX: 0, DY: -1}
	Down  = Direction{DX: 0, DY: 1}

	MOVES = map[Direction]map[rune][]Direction{
		Right: {
			'.':  {Right},
			'-':  {Right},
			'|':  {Up, Down},
			'\\': {Down},
			'/':  {Up},
		},
		Left: {
			'.':  {Left},
			'-':  {Left},
			'|':  {Up, Down},
			'\\': {Up},
			'/':  {Down},
		},
		Up: {
			'.':  {Up},
			'-':  {Left, Right},
			'|':  {Up},
			'\\': {Left},
			'/':  {Right},
		},
		Down: {
			'.':  {Down},
			'-':  {Left, Right},
			'|':  {Down},
			'\\': {Right},
			'/':  {Left},
		},
	}
)

func main() {
	grid, _ := readInput("2023/day16/src/input.txt")
	energizedMax := -1
	for y := range grid {
		count := simulateBeam(grid, Point{X: -1, Y: y}, Right)
		if count > energizedMax {
			energizedMax = count
		}
		count = simulateBeam(grid, Point{X: len(grid[0]), Y: y}, Left)
		if count > energizedMax {
			energizedMax = count
		}
	}
	for x := range grid[0] {
		count := simulateBeam(grid, Point{X: x, Y: -1}, Down)
		if count > energizedMax {
			energizedMax = count
		}

		count = simulateBeam(grid, Point{X: x, Y: len(grid)}, Up)
		if count > energizedMax {
			energizedMax = count
		}
	}
	fmt.Println("Maximum number of energized tiles:", energizedMax)
}

func readInput(filename string) ([][]rune, error) {
	file, _ := os.Open(filename)
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}
	return grid, scanner.Err()
}

func simulateBeam(grid [][]rune, startPosition Point, startDir Direction) int {
	energized := make(map[Point]map[Direction]bool)
	beams := []struct {
		Position Point
		Dir      Direction
	}{{Position: startPosition, Dir: startDir}}

	for len(beams) > 0 {
		beam := beams[len(beams)-1]
		beams = beams[:len(beams)-1]

		x, y := beam.Position.X+beam.Dir.DX, beam.Position.Y+beam.Dir.DY

		if x < 0 || x >= len(grid[0]) || y < 0 || y >= len(grid) {
			continue
		}
		if _, exists := energized[Point{x, y}]; !exists {
			energized[Point{x, y}] = make(map[Direction]bool)
		}
		if energized[Point{x, y}][beam.Dir] {
			continue
		}

		energized[Point{x, y}][beam.Dir] = true
		for _, newDir := range MOVES[beam.Dir][grid[y][x]] {
			beams = append(beams, struct {
				Position Point
				Dir      Direction
			}{Position: Point{x, y}, Dir: newDir})
		}
	}
	return len(energized)
}
