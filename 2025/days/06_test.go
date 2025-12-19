package aoc

import "testing"

func TestDay06(t *testing.T) {
	testFilepath := GetFileName(6, false)
	puzzleFilepath := GetFileName(6, true)

	type args struct {
		filepath string
		part     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Part 1 Test input", args: args{filepath: testFilepath, part: 1}, want: 4277556},
		{name: "Part 1 Puzzle input", args: args{filepath: puzzleFilepath, part: 1}, want: 5316572080628},
		{name: "Part 2 Test input", args: args{filepath: testFilepath, part: 2}, want: 3263827},
		{name: "Part 2 Puzzle input", args: args{filepath: puzzleFilepath, part: 2}, want: 11299263623062},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day06(tt.args.filepath, tt.args.part); got != tt.want {
				t.Errorf("Day06() = %v, want %v", got, tt.want)
			}
		})
	}
}
