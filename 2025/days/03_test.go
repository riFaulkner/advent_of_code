package aoc

import "testing"

const (
	d3TestInputPath   = "../inputs/tests/03-p1.txt"
	d3PuzzleInputPath = "../inputs/puzzles/03-p1.txt"
)

func TestDay03(t *testing.T) {
	type args struct {
		filepath string
		part     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Part 1 Test", args: args{filepath: d3TestInputPath, part: 1}, want: 357},
		{name: "Part 1 Puzzle", args: args{filepath: d3PuzzleInputPath, part: 1}, want: 17278},
		{name: "Part 2 Test", args: args{filepath: d3TestInputPath, part: 2}, want: 3121910778619},
		{name: "Part 2 Puzzle", args: args{filepath: d3PuzzleInputPath, part: 2}, want: 171528556468625},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day03(tt.args.filepath, tt.args.part, true); got != tt.want {
				t.Errorf("Day03() = %v, want %v", got, tt.want)
			}
		})
	}
}
