package aoc

import "testing"

const (
	d5TestInputPath   = "../inputs/tests/05-p1.txt"
	d5PuzzleInputPath = "../inputs/puzzles/05-p1.txt"
)

func TestDay05(t *testing.T) {
	type args struct {
		filepath string
		part     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Part 1 Test Input", args: args{filepath: d5TestInputPath, part: 1}, want: 3},
		{name: "Part 1 Puzzle Input", args: args{filepath: d5PuzzleInputPath, part: 1}, want: 737},
		{name: "Part 2 Test Input", args: args{filepath: d5TestInputPath, part: 2}, want: 14},
		{name: "Part 2 Puzzle Input", args: args{filepath: d5PuzzleInputPath, part: 2}, want: 357485433193284}, // originally missed it by 8, correction got it right tho!
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day05(tt.args.filepath, tt.args.part, true); got != tt.want {
				t.Errorf("Day05() = %v, want %v", got, tt.want)
			}
		})
	}
}
