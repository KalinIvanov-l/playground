package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("advent_of_code/2024/day19/input.txt")

	split := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	var ways func(string) int
	cache := map[string]int{}

	ways = func(design string) (n int) {
		if n, ok := cache[design]; ok {
			return n
		}
		defer func() { cache[design] = n }()

		if design == "" {
			return 1
		}
		for _, s := range strings.Split(split[0], ", ") {
			if strings.HasPrefix(design, s) {
				n += ways(design[len(s):])
			}
		}
		return n
	}

	part1, part2 := 0, 0
	for _, s := range strings.Fields(split[1]) {
		if w := ways(s); w > 0 {
			part1++
			part2 += w
		}
	}
	fmt.Println(part1)
	fmt.Println(part2)
}
