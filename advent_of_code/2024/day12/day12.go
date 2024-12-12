package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

var dirs = []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

func main() {
	input, _ := os.ReadFile("advent_of_code/2024/day12/input.txt")
	grid := parseGrid(string(input))
	seen := map[image.Point]bool{}
	part1, part2 := 0, 0

	for p := range grid {
		if seen[p] {
			continue
		}
		area, perimeter, sides := 0, 0, 0
		queue := []image.Point{p}
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			if seen[cur] {
				continue
			}
			seen[cur] = true
			area++
			for _, d := range dirs {
				n := cur.Add(d)
				if grid[n] != grid[cur] {
					perimeter++
					sides += neighborSides(grid, cur, d)
				} else if !seen[n] {
					queue = append(queue, n)
				}
			}
		}
		part1 += area * perimeter
		part2 += area * sides
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

func parseGrid(input string) map[image.Point]rune {
	grid := map[image.Point]rune{}
	for y, line := range strings.Fields(input) {
		for x, r := range line {
			grid[image.Point{X: x, Y: y}] = r
		}
	}
	return grid
}

func neighborSides(grid map[image.Point]rune, p image.Point, d image.Point) int {
	if grid[p.Add(image.Point{X: -d.Y, Y: d.X})] != grid[p] || grid[p.Add(d).Add(image.Point{X: -d.Y, Y: d.X})] == grid[p] {
		return 1
	}
	return 0
}
