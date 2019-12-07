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
	phaseSettings []int
	amplifiers    []Amplifier
}

func NewAmplifier(input []int) Amplifier {
	return Amplifier{append([]int(nil), input...), true, 0}
}

func NewAmplifiers(input []int, len int) []Amplifier {
	var amplifiers []Amplifier
	for i := 0; i < len; i++ {
		amplifiers = append(amplifiers, NewAmplifier(input))
	}
	return amplifiers
}

func main() {
	lines := util.ReadLines("day7/input.txt")

	input := toIntArray(strings.Split(lines[0], ","))

	maxOutput, phaseSettings := MaxOutput(input, false)
	fmt.Println("Max output signal: ", maxOutput, phaseSettings)
	maxOutput, phaseSettings = MaxOutput(input, true)
	fmt.Println("Feedback Max output signal: ", maxOutput, phaseSettings)
}

func Permutations(nums []int) [][]int {
	result := make([][]int, 0)
	permute(nums, 0, &result)
	return result
}

func permute(nums []int, start int, result *[][]int) {
	if start == len(nums) {
		*result = append(*result, append([]int(nil), nums...))
		return
	}

	for i := start; i < len(nums); i++ {
		nums[start], nums[i] = nums[i], nums[start]
		permute(nums, start+1, result)
		nums[start], nums[i] = nums[i], nums[start]
	}
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

func MaxOutput(input []int, feedbackLoop bool) (int, []int) {
	maxOutput := 0
	var maxPhaseSettings []int

	var phaseSet []int
	if feedbackLoop {
		phaseSet = []int{5, 6, 7, 8, 9}
	} else {
		phaseSet = []int{0, 1, 2, 3, 4}
	}

	for _, phaseSettings := range Permutations(phaseSet) {
		output := initAndRun(input, phaseSettings, feedbackLoop)
		if output > maxOutput {
			maxOutput = output
			maxPhaseSettings = phaseSettings
		}
	}
	return maxOutput, maxPhaseSettings
}

func initAndRun(input []int, phaseSettings []int, feedbackLoop bool) int {
	a := AmplifierSystem{
		phaseSettings: phaseSettings,
		amplifiers:    NewAmplifiers(input, 5),
	}

	output := 0
	if feedbackLoop {
		output = a.runFeedbackLoop()
	} else {
		output = a.run(0)
	}
	return output
}

func (a *AmplifierSystem) run(firstInput int) int {
	output := firstInput
	for i := range a.amplifiers {
		output = a.amplifiers[i].Intcode(a.phaseSettings[i], output)
		if output == -1 {
			break
		}
	}

	return output
}

func (a *AmplifierSystem) runFeedbackLoop() int {
	finalOutput := 0
	for true {
		output := a.run(finalOutput)
		if output == -1 {
			break
		}
		finalOutput = output
	}

	return finalOutput
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
