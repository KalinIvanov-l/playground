package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("advent_of_code/2024/day06/input.txt")

	grid := make(map[image.Point]rune)
	var start image.Point
	for y, line := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		for x, r := range line {
			p := image.Point{X: x, Y: y}
			grid[p] = r
			if r == '^' {
				start = p
			}
		}
	}
	part1 := func(p image.Point, d int) int {
		delta := []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
		seen := make(map[image.Point][]int)
		for {
			if _, ok := grid[p]; !ok {
				return len(seen)
			}
			if ds, ok := seen[p]; ok && len(ds) > 0 {
				for _, dir := range ds {
					if dir == d {
						return -1
					}
				}
			}
			seen[p] = append(seen[p], d)
			next := p.Add(delta[d])
			if grid[next] == '#' {
				d = (d + 1) % 4
			} else {
				p = next
			}
		}
	}
	part2 := 0
	for p, r := range grid {
		if r != '.' {
			continue
		}
		grid[p] = '#'
		if part1(start, 0) == -1 {
			part2++
		}
		grid[p] = '.'
	}

	fmt.Println(part1(start, 0))
	fmt.Println(part2)
}
