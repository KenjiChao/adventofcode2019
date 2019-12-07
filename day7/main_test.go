package main

import (
	"strconv"
	"testing"
)

func TestAmplifier(t *testing.T) {
	tests := []struct {
		input    []int
		output   int
		settings int
	}{
		{[]int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}, 43210, 43210},
	}
	for i, tt := range tests {
		t.Run("Test "+strconv.Itoa(i), func(t *testing.T) {
			if got1, got2 := Amplifier(tt.input); got1 != tt.output || got2 != tt.settings {
				t.Errorf("Amplifier() = %v, %v, want %v, %v", got1, got2, tt.output, tt.settings)
			}
		})
	}
}
