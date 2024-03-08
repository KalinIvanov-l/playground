package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	X, Y int
	Type rune
}

func NewPoint(x, y int, t rune) *Point {
	return &Point{X: x, Y: y, Type: t}
}

func (p *Point) IsWalkable() bool {
	return p.Type != '#'
}

type Node struct {
	Point    *Point
	Children []*Node
}

func (n *Node) AddChild(child *Node) {
	n.Children = append(n.Children, child)
}

func main() {
	file, _ := os.Open("2023/day23/src/input.txt")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	var grid [][]*Point
	y := 0

	for scanner.Scan() {
		line := scanner.Text()
		var row []*Point
		for x, char := range line {
			row = append(row, NewPoint(x, y, char))
		}
		grid = append(grid, row)
		y++
	}
	start, end := findStartEnd(grid)
	maxLength := findLongestPath(grid, start, end, make(map[string]bool), 0)
	fmt.Println("Longest Path Length:", maxLength)
}

func findStartEnd(grid [][]*Point) (*Point, *Point) {
	var start, end *Point
	for x, cell := range grid[0] {
		if cell.IsWalkable() {
			start = grid[0][x]
			break
		}
	}
	for x, cell := range grid[len(grid)-1] {
		if cell.IsWalkable() {
			end = grid[len(grid)-1][x]
			break
		}
	}
	return start, end
}

func findLongestPath(grid [][]*Point, current, end *Point, visited map[string]bool, length int) int {
	if current.X == end.X && current.Y == end.Y {
		return length
	}

	key := fmt.Sprintf("%d,%d", current.X, current.Y)
	if visited[key] {
		return 0
	}
	visited[key] = true

	maxLength := 0
	directions := getDirections()

	for _, dir := range directions {
		nextX, nextY := current.X+dir.X, current.Y+dir.Y
		if nextX >= 0 && nextY >= 0 && nextX < len(grid[0]) && nextY < len(grid) && grid[nextY][nextX].IsWalkable() {
			newLength := findLongestPath(grid, grid[nextY][nextX], end, cloneMap(visited), length+1)
			if newLength > maxLength {
				maxLength = newLength
			}
		}
	}
	return maxLength
}

func getDirections() []*Point {
	return []*Point{{X: 0, Y: -1}, {X: 0, Y: 1}, {X: -1, Y: 0}, {X: 1, Y: 0}}
}

func cloneMap(m map[string]bool) map[string]bool {
	newMap := make(map[string]bool)
	for k, v := range m {
		newMap[k] = v
	}
	return newMap
}
