package aoc

import "testing"

const (
	d4TestInputPath   = "../inputs/tests/04.txt"
	d4PuzzleInputPath = "../inputs/puzzles/04.txt"
)

func TestDay04(t *testing.T) {
	type args struct {
		filepath string
		part     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Part 1 Test Input", args: args{filepath: d4TestInputPath, part: 1}, want: 13},
		{name: "Part 1 Puzzle Input", args: args{filepath: d4PuzzleInputPath, part: 1}, want: 1457},
		{name: "Part 2 Test Input", args: args{filepath: d4TestInputPath, part: 2}, want: 43},
		{name: "Part 2 Puzzle Input", args: args{filepath: d4PuzzleInputPath, part: 2}, want: 8310},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day04(tt.args.filepath, tt.args.part); got != tt.want {
				t.Errorf("Day04() = %v, want %v", got, tt.want)
			}
		})
	}
}
