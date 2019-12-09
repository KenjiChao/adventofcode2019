package main

import (
	"reflect"
	"strconv"
	"testing"
)

func TestIntcode(t *testing.T) {
	tests := []struct {
		original   []int
		magicInput int
		want       []int
	}{
		{[]int{104, 1125899906842624, 99}, 1, []int{1125899906842624}},
		{[]int{1102, 34915192, 34915192, 7, 4, 7, 99, 0}, 1, []int{1219070632396864}},
		{[]int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}, 1, []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}},
	}
	for i, tt := range tests {
		t.Run("Test "+strconv.Itoa(i), func(t *testing.T) {
			if got := Intcode(tt.original, tt.magicInput); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intcode() = %v, want %v", got, tt.want)
			}
		})
	}
}
