package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("2023/day09/src/input.txt")
	defer file.Close()

	var inputHistories [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		history, _ := parseLineToHistory(line)
		inputHistories = append(inputHistories, history)
	}

	sum := 0
	for _, history := range inputHistories {
		nextValue := extrapolateNextValue(history)
		sum += nextValue
	}
	fmt.Println("Sum of next values:", sum)
}

func parseLineToHistory(line string) ([]int, error) {
	var history []int
	values := strings.Split(line, " ")
	for _, value := range values {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		history = append(history, intValue)
	}
	return history, nil
}

func extrapolateNextValue(history []int) int {
	sequences := make([][]int, 0)
	sequences = append(sequences, history)

	for {
		lastSeq := sequences[len(sequences)-1]
		newSeq := make([]int, 0)
		isZeroSeq := true

		for i := 0; i < len(lastSeq)-1; i++ {
			diff := lastSeq[i+1] - lastSeq[i]
			newSeq = append(newSeq, diff)
			if diff != 0 {
				isZeroSeq = false
			}
		}
		sequences = append(sequences, newSeq)
		if isZeroSeq {
			break
		}
	}
	for i := len(sequences) - 2; i >= 0; i-- {
		sequences[i] = append([]int{sequences[i][0] - sequences[i+1][0]}, sequences[i]...)
	}
	return sequences[0][0]
}
