package main

import (
	"fmt"
	"image"
	"os"
	"slices"
	"strings"
)

func main() {
	input, _ := os.ReadFile("advent_of_code/2024/day06/input.txt")

	grid, start := map[image.Point]rune{}, image.Point{}
	for y, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		for x, r := range s {
			if r == '^' {
				start = image.Point{X: x, Y: y}
			}
			grid[image.Point{X: x, Y: y}] = r
		}
	}
	part1 := func(p image.Point, d int) int {
		delta := []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
		seen := map[image.Point][]int{}
		for {
			if _, ok := grid[p]; !ok {
				return len(seen)
			}
			if ds, ok := seen[p]; ok && slices.Contains(ds, d) {
				return -1
			}
			seen[p] = append(seen[p], d)
			if n := p.Add(delta[d]); grid[n] == '#' {
				d = (d + 1) % len(delta)
			} else {
				p = n
			}
		}
	}
	part2 := 0
	for p := range grid {
		if grid[p] != '.' {
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
