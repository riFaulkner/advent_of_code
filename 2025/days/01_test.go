package aoc

import (
	"fmt"
	"testing"
)

const (
	testFilePath   = "../inputs/tests/01-p1.txt"
	puzzleFilePath = "../inputs/puzzles/01-p1.txt"
)

func TestDay01(t *testing.T) {
	type args struct {
		filepath string
		part     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Happy Path Part 1", args: args{filepath: testFilePath, part: 1}, want: 3},
		{name: "Happy Path Part 2", args: args{filepath: testFilePath, part: 2}, want: 6},
		{name: "Puzzle 1", args: args{filepath: puzzleFilePath, part: 1}, want: 1118},
		{name: "Puzzle 2", args: args{filepath: puzzleFilePath, part: 2}, want: 6289},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day01(tt.args.filepath, tt.args.part, true); got != tt.want {
				t.Errorf("Day01() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestDay01Cases(t *testing.T) {
	type args struct {
		turns           int
		currentPosition int
		part            int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Turn to the right", args: args{turns: 25, currentPosition: 50, part: 2}, want: 0},
		{name: "Turn to the left", args: args{turns: -25, currentPosition: 50, part: 2}, want: 0},
		{name: "Turn way to the left", args: args{turns: -125, currentPosition: 50, part: 2}, want: 1},
		{name: "Turn way, way to the left", args: args{turns: -175, currentPosition: 50, part: 2}, want: 2},
		{name: "Turn little to the left", args: args{turns: -75, currentPosition: 50, part: 2}, want: 1},
		{name: "Turn littler to the left", args: args{turns: -50, currentPosition: 50, part: 2}, want: 1},
		{name: "Turn littler to the right", args: args{turns: 50, currentPosition: 50, part: 2}, want: 1},
	}
	fmt.Println("")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processLine(tt.args.turns, tt.args.currentPosition, tt.args.part, false); got != tt.want {
				t.Errorf("processLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
