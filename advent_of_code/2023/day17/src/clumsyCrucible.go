package main

import (
	"container/heap"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// State represents the current state in the pathfinding process.
type State struct {
	loss, chain, direction int
	pos                    complex128
}

// PriorityQueue implements heap.Interface and holds States.
type PriorityQueue []*State

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].loss < pq[j].loss
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*State))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// readGrid reads the grid from a file and returns a map with the heat loss for each coordinate.
func readGrid(filename string) (map[complex128]int, complex128) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(bytes), "\n")
	grid := make(map[complex128]int)
	var maxX, maxY int
	for y, line := range lines {
		for x, char := range line {
			value, err := strconv.Atoi(string(char))
			if err != nil {
				panic(err)
			}
			grid[complex(float64(x), float64(y))] = value
			if x > maxX {
				maxX = x
			}
		}
		maxY = y
	}
	end := complex(float64(maxX), float64(maxY))
	return grid, end
}

// findPath implements the pathfinding logic using a priority queue.
func findPath(grid map[complex128]int, start, end complex128, a, b, c int) int {
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &State{0, 1, 0, start})
	best := make(map[complex128]map[int]map[int]int)

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*State)
		if current.pos == end && current.chain >= a {
			return current.loss
		}

		if _, ok := best[current.pos]; !ok {
			best[current.pos] = make(map[int]map[int]int)
		}
		if _, ok := best[current.pos][current.direction]; !ok {
			best[current.pos][current.direction] = make(map[int]int)
		}
		if val, ok := best[current.pos][current.direction][current.chain]; ok && val <= current.loss {
			continue
		}
		best[current.pos][current.direction][current.chain] = current.loss

		directions := []complex128{1, 1i, -1, -1i}
		for e, d := range directions {
			nextChain := 1
			if e%2 == 0 {
				nextChain = current.chain + 1
			}
			if (e%2 == 0 && current.chain < b) || (e%2 != 0 && current.chain == c) {
				continue
			}
			neighbour := current.pos + d
			if loss, ok := grid[neighbour]; ok {
				heap.Push(pq, &State{current.loss + loss, nextChain, e, neighbour})
			}
		}
	}

	return -1
}

func main() {
	grid, end := readGrid("2023/day17/src/input.txt")
	start := complex(0, 0)
	fmt.Println("Minimum heat loss:", findPath(grid, start, end, 0, 0, 3))
	fmt.Println("Minimum heat loss:", findPath(grid, start, end, 4, 4, 10))
}
