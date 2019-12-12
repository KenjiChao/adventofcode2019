package main

import (
	"fmt"
	util "github.com/adventofcode"
	"math"
	"reflect"
	"strconv"
	"strings"
)

type Tuple struct {
	x, y, z int
}

func (t *Tuple) energy() int {
	return int(math.Abs(float64(t.x)) + math.Abs(float64(t.y)) + math.Abs(float64(t.z)))
}

func (t *Tuple) move(v Tuple) {
	t.x += v.x
	t.y += v.y
	t.z += v.z
}

func (t *Tuple) diff(v Tuple) Tuple {
	return Tuple{
		x: compareInt(t.x, v.x),
		y: compareInt(t.y, v.y),
		z: compareInt(t.z, v.z),
	}
}

func compareInt(a, b int) int {
	if a == b {
		return 0
	}
	if a > b {
		return -1
	}
	return 1
}

type Point struct {
	position Tuple
	velocity Tuple
}

func (p Point) energy() int {
	return p.position.energy() * p.velocity.energy()
}

type State []Point

func (s *State) updateVelocity() {
	for i := range *s {
		for j := range *s {
			if i == j {
				continue
			}
			positionDiff := (*s)[i].position.diff((*s)[j].position)
			(*s)[i].velocity.move(positionDiff)
		}
	}
}

func (s *State) move() {
	for i := range *s {
		(*s)[i].position.move((*s)[i].velocity)
	}
}

func (s *State) nextStep() {
	s.updateVelocity()
	s.move()
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
	initialState := NewState(lines)
	fmt.Println(reflect.DeepEqual(state, initialState))

	for i := 0; i < 1000; i++ {
		state.nextStep()
	}
	fmt.Println(state.totalEnergy())

	state = NewState(lines)
	state.nextStep()
	xsteps := 1
	for ; !isXAxisEquals(state, initialState); xsteps++ {
		state.nextStep()
	}

	state = NewState(lines)
	state.nextStep()
	ysteps := 1
	for ; !isYAxisEquals(state, initialState); ysteps++ {
		state.nextStep()
	}

	state = NewState(lines)
	state.nextStep()
	zsteps := 1
	for ; !isZAxisEquals(state, initialState); zsteps++ {
		state.nextStep()
	}
	fmt.Println(xsteps, ysteps, zsteps)
	fmt.Println(LCM(xsteps, ysteps, zsteps))
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func isXAxisEquals(s1, s2 State) bool {
	for i := range s1 {
		if (s1[i].position.x != s2[i].position.x) || (s1[i].velocity.x != s2[i].velocity.x) {
			return false
		}
	}
	return true
}

func isYAxisEquals(s1, s2 State) bool {
	for i := range s1 {
		if (s1[i].position.y != s2[i].position.y) || (s1[i].velocity.y != s2[i].velocity.y) {
			return false
		}
	}
	return true
}

func isZAxisEquals(s1, s2 State) bool {
	for i := range s1 {
		if (s1[i].position.z != s2[i].position.z) || (s1[i].velocity.z != s2[i].velocity.z) {
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
