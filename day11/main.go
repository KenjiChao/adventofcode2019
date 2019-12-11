package main

import (
	"fmt"
	util "github.com/adventofcode"
	"log"
	"math"
	"strconv"
	"strings"
)

const Add = 1
const Multiple = 2
const Input = 3
const Output = 4
const JumpIfTrue = 5
const JumpIfFalse = 6
const LessThan = 7
const Equals = 8
const RelativeBaseOffset = 9
const Halt = 99

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

type Point struct {
	x, y int
}

type Position struct {
	point     Point
	direction Direction
}

func main() {
	lines := util.ReadLines("day11/input.txt")

	input := toIntArray(strings.Split(lines[0], ","))
	fmt.Println("Part 1:", Intcode(input, Position{
		point:     Point{},
		direction: Up,
	}, 0))
	fmt.Println("Part 2:", Intcode(input, Position{
		point:     Point{},
		direction: Up,
	}, 1))
}

func toIntArray(input []string) []int {
	output := make([]int, len(input))
	for i, data := range input {
		val, err := strconv.Atoi(data)
		output[i] = int(val)
		if err != nil {
			log.Fatal(err)
		}
	}
	return output
}

func Intcode(original []int, position Position, defaultColor int) int {
	input := make([]int, len(original))
	copy(input, original)
	var output []int

	colorMap := make(map[Point]int)
	stepMap := map[Direction]Point{
		Up:    {0, 1,},
		Right: {1, 0,},
		Down:  {0, -1,},
		Left:  {-1, 0,},
	}

	pointRange := []int{0, 0, 0, 0}

	getMemoryPointer := func(index int) *int {
		for int(len(input)) <= index {
			input = append(input, 0)
		}
		return &input[index]
	}
	var index, relativeBase int
	for {
		opcode := input[index] % 100
		getParameter := func(offset int) *int {
			parameter := input[index+offset]
			mode := input[index] / int(math.Pow(float64(10), float64(offset+1))) % 10
			switch mode {
			case 0: // position mode
				return getMemoryPointer(parameter)
			case 1: // immediate mode
				return &parameter
			case 2: // relative mode
				return getMemoryPointer(relativeBase + parameter)
			default:
				panic(fmt.Sprintf("fault: invalid parameter mode: ip=%d instruction=%d offset=%d mode=%d", index, input[index], offset, mode))
			}
		}

		switch opcode {
		case Add:
			a, b, c := getParameter(1), getParameter(2), getParameter(3)
			*c = *a + *b
			index += 4
		case Multiple:
			a, b, c := getParameter(1), getParameter(2), getParameter(3)
			*c = *a * *b
			index += 4
		case Input:
			a := getParameter(1)
			var color int
			val, ok := colorMap[position.point]
			if ok {
				color = val
			} else {
				color = defaultColor
			}

			*a = color
			index += 2
		case Output:
			a := getParameter(1)
			output = append(output, *a)
			if len(output)%2 == 0 {
				// fmt.Println("current:", position)
				instruction := output[len(output)-2:]
				// fmt.Println("instruction:", instruction)
				colorMap[position.point] = instruction[0]
				nextDirection := nextDirection(position.direction, instruction[1])
				position = Position{
					point: Point{
						x: position.point.x + stepMap[nextDirection].x,
						y: position.point.y + stepMap[nextDirection].y,
					},
					direction: nextDirection,
				}

				pointRange[0] = int(math.Min(float64(pointRange[0]), float64(position.point.x)))
				pointRange[1] = int(math.Max(float64(pointRange[1]), float64(position.point.x)))
				pointRange[2] = int(math.Min(float64(pointRange[2]), float64(position.point.y)))
				pointRange[3] = int(math.Max(float64(pointRange[3]), float64(position.point.y)))
				//fmt.Println("new:", position)
			}
			index += 2
		case JumpIfTrue:
			a, b := getParameter(1), getParameter(2)

			if *a != 0 {
				index = *b
			} else {
				index += 3
			}
		case JumpIfFalse:
			a, b := getParameter(1), getParameter(2)

			if *a == 0 {
				index = *b
			} else {
				index += 3
			}
		case LessThan:
			a, b, c := getParameter(1), getParameter(2), getParameter(3)
			if *a < *b {
				*c = 1
			} else {
				*c = 0
			}
			index += 4
		case Equals:
			a, b, c := getParameter(1), getParameter(2), getParameter(3)
			if *a == *b {
				*c = 1
			} else {
				*c = 0
			}
			index += 4
		case RelativeBaseOffset:
			a := getParameter(1)
			relativeBase += *a
			index += 2
		case Halt:
			xRange := pointRange[1] - pointRange[0] + 1
			yRange := pointRange[3] - pointRange[2] + 1
			grid := make([][]int, yRange)
			for i := range grid {
				grid[i] = make([]int, xRange)
			}
			offset := Point{
				x: -pointRange[0],
				y: -pointRange[2],
			}

			for point, color := range colorMap {
				grid[point.y+offset.y][point.x+offset.x] = color
			}

			sb := strings.Builder{}
			for i := range grid {
				for j := range grid[i] {
					if grid[i][j] == 1 {
						sb.WriteRune('O')
					} else {
						sb.WriteRune(' ')
					}
				}
				sb.WriteRune('\n')
			}
			fmt.Println(sb.String())

			return len(colorMap)
		default:
			log.Fatal("Invalid opcode: ", opcode)
		}

	}
}

func nextDirection(d Direction, code int) Direction {
	newDirection := code*2 - 1 + int(d)
	if newDirection < 0 {
		newDirection += 4
	}
	if newDirection >= 4 {
		newDirection -= 4
	}
	return Direction(newDirection)
}
