package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	x, y int
}

func main() {
	file, err := os.Open("day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Closet Manhattan Distance:", ClosetManhattanDistance(Intersections(strings.Split(lines[0], ","), strings.Split(lines[1], ","))))
}

func ClosetManhattanDistance(intersections []Position) int {
	minDistance := manhattanDistance(intersections[0])
	for _, intersection := range intersections {
		minDistance = int(math.Min(float64(minDistance), float64(manhattanDistance(intersection))))
	}
	return minDistance
}

func Intersections(wire1, wire2 []string) []Position {
	var intersections []Position
	wire1Positions := positions(wire1)
	wire2Positions := positions(wire2)
	for position := range wire1Positions {
		if wire2Positions[position] {
			intersections = append(intersections, position)
		}
	}
	return intersections
}

func positions(wire []string) map[Position]bool {
	positions := make(map[Position]bool)
	currentPosition := &Position{}
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
			positions[*currentPosition] = true
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
