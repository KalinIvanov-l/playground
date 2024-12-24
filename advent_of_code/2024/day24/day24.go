package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Gate struct {
	Op     string
	Op1    string
	Op2    string
	Output string
}

func main() {
	inputs, gates, _ := parseInput("advent_of_code/2024/day24/input.txt")

	wireValues := simulateGates(inputs, gates)
	part1 := computeOutputValue(wireValues)
	fmt.Println(part1)

	part2 := swapAndJoinWires(gates)
	fmt.Println(part2)
}

func parseInput(filename string) (map[string]int, []Gate, [][]string) {
	file, _ := os.Open(filename)

	scanner := bufio.NewScanner(file)
	inputs := make(map[string]int)
	var gates []Gate
	var parts [][]string
	var currentPart []string
	readingInputs := true

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			if len(currentPart) > 0 {
				parts = append(parts, currentPart)
				currentPart = []string{}
			}
			readingInputs = false
			continue
		}

		if readingInputs {
			parts := strings.Split(line, ": ")
			val, _ := strconv.Atoi(parts[1])
			inputs[parts[0]] = val
		} else {
			parts := strings.Split(line, " ")
			gates = append(gates, Gate{
				Op:     parts[1],
				Op1:    parts[0],
				Op2:    parts[2],
				Output: parts[4],
			})
		}
	}
	if len(currentPart) > 0 {
		parts = append(parts, currentPart)
	}
	return inputs, gates, parts
}

func simulateGates(inputs map[string]int, gates []Gate) map[string]int {
	wireValues := make(map[string]int)
	for k, v := range inputs {
		wireValues[k] = v
	}

	waitingWires := map[string]bool{}
	for _, gate := range gates {
		if strings.HasPrefix(gate.Output, "z") {
			waitingWires[gate.Output] = true
		}
	}

	for len(waitingWires) > 0 {
		for _, gate := range gates {
			val1, ok1 := wireValues[gate.Op1]
			val2, ok2 := wireValues[gate.Op2]
			if ok1 && ok2 {
				var result int
				switch gate.Op {
				case "AND":
					result = val1 & val2
				case "OR":
					result = val1 | val2
				case "XOR":
					result = val1 ^ val2
				}
				wireValues[gate.Output] = result
				delete(waitingWires, gate.Output)
			}
		}
	}
	return wireValues
}

func computeOutputValue(wireValues map[string]int) int {
	result := 0
	for wire, value := range wireValues {
		if strings.HasPrefix(wire, "z") && value > 0 {
			bitPos, _ := strconv.Atoi(wire[1:])
			result += 1 << bitPos
		}
	}
	return result
}

func find(a, b, operator string, gates []Gate) string {
	for _, gate := range gates {
		if (gate.Op1 == a && gate.Op2 == b) || (gate.Op1 == b && gate.Op2 == a) {
			if gate.Op == operator {
				return gate.Output
			}
		}
	}
	return ""
}

func swapAndJoinWires(gates []Gate) string {
	var swapped []string
	var c0 string

	for i := 0; i < 45; i++ {
		n := fmt.Sprintf("%02d", i)
		var m1, n1, r1, z1, c1 string

		m1 = find("x"+n, "y"+n, "XOR", gates)
		n1 = find("x"+n, "y"+n, "AND", gates)

		if c0 != "" {
			r1 = find(c0, m1, "AND", gates)
			if r1 == "" {
				m1, n1 = n1, m1
				swapped = append(swapped, m1, n1)
				r1 = find(c0, m1, "AND", gates)
			}

			z1 = find(c0, m1, "XOR", gates)
			if strings.HasPrefix(m1, "z") {
				m1, z1 = z1, m1
				swapped = append(swapped, m1, z1)
			}

			if strings.HasPrefix(n1, "z") {
				n1, z1 = z1, n1
				swapped = append(swapped, n1, z1)
			}

			if strings.HasPrefix(r1, "z") {
				r1, z1 = z1, r1
				swapped = append(swapped, r1, z1)
			}
			c1 = find(r1, n1, "OR", gates)
		}

		if strings.HasPrefix(c1, "z") && c1 != "z45" {
			c1, z1 = z1, c1
			swapped = append(swapped, c1, z1)
		}

		if c0 == "" {
			c0 = n1
		} else {
			c0 = c1
		}
	}

	sort.Strings(swapped)
	return strings.Join(swapped, ",")
}
