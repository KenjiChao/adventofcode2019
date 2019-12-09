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

func main() {
	lines := util.ReadLines("day9/input.txt")

	input := toIntArray(strings.Split(lines[0], ","))
	fmt.Println("Part 1:", Intcode(input, 1))
	fmt.Println("Part 2:", Intcode(input, 2))
}

func toIntArray(input []string) []int64 {
	output := make([]int64, len(input))
	for i, data := range input {
		val, err := strconv.Atoi(data)
		output[i] = int64(val)
		if err != nil {
			log.Fatal(err)
		}
	}
	return output
}

func Intcode(original []int64, magicInput int64) (output []int64) {
	input := make([]int64, len(original))
	copy(input, original)
	getMemoryPointer := func(index int64) *int64 {
		for int64(len(input)) <= index {
			input = append(input, 0)
		}
		return &input[index]
	}
	var index, relativeBase int64
	for {
		opcode := input[index]%100
		getParameter := func(offset int64) *int64 {
			parameter := input[index+offset]
			mode := input[index] / int64(math.Pow(float64(10), float64(offset+1))) % 10
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
			*a = magicInput
			index += 2
		case Output:
			a := getParameter(1)
			output = append(output, *a)
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
			return
		default:
			log.Fatal("Invalid opcode: ", opcode)
		}

	}
}
