package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func diff(a, b string) int {
	count := 0
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			count++
		}
	}
	return count
}

func solve(grid [][]string) int {
	max := len(grid) - 1
	diffs := make(map[int][]int)

	for n := 0; n <= max; n++ {
		a := n
		b := n + 1
		var line []int

		for a >= 0 && b <= max {
			prev := grid[a]
			next := grid[b]

			line = append(line, diff(strings.Join(prev, ""), strings.Join(next, "")))
			a--
			b++
		}

		if len(line) > 0 {
			diffs[n] = line
		}
	}

	for k, v := range diffs {
		sum := 0
		for _, d := range v {
			sum += d
		}
		if sum == 1 {
			return k + 1
		}
	}
	return 0
}

func solveTask(grid [][]string) int {
	max := len(grid[0]) - 1
	diffs := make(map[int][]int)

	for n := 0; n <= max; n++ {
		a := n
		b := n + 1
		var line []int

		for a >= 0 && b <= max {
			var prev, next string
			for _, row := range grid {
				if a < len(row) {
					prev += row[a]
				}
				if b < len(row) {
					next += row[b]
				}
			}
			line = append(line, diff(prev, next))
			a--
			b++
		}

		if len(line) > 0 {
			diffs[n] = line
		}
	}

	for k, v := range diffs {
		sum := 0
		for _, d := range v {
			sum += d
		}
		if sum == 1 {
			return k + 1
		}
	}
	return 0
}

func main() {
	file, _ := os.Open("2023/day13/src/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var grids [][][]string
	var grid [][]string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			grids = append(grids, grid)
			grid = nil
			continue
		}
		grid = append(grid, strings.Split(line, ""))
	}
	if grid != nil {
		grids = append(grids, grid)
	}

	totalScore := 0
	for _, g := range grids {
		totalScore += solve(g) * 100
		totalScore += solveTask(g)
	}
	fmt.Println("Total Score with Fixes:", totalScore)
}
