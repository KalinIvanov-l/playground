package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
)

type Position struct {
	x, y int
}

type State struct {
	pos   Position
	dir   Position
	steps int
	cost  int
}

type PriorityQueue []*State

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*State)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

func main() {
	file, err := os.Open("2023/day17/src/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	costs := make(map[Position]int)
	maxX, maxY := 0, 0
	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		for x, c := range line {
			costs[Position{x, y}] = int(c - '0')
			maxX = max(maxX, x)
		}
		maxY = max(maxY, y)
	}

	leastHeatLoss := solve(costs, Position{maxX, maxY})
	fmt.Println("Least heat loss:", leastHeatLoss)
}

func solve(costs map[Position]int, target Position) int {
	directions := []Position{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	costMap := make(map[State]int)
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	for _, d := range directions {
		startState := State{Position{0, 0}, d, 1, 0}
		costMap[startState] = 0
		heap.Push(&pq, &startState)
	}

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*State)
		if current.pos == target && current.steps <= 3 {
			return costMap[*current]
		}

		for _, d := range directions {
			newPos := Position{current.pos.x + d.x, current.pos.y + d.y}
			if _, exists := costs[newPos]; !exists {
				continue
			}

			newSteps := 1
			if current.dir == d {
				newSteps = current.steps + 1
			}

			if newSteps > 3 {
				continue
			}

			newCost := costMap[*current] + costs[newPos]
			newState := State{newPos, d, newSteps, newCost}
			if c, exists := costMap[newState]; !exists || newCost < c {
				costMap[newState] = newCost
				heap.Push(&pq, &newState)
			}
		}
	}

	return math.MaxInt
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
