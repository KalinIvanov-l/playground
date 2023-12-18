package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func main() {
	file, _ := os.Open("2023/day18/src/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	polygon := []Point{{0, 0}}
	pCnt := 0

	deltas := map[string]Point{
		"R": {1, 0},
		"L": {-1, 0},
		"U": {0, -1},
		"D": {0, 1},
	}

	hex2d := map[string]string{
		"0": "R",
		"1": "D",
		"2": "L",
		"3": "U",
	}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		color := parts[2]
		direction := hex2d[string(color[len(color)-2])]
		delta := deltas[direction]
		steps, _ := strconv.ParseInt(color[2:len(color)-2], 16, 64)

		pCnt += int(steps)
		lastPoint := polygon[len(polygon)-1]
		polygon = append(polygon, Point{lastPoint.x + delta.x*int(steps), lastPoint.y + delta.y*int(steps)})
	}

	totalCapacity := solve(polygon) + pCnt/2 + 1
	fmt.Printf("The lagoon can hold %d cubic meters of lava.\n", totalCapacity)
}

func solve(points []Point) int {
	N := len(points)
	firstx, firsty := points[0].x, points[0].y
	prevx, prevy := firstx, firsty
	res := 0

	for i := 1; i < N; i++ {
		nextx, nexty := points[i].x, points[i].y
		res += getInfo(prevx, prevy, nextx, nexty)
		prevx = nextx
		prevy = nexty
	}
	res += getInfo(prevx, prevy, firstx, firsty)
	return abs(res) / 2
}

func getInfo(x1, y1, x2, y2 int) int {
	return x1*y2 - y1*x2
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
