package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(numberOfMatches(152085, 670283))
}

func numberOfMatches(from, to int) int {
	count := 0
	for i := from; i <= to; i++ {
		if isMatch(i) {
			count++
		}
	}
	return count
}

func isMatch(number int) bool {
	// Rule #1: 6 digits
	if number < 100000 || number >= 1000000 {
		return false
	}
	digits := toDigitArray(number)
	return isNonDecreasing(digits) && hasAdjacentSameDigits(digits)
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

func toDigitArray(number int) []int {
	numberStr := strconv.Itoa(number)
	digits := make([]int, len(numberStr))
	for i, r := range numberStr {
		digits[i], _ = strconv.Atoi(string(r))
	}
	return digits
}
