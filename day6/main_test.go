package main

import "testing"

func TestHops(t *testing.T) {
	tests := []struct {
		name  string
		lines []string
		want  int
	}{
		{"Example", []string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L", "K)YOU", "I)SAN"}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hops(tt.lines); got != tt.want {
				t.Errorf("Hops() = %v, want %v", got, tt.want)
			}
		})
	}
}
