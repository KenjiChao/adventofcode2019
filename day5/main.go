package main

import (
	"fmt"
	util "github.com/adventofcode"
	"log"
	"strconv"
	"strings"
)

const OpcodeAdd = 1
const OpcodeMultiple = 2
const OpcodeInput = 3
const OpcodeOutput = 4
const OpcodeHalt = 99

const MagicInput = 1

func main() {
	lines := util.ReadLines("day5/input.txt")

	input := toIntArray(strings.Split(lines[0], ","))
	Intcode(input)
}

func toIntArray(input []string) []int {
	output := make([]int, len(input))
	var err error
	for i, data := range input {
		output[i], err = strconv.Atoi(data)
		if err != nil {
			log.Fatal(err)
		}
	}
	return output
}

func Intcode(original []int) []int {
	input := append([]int(nil), original...)
	index := 0
	for index < len(input) {
		opcode, parameterMode := input[index]%100, input[index]/100
		//fmt.Println("index:", index, "opcode:", opcode, "parameterMode:", parameterMode)
		//fmt.Println(input[index : index+5])
		if opcode == OpcodeHalt {
			break
		}
		switch opcode {
		case OpcodeAdd:
			parameterModes := parameterModeList(parameterMode, 3)
			if parameterModes[2] == Immediate {
				log.Fatal("Can't have immediate mode on write operation - Add")
			}
			input[input[index+3]] = parameterValue(input, index+1, parameterModes[0]) + parameterValue(input, index+2, parameterModes[1])
			index += 4
		case OpcodeMultiple:
			parameterModes := parameterModeList(parameterMode, 3)
			if parameterModes[2] == Immediate {
				log.Fatal("Can't have immediate mode on write operation - Multiple")
			}
			input[input[index+3]] = parameterValue(input, index+1, parameterModes[0]) * parameterValue(input, index+2, parameterModes[1])
			index += 4
		case OpcodeInput:
			parameterModes := parameterModeList(parameterMode, 1)
			if parameterModes[0] == Immediate {
				log.Fatal("Can't have immediate mode on write operation - Input")
			}
			if index == 0 {
				input[input[index+1]] = MagicInput
			} else {
				log.Fatal("no magic input")
			}
			index += 2
		case OpcodeOutput:
			parameterModes := parameterModeList(parameterMode, 1)
			output := parameterValue(input, index+1, parameterModes[0])
			if output == 0 {
				fmt.Println("Test passed! index:", index)
			} else if index+2 < len(input) && input[index+2]%100 == OpcodeHalt {
				fmt.Println("Find the code: ", output, "index:", index)
			} else {
				log.Fatal("Invalid output value, output: ", output, "index:", index)
			}
			index += 2
		default:
			log.Fatal("Invalid opcode: ", opcode)
		}

	}
	return input
}

func parameterValue(input []int, index int, mode ParameterMode) int {
	switch mode {
	case Position:
		return input[input[index]]
	case Immediate:
		return input[index]
	default:
		log.Fatal("Invalid mode:", mode)
		return -1
	}
}

func parameterModeList(parameterMode, len int) []ParameterMode {
	parameterModeList := make([]ParameterMode, len)
	for i := 0; i < len; i++ {
		currentMode := Position
		if parameterMode%10 == 1 {
			currentMode = Immediate
		} else if parameterMode%10 != 0 {
			log.Fatal("Invalid parameter mode:", parameterMode%10)
		}
		parameterModeList[i] = currentMode
		parameterMode = parameterMode / 10
	}
	return parameterModeList
}

type ParameterMode int

const (
	Position ParameterMode = iota
	Immediate
)
