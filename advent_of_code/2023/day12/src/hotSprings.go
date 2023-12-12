package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type key struct {
	lava string
	nums string
	has  bool
}

var cache map[key]int

func sliceToString(slice []int) string {
	return strings.Trim(strings.Replace(fmt.Sprint(slice), " ", ",", -1), "[]")
}

func solve(lava string, nums []int, has bool) int {
	k := key{lava, sliceToString(nums), has}
	if val, found := cache[k]; found {
		return val
	}
	if len(lava) == 0 {
		if len(nums) == 0 || (len(nums) == 1 && nums[0] == 0) {
			return 1
		}
		return 0
	}

	var result int
	switch {
	case lava[0] == '#' && len(nums) > 0 && nums[0] > 0:
		numsCopy := append([]int{nums[0] - 1}, nums[1:]...)
		result = solve(lava[1:], numsCopy, true)
	case lava[0] == '.' && has && len(nums) > 0 && nums[0] == 0:
		result = solve(lava[1:], nums[1:], false)
	case lava[0] == '.' && !has:
		result = solve(lava[1:], nums, false)
	case lava[0] == '?':
		result = solve("#"+lava[1:], nums, has) + solve("."+lava[1:], nums, has)
	}
	cache[k] = result
	return result
}

func unfoldConditions(lava string, nums []int) (string, []int) {
	unfoldedLava := strings.Join([]string{lava, lava, lava, lava, lava}, "?")
	var unfoldedNums []int
	for i := 0; i < 5; i++ {
		unfoldedNums = append(unfoldedNums, nums...)
	}
	return unfoldedLava, unfoldedNums
}

func main() {
	cache = make(map[key]int)
	file, _ := os.Open("2023/day12/src/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		lava, numsStr := parts[0], strings.Split(parts[1], ",")
		var nums []int
		for _, s := range numsStr {
			n, _ := strconv.Atoi(s)
			nums = append(nums, n)
		}
		unfoldedLava, unfoldedNums := unfoldConditions(lava, nums)
		total += solve(unfoldedLava, unfoldedNums, false)
	}
	fmt.Println("Total number of arrangements after unfolding:", total)
}
