package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

var dirMap = map[string]Point{
	"U": {1, 0},
	"D": {-1, 0},
	"R": {0, 1},
	"L": {0, -1},
}

func main() {
	file, _ := os.Open("2023/day18/src/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var xs, ys []int
	loc := Point{0, 0}
	pts := make(map[Point]bool)
	pts[loc] = true

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		curDir := dirMap[parts[0]]
		curLen, _ := strconv.Atoi(parts[1])
		for i := 0; i <= curLen; i++ {
			newPoint := Point{loc.x + i*curDir.x, loc.y + i*curDir.y}
			pts[newPoint] = true
		}
		loc.x += curLen * curDir.x
		loc.y += curLen * curDir.y
		xs = append(xs, loc.x)
		ys = append(ys, loc.y)
	}
	A := polyArea(xs, ys)
	b := len(pts)
	if b%2 != 0 {
		panic("Boundary points count should be even")
	}
	I := int(A + 1 - float64(b)/2)

	totalCapacity := I + b
	fmt.Printf("The lagoon can hold %d cubic meters of lava.\n", totalCapacity)
}

func polyArea(xs, ys []int) float64 {
	var area float64
	n := len(xs)
	for i := 0; i < n; i++ {
		j := (i + 1) % n
		area += float64(xs[i]*ys[j] - xs[j]*ys[i])
	}
	return math.Abs(area) / 2
}
