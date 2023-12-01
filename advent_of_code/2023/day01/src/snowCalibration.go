package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fileName := "2023/day01/src/input.txt"
	totalSum := 0

	file, _ := os.Open(fileName)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		calibrationValue := calculateValuePart2(line)
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

func calculateValuePart2(line string) int {
	replacements := map[string]string{
		"one":   "one1one",
		"two":   "two2two",
		"three": "three3three",
		"four":  "four4four",
		"five":  "five5five",
		"six":   "six6six",
		"seven": "seven7seven",
		"eight": "eight8eight",
		"nine":  "nine9nine",
	}

	for word, replacement := range replacements {
		line = strings.ReplaceAll(line, word, replacement)
	}

	re := regexp.MustCompile(`\d`)
	matches := re.FindAllString(line, -1)

	firstDigit := matches[0]
	lastDigit := matches[len(matches)-1]

	number, _ := strconv.Atoi(firstDigit + lastDigit)
	return number
}
