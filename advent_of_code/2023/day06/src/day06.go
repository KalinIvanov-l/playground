package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Race struct {
	Time     int
	Distance int
}

// first part.
func parseInput(input string) []Race {
	reNum := regexp.MustCompile(`\d+`)
	digStr := reNum.FindAllString(input, -1)

	var digs []int
	for _, str := range digStr {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			continue
		}
		digs = append(digs, num)
	}

	var races []Race
	for i := 0; i < len(digs)/2; i++ {
		races = append(races, Race{Time: digs[i], Distance: digs[len(digs)/2+i]})
	}
	return races
}

// second part.
func parseInputForPartTwo(input string) Race {
	reNum := regexp.MustCompile(`\d+`)
	digStr := reNum.FindAllString(input, -1)

	timeStr := ""
	distanceStr := ""
	for i, str := range digStr {
		if i < len(digStr)/2 {
			timeStr += str
		} else {
			distanceStr += str
		}
	}
	time, _ := strconv.Atoi(timeStr)
	distance, _ := strconv.Atoi(distanceStr)
	return Race{Time: time, Distance: distance}
}

func calculateWaysToWin(race Race) int {
	waysToWin := 0
	for holdTime := 0; holdTime < race.Time; holdTime++ {
		speed := holdTime
		travelTime := race.Time - holdTime
		distanceTravelled := speed * travelTime
		if distanceTravelled > race.Distance {
			waysToWin++
		}
	}
	return waysToWin
}

func main() {
	file, _ := os.Open("2023/day06/src/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := ""
	for scanner.Scan() {
		input += scanner.Text() + "\n"
	}

	race := parseInputForPartTwo(input)
	totalWays := calculateWaysToWin(race)
	fmt.Println("Total number of ways to win the race:", totalWays)
}
