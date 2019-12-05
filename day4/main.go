package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Part1: ", numberOfMatches(152085, 670283, false))
	fmt.Println("Part1: ", numberOfMatches(152085, 670283, true))
}

func numberOfMatches(from, to int, part2 bool) int {
	count := 0
	for i := from; i <= to; i++ {
		if isMatch(i, part2) {
			count++
		}
	}
	return count
}

func isMatch(number int, part2 bool) bool {
	// Rule #1: 6 digits
	if number < 100000 || number >= 1000000 {
		return false
	}
	digits := toDigitArray(number)
	if !part2 {
		return isNonDecreasing(digits) && hasAdjacentSameDigits(digits)
	} else {
		return isNonDecreasing(digits) && hasAdjacentSameDigitsPart2(digits)
	}
}

func isNonDecreasing(digits []int) bool {
	for i := 1; i < len(digits); i++ {
		if digits[i] < digits[i-1] {
			return false
		}
	}
	return true
}

func hasAdjacentSameDigits(digits []int) bool {
	for i := 1; i < len(digits); i++ {
		if digits[i] == digits[i-1] {
			return true
		}
	}
	return false
}

func hasAdjacentSameDigitsPart2(digits []int) bool {
	for i := 1; i < len(digits); i++ {
		if digits[i] == digits[i-1] {
			sameDigits := 2
			for i+1 < len(digits) && digits[i] == digits[i+1] {
				sameDigits++
				i++
			}
			if sameDigits == 2 {
				return true
			}
		}
	}
	return false
}

func toDigitArray(number int) []int {
	numberStr := strconv.Itoa(number)
	digits := make([]int, len(numberStr))
	for i, r := range numberStr {
		digits[i], _ = strconv.Atoi(string(r))
	}
	return digits
}
