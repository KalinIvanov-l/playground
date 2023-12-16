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
	result := simulateBeam(grid)
	fmt.Println("Number of energized tiles:", result)
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

func simulateBeam(grid [][]rune) int {
	energized := make(map[Point]map[Direction]bool)
	beams := []struct {
		Position Point
		Dir      Direction
	}{{Position: Point{X: -1, Y: 0}, Dir: Right}}

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
