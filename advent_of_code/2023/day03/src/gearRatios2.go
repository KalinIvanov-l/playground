package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Set map[int]struct{}

func main() {
	file, _ := os.Open("2023/day03/src/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input string
	for scanner.Scan() {
		input += scanner.Text() + "\n"
	}

	total := part2(input)
	fmt.Println("Sum of all gear ratios:", total)
}

func part2(input string) int {
	lines := strings.Split(input, "\n")
	rowLength := len(lines[0])
	data := strings.Join(lines, "")

	reStar := regexp.MustCompile(`\*`)
	stars := reStar.FindAllStringIndex(data, -1)

	partIdxs, partValues := findAllPossibleParts(data)

	total := 0
	for _, star := range stars {
		parts := getAdjacentPartIndexes(star[0], rowLength, partIdxs)
		if len(parts) == 2 {
			ratio := 1
			for k := range parts {
				ratio *= partValues[k]
			}
			total += ratio
		}
	}
	return total
}

func findAllPossibleParts(data string) ([][]int, []int) {
	reDigit := regexp.MustCompile(`\d+`)
	partIndexes := reDigit.FindAllStringIndex(data, -1)

	partValues := make([]int, len(partIndexes))
	for i, s := range reDigit.FindAllString(data, -1) {
		partValues[i], _ = strconv.Atoi(s)
	}
	return partIndexes, partValues
}

func getAdjacentPartIndexes(symbolIdx int, rowLength int, partIndexes [][]int) Set {
	surround := [8]int{
		-rowLength - 1,
		-rowLength,
		-rowLength + 1,
		-1, 1,
		rowLength - 1,
		rowLength,
		rowLength + 1,
	}

	parts := make(Set)
	for _, dx := range surround {
		c := symbolIdx + dx
		for idx, digitIdx := range partIndexes {
			if c >= digitIdx[0] && c < digitIdx[1] {
				parts[idx] = struct{}{}
			}
		}
	}
	return parts
}
