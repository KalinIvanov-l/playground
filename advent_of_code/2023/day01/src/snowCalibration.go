package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	fileName := "2023/day01/src/input.txt"
	totalSum := 0

	file, _ := os.Open(fileName)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		calibrationValue := calculateValue(line)
		totalSum += calibrationValue
	}
	fmt.Printf("Total sum of calibration values: %d\n", totalSum)
}

func calculateValue(line string) int {
	re := regexp.MustCompile(`(\d)`)
	matches := re.FindAllStringSubmatch(line, -1)

	firstDigit := matches[0][1]
	lastDigit := matches[len(matches)-1][1]

	number, _ := strconv.Atoi(firstDigit + lastDigit)
	return number
}
