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
const OpcodeJumpIfTrue = 5
const OpcodeJumpIfFalse = 6
const OpcodeLessThan = 7
const OpcodeEquals = 8
const OpcodeHalt = 99

func main() {
	lines := util.ReadLines("day7/input.txt")

	input := toIntArray(strings.Split(lines[0], ","))

	maxOutput, phaseSettings := Amplifier(input)
	fmt.Println("Max output signal: ", maxOutput, phaseSettings)
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

func Amplifier(input []int) (int, int) {
	maxOutput := 0
	phaseSettings := 0

	openSet := map[int]bool{
		0: true,
		1: true,
		2: true,
		3: true,
		4: true,
	}
	for i, oki := range openSet {
		if oki {
			openSet[i] = false
			for j, okj := range openSet {
				if okj {
					openSet[j] = false
					for k, okk := range openSet {
						if okk {
							openSet[k] = false
							for l, okl := range openSet {
								if okl {
									openSet[l] = false
									for m, okm := range openSet {
										if okm {
											openSet[m] = false
											output := runAmplifier(input, i, j, k, l, m)
											if output > maxOutput {
												maxOutput = output
												phaseSettings = i*10000 + j*1000 + k*100 + l*10 + m
											}
											openSet[m] = true
										}
									}
									openSet[l] = true
								}
							}
							openSet[k] = true
						}
					}
					openSet[j] = true
				}
			}
			openSet[i] = true
		}
	}

	return maxOutput, phaseSettings
}

func runAmplifier(input []int, phase1, phase2, phase3, phase4, phase5 int) int {
	output1 := Intcode(input, phase1, 0)
	output2 := Intcode(input, phase2, output1)
	output3 := Intcode(input, phase3, output2)
	output4 := Intcode(input, phase4, output3)
	output5 := Intcode(input, phase5, output4)
	return output5
}

func Intcode(original []int, phaseCode int, inputValue int) int {
	input := append([]int(nil), original...)
	index := 0
	for index < len(input) {
		opcode, parameterMode := input[index]%100, input[index]/100
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
				input[input[index+1]] = phaseCode
				//fmt.Println("OpcodeInput, phaseCode: ", phaseCode)
			} else {
				input[input[index+1]] = inputValue
				//fmt.Println("OpcodeInput, inputValue: ", inputValue)
			}
			index += 2
		case OpcodeOutput:
			parameterModes := parameterModeList(parameterMode, 1)
			output := parameterValue(input, index+1, parameterModes[0])
			return output
			//if output == 0 {
			//	fmt.Println("Test passed! index:", index)
			//} else if index+2 < len(input) && input[index+2]%100 == OpcodeHalt {
			//	return output
			//} else {
			//	log.Fatal("Invalid output value, output: ", output, "index:", index)
			//}
			//index += 2
		case OpcodeJumpIfTrue:
			parameterModes := parameterModeList(parameterMode, 2)
			firstParameter := parameterValue(input, index+1, parameterModes[0])
			secondParameter := parameterValue(input, index+2, parameterModes[1])
			if firstParameter != 0 {
				index = secondParameter
			} else {
				index += 3
			}
		case OpcodeJumpIfFalse:
			parameterModes := parameterModeList(parameterMode, 2)
			firstParameter := parameterValue(input, index+1, parameterModes[0])
			secondParameter := parameterValue(input, index+2, parameterModes[1])
			if firstParameter == 0 {
				index = secondParameter
			} else {
				index += 3
			}
		case OpcodeLessThan:
			parameterModes := parameterModeList(parameterMode, 3)
			firstParameter := parameterValue(input, index+1, parameterModes[0])
			secondParameter := parameterValue(input, index+2, parameterModes[1])
			thirdParameter := input[index+3]
			if firstParameter < secondParameter {
				input[thirdParameter] = 1
			} else {
				input[thirdParameter] = 0
			}
			index += 4
		case OpcodeEquals:
			parameterModes := parameterModeList(parameterMode, 3)
			firstParameter := parameterValue(input, index+1, parameterModes[0])
			secondParameter := parameterValue(input, index+2, parameterModes[1])
			thirdParameter := input[index+3]
			if firstParameter == secondParameter {
				input[thirdParameter] = 1
			} else {
				input[thirdParameter] = 0
			}
			index += 4
		default:
			log.Fatal("Invalid opcode: ", opcode)
		}

	}
	return -1
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
