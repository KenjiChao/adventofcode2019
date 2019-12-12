package main

import (
	"fmt"
	util "github.com/adventofcode"
	"math"
	"strconv"
	"strings"
)

type Position struct {
	x, y, z int
}

func (p Position) potentialEnergy() int {
	return int(math.Abs(float64(p.x)) + math.Abs(float64(p.y)) + math.Abs(float64(p.z)))
}

type Velocity struct {
	x, y, z int
}

func (v Velocity) kineticEnergy() int {
	return int(math.Abs(float64(v.x)) + math.Abs(float64(v.y)) + math.Abs(float64(v.z)))
}

type Point struct {
	p Position
	v Velocity
}

func (p Point) energy() int {
	return p.p.potentialEnergy() * p.v.kineticEnergy()
}

type State []Point

func (s *State) nextStep() {
	for i := range *s {
		dx, dy, dz := 0, 0, 0
		for j := range *s {
			if i == j {
				continue
			}
			a, b, c := diffVelocity((*s)[i].p, (*s)[j].p)
			dx += a
			dy += b
			dz += c
		}
		// Update Velocity
		(*s)[i].v.x += dx
		(*s)[i].v.y += dy
		(*s)[i].v.z += dz
	}

	// Update Position
	for i := range *s {
		(*s)[i].p.x += (*s)[i].v.x
		(*s)[i].p.y += (*s)[i].v.y
		(*s)[i].p.z += (*s)[i].v.z
	}
}

func diffVelocity(p, q Position) (int, int, int) {
	return diffStep(p.x, q.x), diffStep(p.y, q.y), diffStep(p.z, q.z)
}

func diffStep(a, b int) int {
	if a == b {
		return 0
	}
	if a > b {
		return -1
	}
	return 1
}

func (s *State) totalEnergy() int {
	total := 0
	for i := range *s {
		total += (*s)[i].energy()
	}

	return total
}

func main() {
	lines := util.ReadLines("day12/input.txt")

	state := NewState(lines)

	for i := 0; i < 1000; i++ {
		state.nextStep()
	}
	fmt.Println(state.totalEnergy())
}

func NewState(lines []string) State {
	var state []Point
	for i := range lines {
		position := strings.Split(lines[i], ",")
		x, _ := strconv.Atoi(strings.Split(position[0], "=")[1])
		y, _ := strconv.Atoi(strings.Split(position[1], "=")[1])
		z, _ := strconv.Atoi(strings.Trim(strings.Split(position[2], "=")[1], ">"))
		state = append(state, Point{
			p: Position{x, y, z},
			v: Velocity{},
		})
	}
	return state
}
