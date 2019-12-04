package main

import (
	"reflect"
	"strconv"
	"testing"
)

func TestIntersections(t *testing.T) {
	type args struct {
		wire1 []string
		wire2 []string
	}
	tests := []struct {
		args args
		want []Position
	}{
		{args: args{
			wire1: []string{"R8", "U5", "L5", "D3"},
			wire2: []string{"U7", "R6", "D4", "L4"},
		}, want: []Position{{6, 5}, {3, 3}}},
	}
	for i, tt := range tests {
		t.Run("Test "+strconv.Itoa(i), func(t *testing.T) {
			if got := Intersections(tt.args.wire1, tt.args.wire2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersections() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClosetManhattanDistance(t *testing.T) {
	type args struct {
		wire1 []string
		wire2 []string
	}
	tests := []struct {
		args args
		want int
	}{
		{args: args{
			wire1: []string{"R8", "U5", "L5", "D3"},
			wire2: []string{"U7", "R6", "D4", "L4"},
		}, want: 6},
	}
	for i, tt := range tests {
		t.Run("Test "+strconv.Itoa(i), func(t *testing.T) {
			if got := ClosetManhattanDistance(Intersections(tt.args.wire1, tt.args.wire2)); got != tt.want {
				t.Errorf("ClosetManhattanDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
