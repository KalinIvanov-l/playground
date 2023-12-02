package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filePath := "2023/day02/src/input.txt"
	file, _ := os.Open(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalPower := 0
	for scanner.Scan() {
		line := scanner.Text()
		minCubes, _ := getMinCubes(line)
		power := minCubes["red"] * minCubes["green"] * minCubes["blue"]
		totalPower += power
	}
	fmt.Printf("Sum of the power of these sets: %d\n", totalPower)
}

/*
First part
*/
//totalSum := 0
//for scanner.Scan() {
//	line := scanner.Text()
//	if feasibleGame(line) {
//		gameID, _ := getGameID(line)
//		totalSum += gameID
//	}
//}
//
//if err := scanner.Err(); err != nil {
//	fmt.Printf("Error reading from file: %v\n", err)
//} else {
//	fmt.Printf("Sum of IDs of feasible games: %d\n", totalSum)
//}

func feasibleGame(game string) bool {
	maxCounts := map[string]int{"red": 12, "green": 13, "blue": 14}
	parts := strings.Split(game, ";")
	for _, part := range parts {
		counts, _ := parseCounts(part)
		for color, count := range counts {
			if count > maxCounts[color] {
				return false
			}
		}
	}
	return true
}

func parseCounts(part string) (map[string]int, error) {
	counts := map[string]int{"red": 0, "green": 0, "blue": 0}
	tokens := strings.Fields(part)

	for i := 0; i < len(tokens); i += 2 {
		countStr := tokens[i]
		color := strings.Trim(tokens[i+1], ", ")
		if _, err := strconv.Atoi(countStr); err != nil {
			continue
		}
		count, _ := strconv.Atoi(countStr)
		if currentMax, exists := counts[color]; exists {
			if count > currentMax {
				counts[color] = count
			}
		}
	}
	return counts, nil
}

func getMinCubes(game string) (map[string]int, error) {
	counts := map[string]int{"red": 0, "green": 0, "blue": 0}
	parts := strings.Split(game, ";")
	for _, part := range parts {
		partCounts, _ := parseCounts(part)
		for color, count := range partCounts {
			if count > counts[color] {
				counts[color] = count
			}
		}
	}
	return counts, nil
}

func getGameID(game string) (int, error) {
	parts := strings.Split(game, ":")
	idStr := strings.TrimSpace(strings.Split(parts[0], " ")[1])
	return strconv.Atoi(idStr)
}
