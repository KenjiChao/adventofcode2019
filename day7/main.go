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
		if output == OpcodeHalt {
			break
		}
	}

	return output
}

func (a *AmplifierSystem) runFeedbackLoop() int {
	finalOutput := 0
	for true {
		output := a.run(finalOutput)
		if output == OpcodeHalt {
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
			index = a.Add(index, parameterMode)
		case OpcodeMultiple:
			index = a.Multiple(index, parameterMode)
		case OpcodeInput:
			if a.usePhaseSetting {
				a.input[a.input[index+1]] = phaseCode
				a.usePhaseSetting = false
			} else {
				a.input[a.input[index+1]] = inputValue
			}
			index += 2
		case OpcodeOutput:
			output := a.parameters(index+1, 1, parameterMode)[0]
			a.index = index + 2
			return output
		case OpcodeJumpIfTrue:
			index = a.JumpIfTrue(index, parameterMode)
		case OpcodeJumpIfFalse:
			index = a.JumpIfFalse(index, parameterMode)
		case OpcodeLessThan:
			index = a.LessThan(index, parameterMode)
		case OpcodeEquals:
			index = a.Equals(index, parameterMode)
		default:
			log.Fatal("Invalid opcode: ", opcode)
		}

	}
	return OpcodeHalt
}

func (a *Amplifier) Add(index, parameterMode int) int {
	parameterValues := a.parameters(index+1, 2, parameterMode)
	a.input[a.input[index+3]] = parameterValues[0] + parameterValues[1]
	return index + 4
}

func (a *Amplifier) Multiple(index, parameterMode int) int {
	parameterValues := a.parameters(index+1, 2, parameterMode)
	a.input[a.input[index+3]] = parameterValues[0] * parameterValues[1]
	return index + 4
}

func (a *Amplifier) JumpIfTrue(index, parameterMode int) int {
	parameterValues := a.parameters(index+1, 2, parameterMode)
	if parameterValues[0] != 0 {
		return parameterValues[1]
	} else {
		return index + 3
	}
}

func (a *Amplifier) JumpIfFalse(index, parameterMode int) int {
	parameterValues := a.parameters(index+1, 2, parameterMode)
	if parameterValues[0] == 0 {
		return parameterValues[1]
	} else {
		return index + 3
	}
}

func (a *Amplifier) LessThan(index, parameterMode int) int {
	parameterValues := a.parameters(index+1, 2, parameterMode)
	firstParameter, secondParameter := parameterValues[0], parameterValues[1]
	thirdParameter := a.input[index+3]
	if firstParameter < secondParameter {
		a.input[thirdParameter] = 1
	} else {
		a.input[thirdParameter] = 0
	}
	return index + 4
}

func (a *Amplifier) Equals(index, parameterMode int) int {
	parameterValues := a.parameters(index+1, 2, parameterMode)
	firstParameter, secondParameter := parameterValues[0], parameterValues[1]
	thirdParameter := a.input[index+3]
	if firstParameter == secondParameter {
		a.input[thirdParameter] = 1
	} else {
		a.input[thirdParameter] = 0
	}
	return index + 4
}

func (a *Amplifier) parameters(startingIndex, numberOfParameters, parameterMode int) []int {
	parameterValues := make([]int, 0)
	for i := 0; i < numberOfParameters; i++ {
		if parameterMode%10 == 1 {
			parameterValues = append(parameterValues, a.input[startingIndex+i])
		} else {
			parameterValues = append(parameterValues, a.input[a.input[startingIndex+i]])
		}
		parameterMode = parameterMode / 10
	}

	return parameterValues
}
