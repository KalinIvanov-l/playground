package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	x, y, z int
}

type Brick []Coord

func main() {
	file, _ := os.Open("2023/day22/src/input.txt")
	defer file.Close()

	var bricks []Brick
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		bricks = append(bricks, parseBrick(line))
	}

	seen := make(map[Coord]bool)
	for _, brick := range bricks {
		for _, coord := range brick {
			seen[coord] = true
		}
	}
	simulateBricksFalling(bricks, seen)
	oldSeen := deepcopy(seen)
	oldBricks := deepcopyBricks(bricks)

	p1, p2 := 0, 0
	for i, brick := range bricks {
		seen = deepcopy(oldSeen)
		bricks = deepcopyBricks(oldBricks)

		for _, coord := range brick {
			delete(seen, coord)
		}

		fall := make(map[int]struct{})
		for {
			a := false
			for j, b := range bricks {
				if j == i {
					continue
				}
				ok := true
				for _, c := range b {
					if c.z == 1 || (seen[Coord{c.x, c.y, c.z - 1}] && !containsCoord(b, Coord{c.x, c.y, c.z - 1})) {
						ok = false
						break
					}
				}
				if ok {
					a = true
					fall[j] = struct{}{}
					for _, c := range b {
						delete(seen, c)
						seen[Coord{c.x, c.y, c.z - 1}] = true
					}
					bricks[j] = moveBrickDown(b)
				}
			}
			if !a {
				break
			}
		}
		if len(fall) == 0 {
			p1++
		}
		p2 += len(fall)
	}
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}

func parseBrick(line string) Brick {
	parts := strings.Split(line, "~")
	start := parseCoord(parts[0])
	end := parseCoord(parts[1])

	var brick Brick
	if start.x == end.x && start.y == end.y {
		for z := start.z; z <= end.z; z++ {
			brick = append(brick, Coord{start.x, start.y, z})
		}
	} else if start.x == end.x && start.z == end.z {
		for y := start.y; y <= end.y; y++ {
			brick = append(brick, Coord{start.x, y, start.z})
		}
	} else if start.y == end.y && start.z == end.z {
		for x := start.x; x <= end.x; x++ {
			brick = append(brick, Coord{x, start.y, start.z})
		}
	}
	return brick
}

func parseCoord(part string) Coord {
	values := strings.Split(part, ",")
	x, _ := strconv.Atoi(values[0])
	y, _ := strconv.Atoi(values[1])
	z, _ := strconv.Atoi(values[2])
	return Coord{x, y, z}
}

func simulateBricksFalling(bricks []Brick, seen map[Coord]bool) {
	for {
		a := false
		for i := range bricks {
			ok := true
			for _, coord := range bricks[i] {
				if coord.z == 1 || (seen[Coord{coord.x, coord.y, coord.z - 1}] && !containsCoord(bricks[i], Coord{coord.x, coord.y, coord.z - 1})) {
					ok = false
					break
				}
			}
			if ok {
				a = true
				for _, coord := range bricks[i] {
					delete(seen, coord)
					seen[Coord{coord.x, coord.y, coord.z - 1}] = true
				}
				bricks[i] = moveBrickDown(bricks[i])
			}
		}
		if !a {
			break
		}
	}
}

func deepcopy(m map[Coord]bool) map[Coord]bool {
	cp := make(map[Coord]bool)
	for k, v := range m {
		cp[k] = v
	}
	return cp
}

func deepcopyBricks(bricks []Brick) []Brick {
	cp := make([]Brick, len(bricks))
	for i, b := range bricks {
		cp[i] = make(Brick, len(b))
		copy(cp[i], b)
	}
	return cp
}

func containsCoord(brick Brick, coord Coord) bool {
	for _, c := range brick {
		if c == coord {
			return true
		}
	}
	return false
}

func moveBrickDown(brick Brick) Brick {
	moved := make(Brick, len(brick))
	for i, coord := range brick {
		moved[i] = Coord{coord.x, coord.y, coord.z - 1}
	}
	return moved
}
