package aoc

import "testing"

const (
	d6TestInputPath   = "../inputs/tests/06-p1.txt"
	d6PuzzleInputPath = "../inputs/puzzles/06.txt"
)

func TestDay06(t *testing.T) {
	type args struct {
		filepath string
		part     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Part 1 Test input", args: args{filepath: d6TestInputPath, part: 1}, want: 4277556},
		{name: "Part 1 Puzzle input", args: args{filepath: d6PuzzleInputPath, part: 1}, want: 5316572080628},
		{name: "Part 2 Test input", args: args{filepath: d6TestInputPath, part: 2}, want: 3263827},
		{name: "Part 2 Puzzle input", args: args{filepath: d6PuzzleInputPath, part: 2}, want: 5316572080628},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day06(tt.args.filepath, tt.args.part); got != tt.want {
				t.Errorf("Day06() = %v, want %v", got, tt.want)
			}
		})
	}
}
