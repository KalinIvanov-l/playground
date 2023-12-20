package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const inputFile = "2023/day20/src/input.txt"

type Module struct {
	Targets    []string
	IsFlipFlop bool
	IsConj     bool
	State      bool
	ConjInputs map[string]bool
}

func ReadModules() map[string]*Module {
	file, _ := os.Open(inputFile)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	modules := make(map[string]*Module)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " -> ")
		if len(parts) != 2 {
			continue
		}

		name := parts[0]
		targets := strings.Split(parts[1], ", ")
		flipFlop := strings.HasPrefix(name, "%")
		conj := strings.HasPrefix(name, "&")

		if flipFlop || conj {
			name = name[1:]
		}

		module := &Module{
			Targets:    targets,
			IsFlipFlop: flipFlop,
			IsConj:     conj,
			State:      false,
		}
		if conj {
			module.ConjInputs = make(map[string]bool)
		}
		modules[name] = module
	}

	for _, module := range modules {
		if module.IsConj {
			for _, target := range module.Targets {
				if targetModule, exists := modules[target]; exists && targetModule.IsConj {
					module.ConjInputs[target] = false
				}
			}
		}
	}
	return modules
}

func SimulatePulses(modules map[string]*Module) int {
	lowPulses, highPulses := 0, 0

	for i := 0; i < 1000; i++ {
		queue := []struct {
			Name   string
			Signal bool
		}{{Name: "broadcaster", Signal: false}}

		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]

			if current.Signal {
				highPulses++
			} else {
				lowPulses++
			}

			module, exists := modules[current.Name]
			if !exists {
				continue
			}

			if module.IsFlipFlop {
				if !current.Signal {
					module.State = !module.State
					newSignal := module.State
					for _, target := range module.Targets {
						queue = append(queue, struct {
							Name   string
							Signal bool
						}{Name: target, Signal: newSignal})
					}
				}
			} else if module.IsConj {
				module.ConjInputs[current.Name] = current.Signal
				allHigh := true
				for _, v := range module.ConjInputs {
					if !v {
						allHigh = false
						break
					}
				}
				newSignal := !allHigh
				for _, target := range module.Targets {
					queue = append(queue, struct {
						Name   string
						Signal bool
					}{Name: target, Signal: newSignal})
				}
			} else {
				for _, target := range module.Targets {
					queue = append(queue, struct {
						Name   string
						Signal bool
					}{Name: target, Signal: current.Signal})
				}
			}
		}
	}
	return lowPulses * highPulses
}

func main() {
	modules := ReadModules()
	result := SimulatePulses(modules)
	fmt.Printf("Part 1: %d\n", result)
}
