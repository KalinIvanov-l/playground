package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Card struct {
	Value string
}

type Hand struct {
	Cards []Card
	Bid   int
	Type  int
	Score int
}

var cardValues = map[string]int{
	"2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "T": 10,
	"J": 11, "Q": 12, "K": 13, "A": 14,
}

func main() {
	file, _ := os.Open("2023/day07/src/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var hands []Hand

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		cardsStr := parts[0]
		var cards []Card
		for _, char := range cardsStr {
			cards = append(cards, Card{Value: string(char)})
		}
		bid, _ := strconv.Atoi(parts[1])
		hands = append(hands, Hand{Cards: cards, Bid: bid})
	}

	for i := range hands {
		hands[i].Type = determineHandType(hands[i])
		hands[i].Score = scoreHand(hands[i])
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].Type == hands[j].Type {
			return scoreHighCards(hands[i].Cards) > scoreHighCards(hands[j].Cards)
		}
		return hands[i].Type > hands[j].Type
	})

	totalWinnings := 0
	for i, hand := range hands {
		rank := len(hands) - i
		totalWinnings += hand.Bid * rank
	}
	fmt.Println(totalWinnings)
}

func scoreHand(hand Hand) int {
	return determineHandType(hand)
}

func scoreHighCards(cards []Card) int {
	score := 0
	for _, card := range cards {
		score = score*100 + cardValues[card.Value]
	}
	return score
}

func determineHandType(hand Hand) int {
	valueCounts := make(map[string]int)
	for _, card := range hand.Cards {
		valueCounts[card.Value]++
	}

	var pairs, threes, fours int
	for _, count := range valueCounts {
		switch count {
		case 2:
			pairs++
		case 3:
			threes++
		case 4:
			fours++
		}
	}

	if len(valueCounts) == 1 {
		return 7
	}
	if fours == 1 {
		return 6
	}
	if threes == 1 && pairs == 1 {
		return 5
	}
	if threes == 1 {
		return 4
	}
	if pairs == 2 {
		return 3
	}
	if pairs == 1 {
		return 2
	}
	return 1
}
