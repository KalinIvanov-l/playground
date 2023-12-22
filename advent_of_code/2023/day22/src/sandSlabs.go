package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Brick struct {
	x1, y1, z1, x2, y2, z2 int
}

func NewBrick(coords []int) Brick {
	return Brick{
		x1: coords[0], y1: coords[1], z1: coords[2],
		x2: coords[3], y2: coords[4], z2: coords[5],
	}
}

func droppedBrick(tallest map[[2]int]int, brick Brick) Brick {
	peak := 0
	for x := brick.x1; x <= brick.x2; x++ {
		for y := brick.y1; y <= brick.y2; y++ {
			if tallest[[2]int{x, y}] > peak {
				peak = tallest[[2]int{x, y}]
			}
		}
	}
	dz := max(brick.z1-peak-1, 0)
	return Brick{brick.x1, brick.y1, brick.z1 - dz, brick.x2, brick.y2, brick.z2 - dz}
}

func drop(tower []Brick) (int, []Brick) {
	tallest := make(map[[2]int]int)
	newTower := make([]Brick, len(tower))
	falls := 0
	for i, brick := range tower {
		newBrick := droppedBrick(tallest, brick)
		if newBrick.z1 != brick.z1 {
			falls++
		}
		newTower[i] = newBrick
		for x := brick.x1; x <= brick.x2; x++ {
			for y := brick.y1; y <= brick.y2; y++ {
				tallest[[2]int{x, y}] = newBrick.z2
			}
		}
	}
	return falls, newTower
}

func solve(bricks []Brick) int {
	_, fallen := drop(bricks)
	p1 := 0
	for i := range fallen {
		removed := append([]Brick{}, fallen[:i]...)
		removed = append(removed, fallen[i+1:]...)
		falls, _ := drop(removed)
		if falls == 0 {
			p1++
		}
	}
	return p1
}

func main() {
	file, _ := os.Open("2023/day22/src/input.txt")
	defer file.Close()

	var bricks []Brick
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		coords := parseLine(line)
		bricks = append(bricks, NewBrick(coords))
	}
	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i].z1 < bricks[j].z1
	})

	p1 := solve(bricks)
	fmt.Printf("Part 1: %d\n", p1)
}

func parseLine(line string) []int {
	parts := strings.Split(line, "~")
	allCoords := strings.Split(parts[0]+","+parts[1], ",")
	var coords []int
	for _, str := range allCoords {
		num, _ := strconv.Atoi(str)
		coords = append(coords, num)
	}
	return coords
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
