package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, _ := os.ReadFile("advent_of_code/2024/day09/input.txt")
	diskMap := string(input)
	blocks, freeSpots := parseDiskMap(diskMap)
	compactBlocks(blocks, freeSpots)

	checksum := calculateChecksum(blocks)
	fmt.Println(checksum)

	blocks, _ = parseDiskMap(diskMap)
	optimizeFileAlignment(blocks)

	checksum = calculateChecksum(blocks)
	fmt.Println(checksum)
}

func parseDiskMap(diskMap string) ([]int, []int) {
	var blocks []int
	var freeSpots []int
	isFile := true
	fileID := 0

	for _, r := range diskMap {
		length, _ := strconv.Atoi(string(r))
		if isFile {
			for i := 0; i < length; i++ {
				blocks = append(blocks, fileID)
			}
			fileID++
		} else {
			for i := 0; i < length; i++ {
				freeSpots = append(freeSpots, len(blocks))
				blocks = append(blocks, -1)
			}
		}
		isFile = !isFile
	}
	return blocks, freeSpots
}

func compactBlocks(blocks []int, freeSpots []int) {
	for _, freePos := range freeSpots {
		if freePos >= len(blocks) {
			break
		}
		for i := len(blocks) - 1; i > freePos; i-- {
			if blocks[i] != -1 {
				blocks[freePos], blocks[i] = blocks[i], -1
				break
			}
		}
	}
	for len(blocks) > 0 && blocks[len(blocks)-1] == -1 {
		blocks = blocks[:len(blocks)-1]
	}
}

func optimizeFileAlignment(blocks []int) {
	var fileStarts []int
	fileID := -1

	for i, block := range blocks {
		if block != -1 && block != fileID {
			fileStarts = append(fileStarts, i)
			fileID = block
		}
	}

	for len(fileStarts) > 0 {
		start := fileStarts[len(fileStarts)-1]
		fileStarts = fileStarts[:len(fileStarts)-1]

		fileID := blocks[start]
		end := start
		for end+1 < len(blocks) && blocks[end+1] == fileID {
			end++
		}
		length := end - start + 1

		dest := findFirstGap(blocks, length)
		if dest == -1 || dest >= start {
			continue
		}
		for i := 0; i < length; i++ {
			blocks[dest+i] = fileID
			blocks[start+i] = -1
		}
	}
}

func findFirstGap(blocks []int, length int) int {
	for i := 0; i+length <= len(blocks); i++ {
		canFit := true
		for j := 0; j < length; j++ {
			if blocks[i+j] != -1 {
				canFit = false
				break
			}
		}
		if canFit {
			return i
		}
	}
	return -1
}

func calculateChecksum(blocks []int) int {
	checksum := 0
	for pos, fileID := range blocks {
		if fileID != -1 {
			checksum += pos * fileID
		}
	}
	return checksum
}
