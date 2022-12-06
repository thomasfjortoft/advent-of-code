package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"one", "bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"two", "nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"three", "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"four", "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
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
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"one", "bvwbjplbgvbhsrlpgdmjqwftvncz", 23},
		{"two", "nppdvjthqldpwncqszvftbrmjlhg", 23},
		{"three", "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29},
		{"four", "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
