package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const (
	pipeVertical      = '|'
	pipeHorizontal    = '-'
	pipeBendNorthEast = 'L'
	pipeBendNorthWest = 'J'
	pipeBendSouthWest = '7'
	pipeBendSouthEast = 'F'
	ground            = '.'
	startingPosition  = 'S'
)

type position struct {
	row, col int
}

func main() {
	file, _ := os.Open("2023/day10/src/input.txt")
	defer file.Close()

	pipeMap, _ := pipeMapFromReader(bufio.NewReader(file))
	loop, _ := findLoop(pipeMap)
	area := areaWithinLoop(pipeMap, loop)
	fmt.Println("Enclosed tiles:", area)
}

func areaWithinLoop(pipeMap [][]byte, loop []position) int {
	zoomedInMap := zoomInMap(pipeMap, loop)
	outsideArea := calculateOutsideArea(zoomedInMap)
	mapArea := len(pipeMap) * len(pipeMap[0])
	return mapArea - outsideArea - len(loop)
}

func zoomInMap(pipeMap [][]byte, loop []position) [][]byte {
	zoomedInMap := make([][]byte, len(pipeMap)*2+1)
	for i := range zoomedInMap {
		zoomedInMap[i] = make([]byte, len(pipeMap[0])*2+1)
		fillWithGround(zoomedInMap[i])
	}
	addPipesToZoomedMap(zoomedInMap, pipeMap, loop)
	return zoomedInMap
}

func fillWithGround(row []byte) {
	for col := range row {
		row[col] = ground
	}
}

func addPipesToZoomedMap(zoomedInMap [][]byte, pipeMap [][]byte, loop []position) {
	for i := range loop {
		pos, nextPos := loop[i], loop[(i+1)%len(loop)]
		addPipeToZoomedMap(zoomedInMap, pos, nextPos, pipeMap)
	}
}

func addPipeToZoomedMap(zoomedInMap [][]byte, pos, nextPos position, pipeMap [][]byte) {
	zoomedInMap[pos.row*2+1][pos.col*2+1] = pipeMap[pos.row][pos.col]
	rowDelta, colDelta := nextPos.row-pos.row, nextPos.col-pos.col
	switch {
	case rowDelta == -1 && colDelta == 0:
		zoomedInMap[pos.row*2][pos.col*2+1] = pipeVertical
	case rowDelta == 1 && colDelta == 0:
		zoomedInMap[pos.row*2+2][pos.col*2+1] = pipeVertical
	case rowDelta == 0 && colDelta == -1:
		zoomedInMap[pos.row*2+1][pos.col*2] = pipeHorizontal
	case rowDelta == 0 && colDelta == 1:
		zoomedInMap[pos.row*2+1][pos.col*2+2] = pipeHorizontal
	default:
		panic("diagonal pipe")
	}
}

func calculateOutsideArea(zoomedInMap [][]byte) int {
	outsideArea := 0
	seen := make([][]bool, len(zoomedInMap))
	for i := range seen {
		seen[i] = make([]bool, len(zoomedInMap[i]))
	}
	processPositionsForOutsideArea(&outsideArea, zoomedInMap, seen)
	return outsideArea
}

func processPositionsForOutsideArea(outsideArea *int, zoomedInMap [][]byte, seen [][]bool) {
	stack := []position{{0, 0}}
	seen[0][0] = true

	for len(stack) > 0 {
		pos := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for _, delta := range []position{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			newPos := position{pos.row + delta.row, pos.col + delta.col}
			if isValidPosition(newPos, zoomedInMap, seen) {
				stack = append(stack, newPos)
				seen[newPos.row][newPos.col] = true
				if newPos.row%2 == 1 && newPos.col%2 == 1 {
					*outsideArea++
				}
			}
		}
	}
}

func isValidPosition(pos position, zoomedInMap [][]byte, seen [][]bool) bool {
	return pos.row >= 0 && pos.row < len(zoomedInMap) &&
		pos.col >= 0 && pos.col < len(zoomedInMap[pos.row]) &&
		!seen[pos.row][pos.col] && zoomedInMap[pos.row][pos.col] == ground
}

func findLoop(pipeMap [][]byte) ([]position, error) {
	startingPosition, _ := findStartingPosition(pipeMap)
	seen := make([][]bool, len(pipeMap))
	for i := range seen {
		seen[i] = make([]bool, len(pipeMap[i]))
	}
	loop := []position{startingPosition}
	seen[startingPosition.row][startingPosition.col] = true

	for {
		neighbors := findConnectedNeighbors(pipeMap, loop[len(loop)-1])
		for len(neighbors) > 0 && seen[neighbors[0].row][neighbors[0].col] {
			neighbors = neighbors[1:]
		}
		if len(neighbors) == 0 {
			break
		}
		loop = append(loop, neighbors[0])
		seen[neighbors[0].row][neighbors[0].col] = true
	}
	return loop, nil
}

func findStartingPosition(pipeMap [][]byte) (position, error) {
	for row := range pipeMap {
		for col := range pipeMap[row] {
			if pipeMap[row][col] == startingPosition {
				return position{row, col}, nil
			}
		}
	}
	return position{}, fmt.Errorf("no starting position found")
}

func findConnectedNeighbors(pipeMap [][]byte, pos position) []position {
	shape := pipeMap[pos.row][pos.col]
	return getNeighborsForShape(shape, pos, pipeMap)
}

func getNeighborsForShape(shape byte, pos position, pipeMap [][]byte) []position {
	switch shape {
	case startingPosition:
		return findStartingPositionConnectedNeighbors(pipeMap, pos)
	case pipeVertical:
		return getVerticalNeighbors(pos, pipeMap)
	case pipeHorizontal:
		return getHorizontalNeighbors(pos, pipeMap)
	case pipeBendNorthEast, pipeBendNorthWest, pipeBendSouthWest, pipeBendSouthEast:
		return getBendNeighbors(shape, pos, pipeMap)
	default:
		return nil
	}
}

func getVerticalNeighbors(pos position, pipeMap [][]byte) []position {
	var neighbors []position
	if pos.row > 0 {
		neighbors = append(neighbors, position{pos.row - 1, pos.col})
	}
	if pos.row < len(pipeMap)-1 {
		neighbors = append(neighbors, position{pos.row + 1, pos.col})
	}
	return neighbors
}

func getHorizontalNeighbors(pos position, pipeMap [][]byte) []position {
	var neighbors []position
	if pos.col > 0 {
		neighbors = append(neighbors, position{pos.row, pos.col - 1})
	}
	if pos.col < len(pipeMap[pos.row])-1 {
		neighbors = append(neighbors, position{pos.row, pos.col + 1})
	}
	return neighbors
}

func getBendNeighbors(shape byte, pos position, pipeMap [][]byte) []position {
	var neighbors []position
	switch shape {
	case pipeBendNorthEast:
		neighbors = appendBendNeighbors(neighbors, pos, -1, 1, pipeMap)
	case pipeBendNorthWest:
		neighbors = appendBendNeighbors(neighbors, pos, -1, -1, pipeMap)
	case pipeBendSouthWest:
		neighbors = appendBendNeighbors(neighbors, pos, 1, -1, pipeMap)
	case pipeBendSouthEast:
		neighbors = appendBendNeighbors(neighbors, pos, 1, 1, pipeMap)
	}
	return neighbors
}

func appendBendNeighbors(neighbors []position, pos position, rowDelta, colDelta int, pipeMap [][]byte) []position {
	if isValidNeighbor(pos.row+rowDelta, pos.col, pipeMap) {
		neighbors = append(neighbors, position{pos.row + rowDelta, pos.col})
	}
	if isValidNeighbor(pos.row, pos.col+colDelta, pipeMap) {
		neighbors = append(neighbors, position{pos.row, pos.col + colDelta})
	}
	return neighbors
}

func isValidNeighbor(row, col int, pipeMap [][]byte) bool {
	return row >= 0 && row < len(pipeMap) && col >= 0 && col < len(pipeMap[row])
}

func findStartingPositionConnectedNeighbors(pipeMap [][]byte, pos position) []position {
	var neighbors []position

	if pos.row > 0 && contains([]byte{pipeVertical, pipeBendSouthEast, pipeBendSouthWest}, pipeMap[pos.row-1][pos.col]) {
		neighbors = append(neighbors, position{pos.row - 1, pos.col})
	}
	if pos.row < len(pipeMap)-1 && contains([]byte{pipeVertical, pipeBendNorthEast, pipeBendNorthWest}, pipeMap[pos.row+1][pos.col]) {
		neighbors = append(neighbors, position{pos.row + 1, pos.col})
	}
	if pos.col > 0 && contains([]byte{pipeHorizontal, pipeBendNorthEast, pipeBendSouthEast}, pipeMap[pos.row][pos.col-1]) {
		neighbors = append(neighbors, position{pos.row, pos.col - 1})
	}
	if pos.col < len(pipeMap[pos.row])-1 && contains([]byte{pipeHorizontal, pipeBendNorthWest, pipeBendSouthWest}, pipeMap[pos.row][pos.col+1]) {
		neighbors = append(neighbors, position{pos.row, pos.col + 1})
	}

	return neighbors
}

func contains(values []byte, value byte) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}
	return false
}

func pipeMapFromReader(r io.Reader) ([][]byte, error) {
	scanner := bufio.NewScanner(r)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	pipeMap := make([][]byte, len(lines))
	for i, line := range lines {
		pipeMap[i] = []byte(line)
	}

	lineLength := len(pipeMap[0])
	for _, line := range pipeMap {
		if len(line) != lineLength {
			return nil, fmt.Errorf("input is not rectangular")
		}
	}

	for row := range pipeMap {
		for col := range pipeMap[row] {
			switch pipeMap[row][col] {
			case
				pipeVertical,
				pipeHorizontal,
				pipeBendNorthEast,
				pipeBendNorthWest,
				pipeBendSouthWest,
				pipeBendSouthEast,
				ground,
				startingPosition:
			default:
				return nil, fmt.Errorf("unknown character %q", pipeMap[row][col])
			}
		}
	}
	return pipeMap, nil
}
