package main

import (
	"strconv"
	"testing"
)

func TestIntcodePart2(t *testing.T) {
	tests := []struct {
		original   []int
		magicInput int
		want       int
	}{
		{[]int{3,9,8,9,10,9,4,9,99,-1,8}, 8, 1},
		{[]int{3,9,8,9,10,9,4,9,99,-1,8}, 0, 0},
		{[]int{3,9,8,9,10,9,4,9,99,-1,8}, -1, 0},
		{[]int{3,9,7,9,10,9,4,9,99,-1,8}, 7, 1},
		{[]int{3,9,7,9,10,9,4,9,99,-1,8}, 8, 0},
		{[]int{3,9,7,9,10,9,4,9,99,-1,8}, 9, 0},
		{[]int{3,3,1108,-1,8,3,4,3,99}, 8, 1},
		{[]int{3,3,1108,-1,8,3,4,3,99}, 0, 0},
		{[]int{3,3,1108,-1,8,3,4,3,99}, -1, 0},
		{[]int{3,3,1107,-1,8,3,4,3,99}, 7, 1},
		{[]int{3,3,1107,-1,8,3,4,3,99}, 8, 0},
		{[]int{3,3,1107,-1,8,3,4,3,99}, 9, 0},
	}
	for i, tt := range tests {
		t.Run("Test " + strconv.Itoa(i), func(t *testing.T) {
			if got := IntcodePart2(tt.original, tt.magicInput); got != tt.want {
				t.Errorf("IntcodePart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
