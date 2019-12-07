package main

import (
	"container/list"
	"fmt"
	util "github.com/adventofcode"
	"strings"
)

func main() {
	lines := util.ReadLines("day6/input.txt")
	fmt.Println(TotalOrbits(lines))
	fmt.Println(Hops(lines))
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

		relations[input[0]] = append(relations[input[0]], input[1])
		degree[input[1]] = degree[input[1]] + 1
	}

	queue := list.New()
	for key, val := range degree {
		if val == 0 {
			queue.PushBack(key)
		}
	}

	totalOrbit, currentOrbit := 0, 0
	for queue.Len() > 0 {
		currentLen := queue.Len()
		totalOrbit += currentOrbit * currentLen
		currentOrbit++

		for i := 0; i < currentLen; i++ {
			planet := queue.Remove(queue.Front()).(string)
			for _, dep := range relations[planet] {
				degree[dep] = degree[dep] - 1
				if degree[dep] == 0 {
					queue.PushBack(dep)
				}
			}
		}
	}
	return totalOrbit
}

const Origin = "YOU"
const Dest = "SAN"

func Hops(lines []string) int {
	relations := make(map[string][]string)

	for _, line := range lines {
		input := strings.Split(line, ")")

		relations[input[0]] = append(relations[input[0]], input[1])
		relations[input[1]] = append(relations[input[1]], input[0])
	}

	queue := list.New()
	queue.PushBack(Origin)
	visited := make(map[string]bool)
	hops := 0

	for queue.Len() > 0 {
		currentLen := queue.Len()
		for i := 0; i < currentLen; i++ {
			planet := queue.Remove(queue.Front()).(string)
			if planet == Dest {
				return hops - 2
			}

			visited[planet] = true
			for _, dep := range relations[planet] {
				if !visited[dep] {
					queue.PushBack(dep)
				}
			}
		}
		hops++
	}

	return 0
}
