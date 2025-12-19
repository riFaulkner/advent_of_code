package aoc

import (
	"testing"
)

func TestDay02(t *testing.T) {
	testFilePath := GetFileName(2, false)
	puzzleFilePath := GetFileName(2, true)
	type args struct {
		filepath string
		part     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Test input part 1", args: args{filepath: testFilePath, part: 1}, want: 1227775554},
		{name: "Puzzle input part 1", args: args{filepath: puzzleFilePath, part: 1}, want: 12850231731},
		{name: "Test input part 2", args: args{filepath: testFilePath, part: 2}, want: 4174379265},

		// Done with brute force, could be so much faster!
		{name: "Puzzle input part 2", args: args{filepath: puzzleFilePath, part: 2}, want: 24774350322},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day02(tt.args.filepath, tt.args.part, false); got != tt.want {
				t.Errorf("Day02() = %v, want %v", got, tt.want)
			}
		})
	}
}
