package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	input, _ := os.ReadFile("advent_of_code/2024/day03/input.txt")
	data := string(input)

	mulRe := regexp.MustCompile(`(do\(\)|don't\(\)|mul\((\d{1,3}),(\d{1,3})\))`)
	part1, part2 := 0, 0
	enabled := true

	for _, match := range mulRe.FindAllStringSubmatch(data, -1) {
		switch match[1] {
		case "do()":
			enabled = true
		case "don't()":
			enabled = false
		default:
			x, _ := strconv.Atoi(match[2])
			y, _ := strconv.Atoi(match[3])
			result := x * y

			part1 += result
			if enabled {
				part2 += result
			}
		}
	}
	fmt.Println(part1)
	fmt.Println(part2)
}
