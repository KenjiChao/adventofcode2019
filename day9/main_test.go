package main

import (
	"reflect"
	"strconv"
	"testing"
)

func TestIntcode(t *testing.T) {
	tests := []struct {
		original   []int64
		magicInput int64
		want       []int64
	}{
		{[]int64{104, 1125899906842624, 99}, 1, []int64{1125899906842624}},
		{[]int64{1102, 34915192, 34915192, 7, 4, 7, 99, 0}, 1, []int64{1219070632396864}},
		{[]int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}, 1, []int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}},
	}
	for i, tt := range tests {
		t.Run("Test "+strconv.Itoa(i), func(t *testing.T) {
			if got := Intcode(tt.original, tt.magicInput); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intcode() = %v, want %v", got, tt.want)
			}
		})
	}
}
