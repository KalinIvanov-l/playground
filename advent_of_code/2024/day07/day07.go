package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("advent_of_code/2024/day07/input.txt")

	part1, part2 := 0, 0
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		test, _ := strconv.Atoi(strings.Split(s, ": ")[0])
		var ns []int
		if err := json.Unmarshal([]byte("["+strings.ReplaceAll(strings.Split(s, ": ")[1], " ", ",")+"]"), &ns); err != nil {
			panic(err)
		}
		part1 += value(test, ns, false)
		part2 += value(test, ns, true)
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

func value(test int, ns []int, p2 bool) int {
	if len(ns) == 1 {
		if ns[0] == test {
			return test
		}
		return 0
	}
	if len(ns) < 2 {
		return 0
	}
	for _, op := range []func(int, int) int{
		func(a, b int) int { return a + b },
		func(a, b int) int { return a * b },
	} {
		if n := value(test, append([]int{op(ns[0], ns[1])}, ns[2:]...), p2); n != 0 {
			return n
		}
	}

	if p2 {
		if n, _ := strconv.Atoi(fmt.Sprintf("%d%d", ns[0], ns[1])); n != 0 {
			return value(test, append([]int{n}, ns[2:]...), p2)
		}
	}
	return 0
}
