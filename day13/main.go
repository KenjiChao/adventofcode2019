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

type Point struct {
	x, y int
}

func (p Point) isEmpty() bool {
	return p.x == 0 && p.y == 0
}

func main() {
	lines := util.ReadLines("day13/input.txt")

	input := toIntArray(strings.Split(lines[0], ","))
	fmt.Println("Part 1:", Intcode(input, false))
	fmt.Println("Part 2:", Intcode(input, true))
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

func Intcode(original []int, playGame bool) int {
	input := make([]int, len(original))
	copy(input, original)
	if playGame {
		input[0] = 2
	}

	var output []int
	blockTiles := 0
	score := 0
	var paddle, ball Point

	getMemoryPointer := func(index int) *int {
		for len(input) <= index {
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
			if paddle.x < ball.x {
				*a = 1
			} else if paddle.x > ball.x {
				*a = -1
			} else {
				*a = 0
			}
			index += 2
		case Output:
			a := getParameter(1)
			output = append(output, *a)
			if len(output)%3 == 0 {
				instruction := output[len(output)-3:]
				x, y, code := instruction[0], instruction[1], instruction[2]
				if x == -1 && y == 0 {
					score = code
				} else {
					switch code {
					case 1:
					case 2:
						blockTiles++
					case 3:
						paddle = Point{x, y}
					case 4:
						ball = Point{x, y}
					}
				}
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
			if playGame {
				return score
			}
			return blockTiles
		default:
			log.Fatal("Invalid opcode: ", opcode)
		}

	}
}
