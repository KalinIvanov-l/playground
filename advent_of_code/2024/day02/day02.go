package main

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	input, _ := os.ReadFile("advent_of_code/2024/day02/input.txt")

	part1, part2 := 0, 0
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		var report []int
		err := json.Unmarshal([]byte("["+strings.ReplaceAll(s, " ", ",")+"]"), &report)
		if err != nil {
			return
		}

		if isSafe(report) {
			part1++
		}
		for i := range report {
			if isSafe(slices.Delete(slices.Clone(report), i, i+1)) {
				part2++
				break
			}
		}
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

func isSafe(levels []int) bool {
	for i := 1; i < len(levels); i++ {
		if d := levels[i] - levels[i-1]; d*(levels[1]-levels[0]) <= 0 || d < -3 || d > 3 {
			return false
		}
	}
	return true
}
