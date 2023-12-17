package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func readInput(filename string) ([][]rune, error) {
	file, _ := os.Open(filename)
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}
	return grid, scanner.Err()
}

type position struct {
	x         int
	y         int
	accWeight int
	up        int
	down      int
	left      int
	right     int
}

func addPositionIfValid(pq *[]position, newX, newY, newLeft, newRight, newUp, newDown, weight int, s []string, visited map[string]int) {
	if newX >= 0 && newY >= 0 && newX < len(s[0]) && newY < len(s) {
		key := fmt.Sprintf("%d,%d,%d,%d,%d,%d", newX, newY, newUp, newDown, newLeft, newRight)
		if v, ok := visited[key]; !ok || ok && v > weight {
			visited[key] = weight
			*pq = append(*pq, position{x: newX, y: newY, up: newUp, down: newDown, left: newLeft, right: newRight, accWeight: weight})
		}
	}
}

func findShortestPath(s []string) int {
	var (
		priorityQueue []position
		sx, sy        = 0, 0
		dx, dy        = len(s[0]) - 1, len(s) - 1
	)
	priorityQueue = append(priorityQueue, position{x: sx + 1, y: sy, right: 1})
	priorityQueue = append(priorityQueue, position{x: sx, y: sy + 1, down: 1})
	visited := make(map[string]int)

	for len(priorityQueue) > 0 {
		sort.SliceStable(priorityQueue, func(i, j int) bool { return priorityQueue[i].accWeight < priorityQueue[j].accWeight })

		entry := priorityQueue[0]
		priorityQueue = priorityQueue[1:]

		iVal, _ := strconv.Atoi(string(s[entry.y][entry.x]))
		weight := entry.accWeight + iVal
		if entry.x == dx && entry.y == dy {
			return weight
		}

		if entry.x > 0 && entry.left < 3 && entry.right == 0 {
			addPositionIfValid(&priorityQueue, entry.x-1, entry.y, entry.left+1, 0, 0, 0, weight, s, visited)
		}
		if entry.x < len(s[0])-1 && entry.right < 3 && entry.left == 0 {
			addPositionIfValid(&priorityQueue, entry.x+1, entry.y, 0, entry.right+1, 0, 0, weight, s, visited)
		}
		if entry.y > 0 && entry.up < 3 && entry.down == 0 {
			addPositionIfValid(&priorityQueue, entry.x, entry.y-1, 0, 0, entry.up+1, 0, weight, s, visited)
		}
		if entry.y < len(s)-1 && entry.down < 3 && entry.up == 0 {
			addPositionIfValid(&priorityQueue, entry.x, entry.y+1, 0, 0, 0, entry.down+1, weight, s, visited)
		}
	}
	return 0
}

func findShortestPath2(s []string, minSteps, maxSteps int) int {
	var (
		priorityQueue []position
		sx            = 0
		sy            = 0
		dx            = len(s[0]) - 1
		dy            = len(s) - 1
	)
	priorityQueue = append(priorityQueue, position{
		x:     sx + 1,
		y:     0,
		up:    0,
		down:  0,
		left:  0,
		right: 1,
	})
	priorityQueue = append(priorityQueue, position{
		x:     0,
		y:     sy + 1,
		up:    0,
		down:  1,
		left:  0,
		right: 0,
	})
	visited := make(map[string]int)
	for len(priorityQueue) > 0 {
		sort.SliceStable(priorityQueue, func(i, j int) bool {
			return priorityQueue[i].accWeight < priorityQueue[j].accWeight
		})

		entry := priorityQueue[0]
		if len(priorityQueue) > 1 {
			priorityQueue = priorityQueue[1:]
		} else {
			priorityQueue = []position{}
		}

		iVal, _ := strconv.Atoi(string(s[entry.y][entry.x]))
		weight := entry.accWeight + iVal
		if entry.x == dx && entry.y == dy && (entry.left >= minSteps || entry.right >= minSteps || entry.up >= minSteps || entry.down >= minSteps) {
			return weight
		}
		key := fmt.Sprintf("%d,%d,%d,%d,%d,%d", entry.x, entry.y, entry.up, entry.down, entry.left, entry.right)
		if v, ok := visited[key]; !ok || ok && v > weight {
			visited[key] = weight
			if entry.x > 0 && entry.right == 0 && (((entry.up >= minSteps || entry.down >= minSteps) && entry.left == 0) || entry.left > 0 && entry.left < maxSteps) {
				priorityQueue = append(priorityQueue, position{
					x:         entry.x - 1,
					y:         entry.y,
					up:        0,
					down:      0,
					left:      entry.left + 1,
					right:     0,
					accWeight: weight,
				})
			}
			if entry.x < len(s[0])-1 && entry.left == 0 && (((entry.up >= minSteps || entry.down >= minSteps) && entry.right == 0) || entry.right > 0 && entry.right < maxSteps) {
				priorityQueue = append(priorityQueue, position{
					x:         entry.x + 1,
					y:         entry.y,
					up:        0,
					down:      0,
					left:      0,
					right:     entry.right + 1,
					accWeight: weight,
				})
			}
			if entry.y > 0 && entry.down == 0 && (((entry.left >= minSteps || entry.right >= minSteps) && entry.up == 0) || entry.up > 0 && entry.up < maxSteps) {
				priorityQueue = append(priorityQueue, position{
					x:         entry.x,
					y:         entry.y - 1,
					up:        entry.up + 1,
					down:      0,
					left:      0,
					right:     0,
					accWeight: weight,
				})
			}
			if entry.y < len(s)-1 && entry.up == 0 && (((entry.left >= minSteps || entry.right >= minSteps) && entry.down == 0) || entry.down > 0 && entry.down < maxSteps) {
				priorityQueue = append(priorityQueue, position{
					x:         entry.x,
					y:         entry.y + 1,
					up:        0,
					down:      entry.down + 1,
					left:      0,
					right:     0,
					accWeight: weight,
				})
			}
		}
	}
	return 0
}

func main() {
	gridRunes, _ := readInput("2023/day17/src/input.txt")
	grid := make([]string, len(gridRunes))
	for i, line := range gridRunes {
		grid[i] = string(line)
	}
	fmt.Println(findShortestPath(grid))
	fmt.Println(findShortestPath2(grid, 4, 10))
}
