package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand string
type Play struct {
	hand Hand
	bid  int
}

func main() {
	file, _ := os.Open("2023/day07/src/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var plays []Play

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		bid, _ := strconv.Atoi(parts[1])
		plays = append(plays, Play{hand: Hand(parts[0]), bid: bid})
	}

	cardValues := []rune{'J', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'Q', 'K', 'A'}
	sort.Slice(plays, func(i, j int) bool {
		return plays[j].hand.beats(plays[i].hand, cardValues, true)
	})

	total := 0
	for idx, play := range plays {
		total += play.bid * (idx + 1)
	}
	fmt.Println(total)
}

func (h1 Hand) beats(h2 Hand, cardValues []rune, joker bool) bool {
	var val1, val2 int
	if joker {
		val1 = getHandValueWithJoker(h1, cardValues)
		val2 = getHandValueWithJoker(h2, cardValues)
	} else {
		val1 = getHandValue(h1)
		val2 = getHandValue(h2)
	}

	if val1 == val2 {
		for i := 0; i < len(h1); i++ {
			cVal1 := getCardValue(rune(h1[i]), cardValues)
			cVal2 := getCardValue(rune(h2[i]), cardValues)
			if cVal1 != cVal2 {
				return cVal1 > cVal2
			}
		}
		return false
	}
	return val1 > val2
}

func getCardValue(r rune, cardValues []rune) int {
	for idx, v := range cardValues {
		if v == r {
			return idx
		}
	}
	return -1
}

func getHandValueWithJoker(h Hand, cardValues []rune) (max int) {
	for _, v := range cardValues {
		tmp := strings.Replace(string(h), "J", string(v), -1)
		val := getHandValue(Hand(tmp))
		if val > max {
			max = val
		}
	}
	return
}

func getHandValue(h Hand) int {
	vals := make(map[rune]int)
	for _, card := range h {
		vals[card]++
	}
	var count []int
	for _, v := range vals {
		count = append(count, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(count)))

	switch count[0] {
	case 5:
		return 7
	case 4:
		return 6
	case 3:
		if count[1] == 2 {
			return 5
		}
		return 4
	case 2:
		if count[1] == 2 {
			return 3
		}
		return 2
	default:
		return 1
	}
}
