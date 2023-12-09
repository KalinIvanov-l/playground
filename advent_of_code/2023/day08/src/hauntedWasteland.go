package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

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
		nodes[node] = [2]string{connectionParts[0], connectionParts[1]}
	}
	startNodes := make([]string, 0)
	for node := range nodes {
		if strings.HasSuffix(node, "A") {
			startNodes = append(startNodes, node)
		}
	}

	pathLengths := make([]int, 0)
	for _, start := range startNodes {
		current := start
		steps := 0
		for !strings.HasSuffix(current, "Z") {
			direction := instructions[steps%len(instructions)]
			if direction == 'L' {
				current = nodes[current][0]
			} else {
				current = nodes[current][1]
			}
			steps++
		}
		pathLengths = append(pathLengths, steps)
	}

	result := pathLengths[0]
	for _, length := range pathLengths[1:] {
		result = lcm(result, length)
	}
	fmt.Println(result)
}
