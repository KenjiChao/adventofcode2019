package main

import (
	"fmt"
	util "github.com/adventofcode"
	"math"
	"strconv"
	"strings"
)

type Tuple struct {
	x, y, z int
}

func (t Tuple) energy() int {
	return int(math.Abs(float64(t.x)) + math.Abs(float64(t.y)) + math.Abs(float64(t.z)))
}

func (t Tuple) diff(v Tuple) Tuple {
	return Tuple{
		x: compareInt(t.x, v.x),
		y: compareInt(t.y, v.y),
		z: compareInt(t.z, v.z),
	}
}

func (t *Tuple) move(v Tuple) {
	t.x += v.x
	t.y += v.y
	t.z += v.z
}

func compareInt(a, b int) int {
	if a == b {
		return 0
	} else if a > b {
		return -1
	} else {
		return 1
	}
}

type Point struct {
	position Tuple
	velocity Tuple
}

func (p Point) energy() int {
	return p.position.energy() * p.velocity.energy()
}

func (p Point) positionDiff(q Point) Tuple {
	return p.position.diff(q.position)
}

func (p Point) xAxis() (int, int) {
	return p.position.x, p.velocity.x
}

func (p Point) yAxis() (int, int) {
	return p.position.y, p.velocity.y
}

func (p Point) zAxis() (int, int) {
	return p.position.z, p.velocity.z
}

func (p *Point) move() {
	p.position.move(p.velocity)
}

func (p *Point) updateVelocity(diff Tuple) {
	p.velocity.move(diff)
}

type State []Point

func (s State) updateVelocity() {
	for i := range s {
		for j := range s {
			if i == j {
				continue
			}
			s[i].updateVelocity(s[i].positionDiff(s[j]))
		}
	}
}

func (s State) move() {
	for i := range s {
		s[i].move()
	}
}

func (s State) nextStep() {
	s.updateVelocity()
	s.move()
}

func (s State) totalEnergy() int {
	total := 0
	for _, p := range s {
		total += p.energy()
	}

	return total
}

func main() {
	lines := util.ReadLines("day12/input.txt")

	initialState := NewState(lines)
	fmt.Println(totalEnergyAfterNSteps(initialState, 1000))

	xSteps := steps(initialState, Point.xAxis)
	ySteps := steps(initialState, Point.yAxis)
	zSteps := steps(initialState, Point.zAxis)
	fmt.Println(xSteps, ySteps, zSteps)
	fmt.Println(lcm(xSteps, ySteps, zSteps))
}

func totalEnergyAfterNSteps(initialState State, n int) int {
	state := append(State{}, initialState...)
	for i := 0; i < n; i++ {
		state.nextStep()
	}
	return state.totalEnergy()
}

func steps(initialState State, fAxis func(p Point) (int, int)) int {
	state := append(State{}, initialState...)
	state.nextStep()
	steps := 1
	for ; !isAxisEquals(state, initialState, fAxis); steps++ {
		state.nextStep()
	}
	return steps
}

func isAxisEquals(s1, s2 State, fAxis func(p Point) (int, int)) bool {
	for i := range s1 {
		s1p, s1v := fAxis(s1[i])
		s2p, s2v := fAxis(s2[i])

		if s1p != s2p || s1v != s2v {
			return false
		}
	}
	return true
}

func NewState(lines []string) State {
	var state []Point
	for i := range lines {
		position := strings.Split(lines[i], ",")
		x, _ := strconv.Atoi(strings.Split(position[0], "=")[1])
		y, _ := strconv.Atoi(strings.Split(position[1], "=")[1])
		z, _ := strconv.Atoi(strings.Trim(strings.Split(position[2], "=")[1], ">"))
		state = append(state, Point{
			position: Tuple{x, y, z},
			velocity: Tuple{},
		})
	}
	return state
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}
