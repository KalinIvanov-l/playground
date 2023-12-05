package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Mapping struct {
	dest, source, dist int
}

func (m *Mapping) getMappedValue(val int) int {
	if val >= m.source && val < m.source+m.dist {
		return (val - m.source) + m.dest
	}
	return -1
}

func convertTextMap(text string) []Mapping {
	var mappings []Mapping
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) == 3 {
			dest, _ := strconv.Atoi(parts[0])
			source, _ := strconv.Atoi(parts[1])
			dist, _ := strconv.Atoi(parts[2])
			mappings = append(mappings, Mapping{dest, source, dist})
		}
	}
	return mappings
}

func readInputFromFile(filePath string) (string, error) {
	file, _ := os.Open(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input strings.Builder
	for scanner.Scan() {
		input.WriteString(scanner.Text() + "\n")
	}
	return input.String(), scanner.Err()
}

func parseSeeds(seedStrings []string) ([]int, error) {
	seeds := make([]int, len(seedStrings))
	for i, s := range seedStrings {
		seed, _ := strconv.Atoi(s)
		seeds[i] = seed
	}
	return seeds, nil
}

// Second part. This part requires more time to execute
func parseSeedRanges(seedStrings []string) ([][2]int, error) {
	seedRanges := make([][2]int, 0, len(seedStrings)/2)
	for i := 0; i < len(seedStrings); i += 2 {
		start, err1 := strconv.Atoi(seedStrings[i])
		length, err2 := strconv.Atoi(seedStrings[i+1])

		if err1 != nil || err2 != nil {
			return nil, fmt.Errorf("invalid seed range: %v, %v", err1, err2)
		}
		seedRanges = append(seedRanges, [2]int{start, length})
	}
	return seedRanges, nil
}

func createMappingSet(parts []string) [][]Mapping {
	mappingSet := make([][]Mapping, len(parts)-1)
	for i, part := range parts[1:] {
		mappingSet[i] = convertTextMap(part)
	}
	return mappingSet
}

// second part
func processSeed(seed int, mappingSet [][]Mapping) int {
	currVal := seed
	for _, mappings := range mappingSet {
		for _, mapEntry := range mappings {
			if newVal := mapEntry.getMappedValue(currVal); newVal != -1 {
				currVal = newVal
				break
			}
		}
	}
	return currVal
}

func main() {
	input, _ := readInputFromFile("2023/day05/src/input.txt")

	parts := strings.Split(input, "\n\n")
	seedStrings := strings.Split(strings.TrimSpace(strings.Split(parts[0], ":")[1]), " ")
	seedRanges, _ := parseSeedRanges(seedStrings)

	mappingSet := createMappingSet(parts)
	minLocation := -1

	for _, rangePair := range seedRanges {
		start, length := rangePair[0], rangePair[1]
		for seed := start; seed < start+length; seed++ {
			processedSeed := processSeed(seed, mappingSet)
			if minLocation == -1 || processedSeed < minLocation {
				minLocation = processedSeed
			}
		}
	}
	fmt.Println("Lowest location number:", minLocation)
}
