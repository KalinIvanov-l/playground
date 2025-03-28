package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	locks, keys := parseInput("advent_of_code/2024/day25/input.txt")
	fmt.Println(part1(locks, keys))
}

func parseInput(name string) ([][]int, [][]int) {
	file, _ := os.Open(name)
	scanner := bufio.NewScanner(file)
	var locks, keys [][]int

	for scanner.Scan() {
		curr := make([]int, 5)
		isLock := false
		for i := 0; i < 7; i++ {
			for c, char := range scanner.Text() {
				if char == '#' {
					curr[c]++
				}
			}
			if i == 0 && slices.Equal(curr, []int{1, 1, 1, 1, 1}) {
				isLock = true
			}
			scanner.Scan()
		}

		for i := range curr {
			curr[i]--
		}
		if isLock {
			locks = append(locks, curr)
		} else {
			keys = append(keys, curr)
		}
	}
	return locks, keys
}

func part1(locks, keys [][]int) (count int) {
	for _, lock := range locks {
		for _, key := range keys {
			check := make([]int, 5)
			for i := range lock {
				check[i] = lock[i] + key[i]
			}

			isValid := true
			for i := range check {
				if check[i] > 5 {
					isValid = false
				}
			}

			if isValid {
				count++
			}
		}
	}
	return
}
