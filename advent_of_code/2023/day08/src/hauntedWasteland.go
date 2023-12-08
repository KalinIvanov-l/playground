package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("2023/day08/src/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	nodes := make(map[string][2]string)
	var instructions string
	firstLine := true

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		if firstLine {
			instructions = line
			firstLine = false
			continue
		}
		parts := strings.Split(line, " = ")
		node := parts[0]
		connections := strings.Trim(parts[1], "()")
		connectionParts := strings.Split(connections, ", ")
		if len(connectionParts) != 2 {
			fmt.Println("Error parsing connections for node:", node)
			return
		}

		nodes[node] = [2]string{connectionParts[0], connectionParts[1]}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	current := "AAA"
	steps := 0
	i := 0

	for current != "ZZZ" {
		if i >= len(instructions) {
			i = 0
		}
		direction := instructions[i]
		if direction == 'L' {
			current = nodes[current][0]
		} else {
			current = nodes[current][1]
		}
		steps++
		i++
	}
	fmt.Println(steps)
}
