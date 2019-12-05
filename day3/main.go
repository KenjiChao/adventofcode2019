package main

import (
	"fmt"
	util "github.com/adventofcode"
	"log"
	"math"
	"strconv"
	"strings"
)

type Position struct {
	x, y int
}

func main() {
	lines := util.ReadLines("day3/input.txt")

	fmt.Println("Closet Manhattan Distance:", ClosetManhattanDistance(Intersections(strings.Split(lines[0], ","), strings.Split(lines[1], ","))))
	fmt.Println("Fewest Combined Steps:", FewestCombinedSteps(strings.Split(lines[0], ","), strings.Split(lines[1], ",")))
}

func ClosetManhattanDistance(intersections []Position) int {
	minDistance := manhattanDistance(intersections[0])
	for _, intersection := range intersections {
		minDistance = int(math.Min(float64(minDistance), float64(manhattanDistance(intersection))))
	}
	return minDistance
}

func FewestCombinedSteps(wire1, wire2 []string) int {
	fewestSteps := math.MaxInt64
	for _, steps := range IntersectionsWithSteps(wire1, wire2) {
		fewestSteps = int(math.Min(float64(fewestSteps), float64(steps)))
	}
	return fewestSteps
}

func IntersectionsWithSteps(wire1, wire2 []string) map[Position]int {
	intersectionsWithSteps := make(map[Position]int)
	wire1Positions := positions(wire1)
	wire2Positions := positions(wire2)
	for position := range wire1Positions {
		if wire2Positions[position] != 0 {
			intersectionsWithSteps[position] = wire1Positions[position] + wire2Positions[position]
		}
	}
	return intersectionsWithSteps
}

func Intersections(wire1, wire2 []string) []Position {
	var intersections []Position
	wire1Positions := positions(wire1)
	wire2Positions := positions(wire2)
	for position := range wire1Positions {
		if wire2Positions[position] != 0 {
			intersections = append(intersections, position)
		}
	}
	return intersections
}

func positions(wire []string) map[Position]int {
	positions := make(map[Position]int)
	currentPosition := &Position{}
	currentSteps := 0
	for _, op := range wire {
		direction := op[0]
		steps, err := strconv.Atoi(op[1:])
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < steps; i++ {
			switch direction {
			case 'R':
				currentPosition.Right()
			case 'L':
				currentPosition.Left()
			case 'U':
				currentPosition.Up()
			case 'D':
				currentPosition.Down()
			default:
				log.Fatal("Invalid operation")
			}
			currentSteps++
			if positions[*currentPosition] == 0 {
				positions[*currentPosition] = currentSteps
			}
		}
	}
	return positions
}

func manhattanDistance(position Position) int {
	return int(math.Abs(float64(position.x)) + math.Abs(float64(position.y)))
}

func (p *Position) Right() {
	p.x++
}

func (p *Position) Left() {
	p.x--
}

func (p *Position) Up() {
	p.y++
}

func (p *Position) Down() {
	p.y--
}
