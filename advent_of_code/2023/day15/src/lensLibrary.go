package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Lens struct {
	Label       string
	FocalLength int
}

func main() {
	data, err := ioutil.ReadFile("2023/day15/src/input.txt")
	if err != nil {
		panic(err)
	}
	steps := strings.Split(strings.TrimSpace(string(data)), ",")
	boxes := make([][]Lens, 256) // Initialize boxes

	for _, step := range steps {
		processStep(step, &boxes)
	}
	totalFocusPower := 0
	for boxIndex, box := range boxes {
		for slotIndex, lens := range box {
			totalFocusPower += (boxIndex + 1) * (slotIndex + 1) * lens.FocalLength
		}
	}
	fmt.Println("Total focusing power:", totalFocusPower)
}

func processStep(step string, boxes *[][]Lens) {
	label := step
	if strings.Contains(step, "=") {
		parts := strings.Split(step, "=")
		label = parts[0]
		focalLength := atomic(parts[1])
		boxIndex := hash(label)
		replaceOrInsertLens(label, focalLength, &(*boxes)[boxIndex])
	} else if strings.Contains(step, "-") {
		label = strings.TrimSuffix(label, "-")
		boxIndex := hash(label)
		removeLens(label, &(*boxes)[boxIndex])
	}
}

func hash(s string) int {
	currentValue := 0
	for _, c := range s {
		currentValue = (currentValue + int(c)) * 17 % 256
	}
	return currentValue
}

func replaceOrInsertLens(label string, focalLength int, box *[]Lens) {
	for i, lens := range *box {
		if lens.Label == label {
			(*box)[i].FocalLength = focalLength
			return
		}
	}
	*box = append(*box, Lens{Label: label, FocalLength: focalLength})
}

func removeLens(label string, box *[]Lens) {
	for i, lens := range *box {
		if lens.Label == label {
			*box = append((*box)[:i], (*box)[i+1:]...)
			return
		}
	}
}

func atomic(s string) int {
	res := 0
	for _, c := range s {
		res = res*10 + int(c-'0')
	}
	return res
}
