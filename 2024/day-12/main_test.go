package main

import (
	"os"
	"testing"
)

func Test_part1(t *testing.T) {
	input, _ := os.ReadFile("./input_test")
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"actual", string(input), 140},
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
		{"actual", string(input), 80},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_part2_1(t *testing.T) {
	input, _ := os.ReadFile("./input_test")
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"actual", string(input), 80},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
