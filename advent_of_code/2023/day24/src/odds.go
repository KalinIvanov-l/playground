package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	InputFile = "2023/day24/src/input.txt"
	Min       = 200000000000000
	Max       = 400000000000000
)

type Point struct {
	X, Y float64
}

func intersect(x1, y1, x2, y2, x3, y3, x4, y4 float64) *Point {
	denom := (y4-y3)*(x2-x1) - (x4-x3)*(y2-y1)
	ua := ((x4-x3)*(y1-y3) - (y4-y3)*(x1-x3)) / denom
	return &Point{X: x1 + ua*(x2-x1), Y: y1 + ua*(y2-y1)}
}

func main() {
	file, _ := os.Open(InputFile)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	bean := 0
	for i := 0; i < len(lines); i++ {
		for j := i + 1; j < len(lines); j++ {
			line1 := strings.Split(lines[i], " @ ")
			line2 := strings.Split(lines[j], " @ ")

			pos1 := strings.Split(line1[0], ", ")
			vel1 := strings.Split(line1[1], ", ")
			pos2 := strings.Split(line2[0], ", ")
			vel2 := strings.Split(line2[1], ", ")

			x1, _ := strconv.ParseFloat(pos1[0], 64)
			y1, _ := strconv.ParseFloat(pos1[1], 64)
			x2 := x1 + parseFloat(vel1[0])
			y2 := y1 + parseFloat(vel1[1])

			x3, _ := strconv.ParseFloat(pos2[0], 64)
			y3, _ := strconv.ParseFloat(pos2[1], 64)
			x4 := x3 + parseFloat(vel2[0])
			y4 := y3 + parseFloat(vel2[1])

			intersection := intersect(x1, y1, x2, y2, x3, y3, x4, y4)
			if intersection != nil {
				x := intersection.X
				y := intersection.Y

				if (x > x1) == (x2 > x1) && (y > y1) == (y2 > y1) &&
					(x > x3) == (x4 > x3) && (y > y3) == (y4 > y3) &&
					x >= Min && x <= Max && y >= Min && y <= Max {
					bean++
				}
			}
		}
	}
	fmt.Println("answer part 1:", bean)
}

func parseFloat(s string) float64 {
	v, _ := strconv.ParseFloat(s, 64)
	return v
}
