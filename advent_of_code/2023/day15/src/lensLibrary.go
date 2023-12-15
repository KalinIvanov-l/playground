package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("2023/day15/src/input.txt")
	if err != nil {
		panic(err)
	}
	steps := strings.Split(strings.TrimSpace(string(data)), ",")
	sum := 0
	for _, step := range steps {
		sum += hash(step)
	}
	fmt.Println("Sum of HASH algorithm results:", sum)
}

func hash(s string) int {
	currentValue := 0
	for _, c := range s {
		currentValue = (currentValue + int(c)) * 17 % 256
	}
	return currentValue
}
