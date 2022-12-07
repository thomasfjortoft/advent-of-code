package main

import (
	"os"
	"testing"
)

func Test_part1(t *testing.T) {
	input, _ := os.ReadFile("./input_test")
	input2, _ := os.ReadFile("./input")
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"actual", string(input), 95437},
		{"actual", string(input2), 2031851},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	input, _ := os.ReadFile("./input_test")
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"actual", string(input), 24933642},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
