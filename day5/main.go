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
const OpcodeHalt = 99

const MagicNumber = 19690720

func main() {
	lines := util.ReadLines("day2/input.txt")

	input := toIntArray(strings.Split(lines[0], ","))
	input[1] = 12
	input[2] = 2
	fmt.Println("Value at position 0:", Intcode(input)[0])

	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			input[1], input[2] = i, j
			if Intcode(input)[0] == MagicNumber {
				fmt.Println("Find the match! noun:", i, "verb:", j)
				fmt.Println("100 * noun + verb =", 100*i+j)
				return
			}
		}
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

func Intcode(original []int) []int {
	input := append([]int(nil), original...)
	for index, opcode := 0, input[0]; opcode != OpcodeHalt; index, opcode = index+4, input[index+4] {
		switch opcode {
		case OpcodeAdd:
			input[input[index+3]] = input[input[index+1]] + input[input[index+2]]
		case OpcodeMultiple:
			input[input[index+3]] = input[input[index+1]] * input[input[index+2]]
		default:
			log.Fatal("Invalid opcode: ", opcode)
		}
	}
	return input
}
