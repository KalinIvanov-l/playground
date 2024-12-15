package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

type Robot struct {
	P, V image.Point
}

func main() {
	input, _ := os.ReadFile("advent_of_code/2024/day14/input.txt")
	area := image.Rectangle{Min: image.Point{}, Max: image.Point{X: 101, Y: 103}}

	var robots []Robot
	quadrants := map[image.Point]int{}
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		var r Robot
		_, err := fmt.Sscanf(s, "p=%d,%d v=%d,%d", &r.P.X, &r.P.Y, &r.V.X, &r.V.Y)
		if err != nil {
			return
		}
		robots = append(robots, r)
		r.P = r.P.Add(r.V.Mul(100)).Mod(area)
		quadrants[image.Point{X: position(r.P.X - area.Dx()/2), Y: position(r.P.Y - area.Dy()/2)}]++
	}
	fmt.Println(quadrants[image.Point{X: -1, Y: -1}] * quadrants[image.Point{X: 1, Y: -1}] *
		quadrants[image.Point{X: 1, Y: 1}] * quadrants[image.Point{X: -1, Y: 1}])

	for t := 1; ; t++ {
		seen := map[image.Point]struct{}{}
		for i := range robots {
			robots[i].P = robots[i].P.Add(robots[i].V).Mod(area)
			seen[robots[i].P] = struct{}{}
		}
		if len(seen) == len(robots) {
			fmt.Println(t)
			break
		}
	}
}

func position(i int) int {
	if i < 0 {
		return -1
	} else if i > 0 {
		return 1
	}
	return 0
}
