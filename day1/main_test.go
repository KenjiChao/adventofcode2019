package main

import (
	"strconv"
	"testing"
)

func TestRequiredFuel(t *testing.T) {
	tests := []struct {
		mass int
		want int
	}{
		{12, 2},
		{14, 2},
		{1968, 654},
		{100756, 33583},
	}
	for _, tt := range tests {
		t.Run("Test for mass "+strconv.Itoa(tt.mass), func(t *testing.T) {
			if got := RequiredFuel(tt.mass); got != tt.want {
				t.Errorf("RequiredFuel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequiredActualFuel(t *testing.T) {
	tests := []struct {
		mass int
		want int
	}{
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}
	for _, tt := range tests {
		t.Run("Test for mass "+strconv.Itoa(tt.mass), func(t *testing.T) {
			if got := RequiredActualFuel(tt.mass); got != tt.want {
				t.Errorf("RequiredFuel() = %v, want %v", got, tt.want)
			}
		})
	}
}