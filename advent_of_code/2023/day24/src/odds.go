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
	LEAST     = 200000000000000
	MOST      = 400000000000000
)

type Line struct {
	a, b, c float64
}

type Stone struct {
	x, y, z    int
	vx, vy, vz int
	line       Line
}

func parseNums(str string) (int, int, int) {
	parts := strings.Split(str, ", ")
	x, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
	y, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
	z, _ := strconv.Atoi(strings.TrimSpace(parts[2]))
	return x, y, z
}

func Solve() {
	file, _ := os.Open(InputFile)
	defer file.Close()

	var stones []Stone
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		p, v, _ := strings.Cut(scanner.Text(), " @ ")
		x, y, z := parseNums(p)
		vx, vy, vz := parseNums(v)

		m := float64(vy) / float64(vx)
		c := float64(y) - m*float64(x)
		stones = append(stones, Stone{x, y, z, vx, vy, vz, Line{m, -1, c}})
	}

	count := 0
	for i := 0; i < len(stones)-1; i++ {
		for j := i + 1; j < len(stones); j++ {
			s1, s2 := stones[i], stones[j]
			a1b2A2b1 := (s1.line.a * s2.line.b) - (s2.line.a * s1.line.b)
			if a1b2A2b1 == 0 {
				continue
			}

			ix, iy := (s1.line.b*s2.line.c-s2.line.b*s1.line.c)/a1b2A2b1, (s1.line.c*s2.line.a-s2.line.c*s1.line.a)/a1b2A2b1
			t1, t2 := (ix-float64(s1.x))/float64(s1.vx), (ix-float64(s2.x))/float64(s2.vx)

			if t1 > 0 && t2 > 0 && ix >= LEAST && ix <= MOST && iy >= LEAST && iy <= MOST {
				count++
			}
		}
	}
	fmt.Println("Part 1: ", count)

	fmt.Println("var('x y z vx vy vz t1 t2 t3 ans')")
	for i := 0; i < 3; i++ {
		fmt.Printf("eq%d = x + (vx * t%d) == %d + (%d * t%d)\n", i*3+1, i+1, stones[i].x, stones[i].vx, i+1)
		fmt.Printf("eq%d = y + (vy * t%d) == %d + (%d * t%d)\n", i*3+2, i+1, stones[i].y, stones[i].vy, i+1)
		fmt.Printf("eq%d = z + (vz * t%d) == %d + (%d * t%d)\n", i*3+3, i+1, stones[i].z, stones[i].vz, i+1)
	}
	fmt.Println("eq10 = ans == x + y + z")
	fmt.Println("print(solve([eq1,eq2,eq3,eq4,eq5,eq6,eq7,eq8,eq9,eq10],x,y,z,vx,vy,vz,t1,t2,t3,ans))")

	fmt.Println()
	fmt.Println("Part 2: ", 149412455352770+174964385672289+233413147425100)
}

func main() {
	Solve()
}
