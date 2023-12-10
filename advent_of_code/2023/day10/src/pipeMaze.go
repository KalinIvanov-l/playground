package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("2023/day10/src/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	var level [][]interface{}
	var startPos complex128

	for y, line := range lines {
		row := []interface{}{}
		for x, c := range line {
			curPos := complex(float64(x), float64(y))
			dirs := []complex128{}
			switch c {
			case '|':
				dirs = []complex128{1i, -1i}
			case '-':
				dirs = []complex128{1, -1}
			case 'L':
				dirs = []complex128{1, -1i}
			case 'J':
				dirs = []complex128{-1, -1i}
			case '7':
				dirs = []complex128{-1, 1i}
			case 'F':
				dirs = []complex128{1, 1i}
			case 'S':
				startPos = curPos
			}
			if len(dirs) > 0 {
				links[curPos] = dirs
			}
			row = append(row, c)
		}
		level = append(level, row)
	}

	path := findPath(links, startPos, []complex128{1, -1, 1i, -1i})
	if path != nil {
		fmt.Println(len(path) / 2)
	}

	pos := startPos
	for _, dir := range path {
		a := pos
		b := pos + dir
		if imag(dir) != 0 {
			level[int(imag(a))][int(real(a))] = complex(imag(dir), imag(dir))
			level[int(imag(b))][int(real(b))] = complex(imag(dir), imag(dir))
		} else {
			realVal, imagVal := int(real(a)), int(imag(a))
			if realVal != 0 || imagVal != 0 {
				level[imagVal][realVal] = complex(float64(realVal), float64(imagVal))
			}

		}
		pos = b
	}

	sum := 0
	for _, row := range level {
		first := interface{}(nil)
		count := false
		rowSum := 0

		for _, e := range row {
			if val, ok := e.(float64); ok {
				if !count {
					first = val
					count = true
				} else {
					count = first == val
				}
			} else {
				if count {
					rowSum++
				}
			}
		}
		sum += rowSum
	}
	fmt.Println(sum)
}

var links = make(map[complex128][]complex128)

func findPath(links map[complex128][]complex128, startPos complex128, directions []complex128) []complex128 {
	for _, dir := range directions {
		path := []complex128{dir}
		pos := startPos

		for {
			dir := path[len(path)-1]
			pos += dir

			if pos == startPos {
				return path
			}
			nextDir, ok := links[pos]
			if !ok {
				break
			}
			var newNextDir []complex128
			for _, d := range nextDir {
				if d != -dir {
					newNextDir = append(newNextDir, d)
				}
			}
			if len(newNextDir) != 1 {
				break
			}
			path = append(path, newNextDir[0])
		}
	}
	return nil
}
