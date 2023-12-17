package main

import (
	"container/heap"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type State struct {
	pos   complex128
	dir   complex128
	loss  int
	index int
}

type PriorityQueue []*State

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].loss < pq[j].loss
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*State)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func readInput(filename string) map[complex128]int {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(bytes), "\n")
	grid := make(map[complex128]int)
	for y, line := range lines {
		for x, char := range line {
			value, err := strconv.Atoi(string(char))
			if err != nil {
				panic(err)
			}
			grid[complex(float64(x), float64(y))] = value
		}
	}
	return grid
}

func main() {
	grid := readInput("2023/day17/src/input.txt")
	target := complex128(0)
	for pos := range grid {
		if real(pos) > real(target) || imag(pos) > imag(target) {
			target = pos
		}
	}

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &State{0, 1, 0, 0})
	heap.Push(&pq, &State{0, 1i, 0, 0})

	visited := make(map[complex128]map[complex128]bool)

	for pq.Len() > 0 {
		state := heap.Pop(&pq).(*State)
		if state.pos == target {
			fmt.Println("Minimum heat loss:", state.loss)
			return
		}
		if _, ok := visited[state.pos]; !ok {
			visited[state.pos] = make(map[complex128]bool)
		}
		if _, visited := visited[state.pos][state.dir]; visited {
			continue
		}
		visited[state.pos][state.dir] = true

		for _, dirChange := range []complex128{1, -1} {
			newDir := state.dir * dirChange
			newPos := state.pos + newDir
			if _, exists := grid[newPos]; exists {
				newLoss := state.loss + grid[newPos]
				heap.Push(&pq, &State{newPos, newDir, newLoss, 0})
			}
		}
	}
}
