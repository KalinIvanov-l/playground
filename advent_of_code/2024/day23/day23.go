package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func isConnected(a, b string, connections map[string]map[string]bool) bool {
	return connections[a][b] && connections[b][a]
}

func isTriangle(a, b, c string, connections map[string]map[string]bool) bool {
	return connections[a][b] && connections[b][c] && connections[a][c]
}

func bronKerbosch(R, P, X []string, connections map[string]map[string]bool) [][]string {
	var cliques [][]string
	if len(P) == 0 && len(X) == 0 {
		cliques = append(cliques, append([]string{}, R...))
		return cliques
	}

	for _, v := range P {
		intersectP := []string{}
		intersectX := []string{}
		for _, u := range append([]string{}, P...) {
			if isConnected(v, u, connections) {
				intersectP = append(intersectP, u)
			}
		}
		for _, u := range append([]string{}, X...) {
			if isConnected(v, u, connections) {
				intersectX = append(intersectX, u)
			}
		}

		cliques = append(cliques, bronKerbosch(append(R, v), intersectP, intersectX, connections)...)
		P = removeFromSlice(P, v)
		X = append(X, v)
	}
	return cliques
}

func removeFromSlice(slice []string, elem string) []string {
	var newSlice []string
	for _, s := range slice {
		if s != elem {
			newSlice = append(newSlice, s)
		}
	}
	return newSlice
}

func main() {
	input, _ := os.ReadFile("advent_of_code/2024/day23/input.txt")

	connections := make(map[string]map[string]bool)
	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			continue
		}
		parts := strings.Split(line, "-")
		computer1 := parts[0]
		computer2 := parts[1]

		if _, exists := connections[computer1]; !exists {
			connections[computer1] = make(map[string]bool)
		}
		if _, exists := connections[computer2]; !exists {
			connections[computer2] = make(map[string]bool)
		}

		connections[computer1][computer2] = true
		connections[computer2][computer1] = true
	}

	var validTriangles [][]string
	for computer1 := range connections {
		for computer2 := range connections {
			if computer1 >= computer2 {
				continue
			}
			for computer3 := range connections {
				if computer2 >= computer3 || computer1 >= computer3 {
					continue
				}
				if isTriangle(computer1, computer2, computer3, connections) {
					if strings.HasPrefix(computer1, "t") ||
						strings.HasPrefix(computer2, "t") ||
						strings.HasPrefix(computer3, "t") {
						validTriangles = append(validTriangles, []string{computer1, computer2, computer3})
					}
				}
			}
		}
	}
	fmt.Println(len(validTriangles))

	var allComputers []string
	for computer := range connections {
		allComputers = append(allComputers, computer)
	}
	cliques := bronKerbosch([]string{}, allComputers, []string{}, connections)

	var largestClique []string
	for _, clique := range cliques {
		if len(clique) > len(largestClique) {
			largestClique = clique
		}
	}

	sort.Strings(largestClique)
	fmt.Println(strings.Join(largestClique, ","))
}
