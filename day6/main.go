package main

import (
	"fmt"
	util "github.com/adventofcode"
	"strings"
)

func main() {
	lines := util.ReadLines("day6/input.txt")
	fmt.Println(TotalOrbits(lines))
}

func TotalOrbits(lines []string) int {
	degree := make(map[string]int)
	relations := make(map[string][]string)

	for _, line := range lines {
		input := strings.Split(line, ")")

		// input[1] orbits input[0]
		if _, ok := degree[input[0]]; !ok {
			degree[input[0]] = 0
		}
		if _, ok := degree[input[1]]; !ok {
			degree[input[1]] = 0
		}
		if _, ok := relations[input[0]]; !ok {
			relations[input[0]] = []string{}
		}
		if _, ok := relations[input[1]]; !ok {
			relations[input[1]] = []string{}
		}

		relations[input[0]] = append(relations[input[0]], input[1])
		degree[input[1]] = degree[input[1]] + 1
	}

	var queue []string
	for key, val := range degree {
		if val == 0 {
			queue = append(queue, key)
		}
	}

	totalOrbit, currentOrbit := 0, 0
	for len(queue) > 0 {
		currentLen := len(queue)
		totalOrbit += currentOrbit * currentLen
		currentOrbit++

		for i := 0; i < currentLen; i++ {
			planet := queue[0]
			queue = queue[1:]
			for _, dep := range relations[planet] {
				degree[dep] = degree[dep] - 1
				if degree[dep] == 0 {
					queue = append(queue, dep)
				}
			}
		}
	}
	return totalOrbit
}
