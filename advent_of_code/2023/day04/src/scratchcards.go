package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputFile := "2023/day04/src/input.txt"
	cards, _ := readInputFile(inputFile)

	totalPoints := part2(cards)
	fmt.Println(totalPoints)
}

func calculateTotalPoints(cards []string) int {
	totalPoints := 0
	for _, card := range cards {
		winning, picked := getNumbers(card)
		points := calculateCardPoints(winning, picked)
		totalPoints += points
	}
	return totalPoints
}

func getNumbers(line string) (winning []int, picked []int) {
	parts := strings.Split(line, "|")
	winning = convertToIntSlice(strings.Fields(parts[0]))
	picked = convertToIntSlice(strings.Fields(parts[1]))
	return
}

func convertToIntSlice(strSlice []string) []int {
	var intSlice []int
	for _, str := range strSlice {
		num, _ := strconv.Atoi(str)
		intSlice = append(intSlice, num)
	}
	return intSlice
}

func calculateCardPoints(winning, picked []int) int {
	points := 0
	for _, p := range picked {
		if contains(winning, p) {
			if points == 0 {
				points = 1
			} else {
				points *= 2
			}
		}
	}
	return points
}

func part2(cards []string) (total int) {
	cardCount := make(map[int]int)

	for idx, card := range cards {
		winning, picked := getNumbers(card)
		matches := calcCardMatches(winning, picked)
		copiesOfCurrent := cardCount[idx]
		for i := 1; i <= matches; i++ {
			if idx+i < len(cards) {
				cardCount[idx+i] += copiesOfCurrent + 1
			}
		}
	}
	total = len(cards)
	for _, count := range cardCount {
		total += count
	}
	return
}

func calcCardMatches(winning []int, picked []int) (matches int) {
	for _, win := range winning {
		for _, pick := range picked {
			if win == pick {
				matches++
				break
			}
		}
	}
	return matches
}

func contains(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func readInputFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
