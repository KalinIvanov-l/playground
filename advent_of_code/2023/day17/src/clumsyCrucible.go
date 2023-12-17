package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Limits struct {
	from, to int
}

var (
	LimitsA = Limits{0, 3}
	LimitsB = Limits{3, 10}
)

type Vec struct {
	x, y int
}

func (v *Vec) add(o Vec) {
	v.x += o.x
	v.y += o.y
}

type Dirname int

const (
	Down Dirname = iota
	Right
	Up
	Left
)

var dirs = [4]Vec{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
var perp = [4][2]Dirname{{Left, Right}, {Down, Up}, {Right, Left}, {Up, Down}}

type Node struct {
	loss  int
	links [4][10]Link // Using 10 as the maximum size for simplicity
}

type Link struct {
	n        *Node
	loss     int
	bestLoss int
}

type BFSNode struct {
	n       *Node
	fromDir Dirname
	loss    int
}

func bfs(start BFSNode, limits Limits, level [][]Node) {
	queue := []BFSNode{start}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		for _, nextDirName := range perp[current.fromDir] {
			for i := limits.from; i < limits.to; i++ {
				link := current.n.links[nextDirName][i]
				if link.n == nil {
					break
				}
				if current.loss >= link.bestLoss {
					continue
				}
				link.bestLoss = current.loss
				queue = append(queue, BFSNode{link.n, nextDirName, current.loss + link.loss})
			}
		}
	}
}

func minLoss(n Node, limits Limits) int {
	min := int(^uint(0) >> 1) // Max int
	for _, links := range n.links {
		for i := limits.from; i < limits.to; i++ {
			link := links[i]
			if link.n == nil {
				break
			}
			if link.bestLoss < min {
				min = link.bestLoss
			}
		}
	}
	return min
}

func fillLinks(level [][]Node) {
	maxY, maxX := len(level), len(level[0])
	for y := range level {
		for x := range level[y] {
			cur := &level[y][x]
			for d, dir := range dirs {
				loss := 0
				pos := Vec{x, y}
				for i := 0; i < LimitsB.to; i++ {
					pos.add(dir)
					if pos.x < 0 || pos.y < 0 || pos.x >= maxX || pos.y >= maxY {
						break
					}
					n := &level[pos.y][pos.x]
					loss += n.loss
					cur.links[d][i] = Link{n: n, loss: loss, bestLoss: math.MaxInt32}
				}
			}
		}
	}
}

func main() {
	file, err := os.Open("2023/day17/src/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var level [][]Node
	for scanner.Scan() {
		var row []Node
		for _, c := range scanner.Text() {
			row = append(row, Node{loss: int(c - '0')})
		}
		level = append(level, row)
	}

	if len(level) == 0 || len(level[0]) == 0 {
		fmt.Println("Empty or invalid grid")
		return
	}

	// Fill links
	fillLinks(level)

	// Run BFS and calculate minimum loss for each limit
	startNode := &level[0][0]
	bfs(BFSNode{n: startNode, fromDir: Right, loss: 0}, LimitsA, level)
	bfs(BFSNode{n: startNode, fromDir: Down, loss: 0}, LimitsA, level)
	fmt.Println("Minimum Heat Loss (Limits A):", minLoss(level[len(level)-1][len(level[0])-1], LimitsA))

	bfs(BFSNode{n: startNode, fromDir: Right, loss: 0}, LimitsB, level)
	bfs(BFSNode{n: startNode, fromDir: Down, loss: 0}, LimitsB, level)
	fmt.Println("Minimum Heat Loss (Limits B):", minLoss(level[len(level)-1][len(level[0])-1], LimitsB))
}
