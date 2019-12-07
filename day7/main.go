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

type Amplifier struct {
	input           []int
	usePhaseSetting bool
	index           int
}

type AmplifierSystem struct {
	phaseSettings int
	ampA          Amplifier
	ampB          Amplifier
	ampC          Amplifier
	ampD          Amplifier
	ampE          Amplifier
}

func main() {
	lines := util.ReadLines("day7/input.txt")

	input := toIntArray(strings.Split(lines[0], ","))

	maxOutput, phaseSettings := MaxOutput(input, false)
	fmt.Println("Max output signal: ", maxOutput, phaseSettings)
	maxOutput, phaseSettings = MaxOutput(input, true)
	fmt.Println("Feedback Max output signal: ", maxOutput, phaseSettings)
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

func MaxOutput(input []int, feedbackLoop bool) (int, int) {
	maxOutput := 0
	phaseSettings := 0

	openSet := map[int]bool{
		0: true,
		1: true,
		2: true,
		3: true,
		4: true,
	}

	if feedbackLoop {
		openSet = map[int]bool{
			5: true,
			6: true,
			7: true,
			8: true,
			9: true,
		}
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

											output := initAndRun(input, i, j, k, l, m, feedbackLoop)
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

func initAndRun(input []int, i, j, k, l, m int, feedbackLoop bool) int {
	a := &AmplifierSystem{
		phaseSettings: i*10000 + j*1000 + k*100 + l*10 + m,
		ampA:          Amplifier{append([]int(nil), input...), true, 0},
		ampB:          Amplifier{append([]int(nil), input...), true, 0},
		ampC:          Amplifier{append([]int(nil), input...), true, 0},
		ampD:          Amplifier{append([]int(nil), input...), true, 0},
		ampE:          Amplifier{append([]int(nil), input...), true, 0},
	}

	output := 0
	if feedbackLoop {
		output = a.runFeedbackLoop()
	} else {
		output = a.run()
	}
	return output
}

func (a *AmplifierSystem) run() int {
	output1 := a.ampA.Intcode(a.phaseSettings/10000, 0)
	output2 := a.ampB.Intcode(a.phaseSettings/1000%10, output1)
	output3 := a.ampC.Intcode(a.phaseSettings/100%10, output2)
	output4 := a.ampD.Intcode(a.phaseSettings/10%10, output3)
	output5 := a.ampE.Intcode(a.phaseSettings%10, output4)
	return output5
}

func (a *AmplifierSystem) runFeedbackLoop() int {
	output := 0
	input := 0
	for true {
		output1 := a.ampA.Intcode(a.phaseSettings/10000, input)
		if output1 == -1 {
			break
		}
		output2 := a.ampB.Intcode(a.phaseSettings/1000%10, output1)
		if output2 == -1 {
			break
		}
		output3 := a.ampC.Intcode(a.phaseSettings/100%10, output2)
		if output3 == -1 {
			break
		}
		output4 := a.ampD.Intcode(a.phaseSettings/10%10, output3)
		if output4 == -1 {
			break
		}
		output5 := a.ampE.Intcode(a.phaseSettings%10, output4)
		if output5 == -1 {
			break
		}
		input = output5
		output = output5
	}

	return output
}

func (a *Amplifier) Intcode(phaseCode int, inputValue int) int {
	index := a.index
	for index < len(a.input) {
		opcode, parameterMode := a.input[index]%100, a.input[index]/100
		if opcode == OpcodeHalt {
			break
		}
		switch opcode {
		case OpcodeAdd:
			parameterModes := parameterModeList(parameterMode, 3)
			if parameterModes[2] == Immediate {
				log.Fatal("Can't have immediate mode on write operation - Add")
			}
			a.input[a.input[index+3]] = parameterValue(a.input, index+1, parameterModes[0]) + parameterValue(a.input, index+2, parameterModes[1])
			index += 4
		case OpcodeMultiple:
			parameterModes := parameterModeList(parameterMode, 3)
			if parameterModes[2] == Immediate {
				log.Fatal("Can't have immediate mode on write operation - Multiple")
			}
			a.input[a.input[index+3]] = parameterValue(a.input, index+1, parameterModes[0]) * parameterValue(a.input, index+2, parameterModes[1])
			index += 4
		case OpcodeInput:
			parameterModes := parameterModeList(parameterMode, 1)
			if parameterModes[0] == Immediate {
				log.Fatal("Can't have immediate mode on write operation - Input")
			}
			if a.usePhaseSetting {
				a.input[a.input[index+1]] = phaseCode
				a.usePhaseSetting = false
			} else {
				a.input[a.input[index+1]] = inputValue
			}
			index += 2
		case OpcodeOutput:
			parameterModes := parameterModeList(parameterMode, 1)
			output := parameterValue(a.input, index+1, parameterModes[0])
			a.index = index + 2
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
			firstParameter := parameterValue(a.input, index+1, parameterModes[0])
			secondParameter := parameterValue(a.input, index+2, parameterModes[1])
			if firstParameter != 0 {
				index = secondParameter
			} else {
				index += 3
			}
		case OpcodeJumpIfFalse:
			parameterModes := parameterModeList(parameterMode, 2)
			firstParameter := parameterValue(a.input, index+1, parameterModes[0])
			secondParameter := parameterValue(a.input, index+2, parameterModes[1])
			if firstParameter == 0 {
				index = secondParameter
			} else {
				index += 3
			}
		case OpcodeLessThan:
			parameterModes := parameterModeList(parameterMode, 3)
			firstParameter := parameterValue(a.input, index+1, parameterModes[0])
			secondParameter := parameterValue(a.input, index+2, parameterModes[1])
			thirdParameter := a.input[index+3]
			if firstParameter < secondParameter {
				a.input[thirdParameter] = 1
			} else {
				a.input[thirdParameter] = 0
			}
			index += 4
		case OpcodeEquals:
			parameterModes := parameterModeList(parameterMode, 3)
			firstParameter := parameterValue(a.input, index+1, parameterModes[0])
			secondParameter := parameterValue(a.input, index+2, parameterModes[1])
			thirdParameter := a.input[index+3]
			if firstParameter == secondParameter {
				a.input[thirdParameter] = 1
			} else {
				a.input[thirdParameter] = 0
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
