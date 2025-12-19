package aoc

import "testing"

func TestDay03(t *testing.T) {
	testFilepath := GetFileName(3, false)
	puzzleFilepath := GetFileName(3, true)
	type args struct {
		filepath string
		part     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Part 1 Test", args: args{filepath: testFilepath, part: 1}, want: 357},
		{name: "Part 1 Puzzle", args: args{filepath: puzzleFilepath, part: 1}, want: 17278},
		{name: "Part 2 Test", args: args{filepath: testFilepath, part: 2}, want: 3121910778619},
		{name: "Part 2 Puzzle", args: args{filepath: puzzleFilepath, part: 2}, want: 171528556468625},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day03(tt.args.filepath, tt.args.part, true); got != tt.want {
				t.Errorf("Day03() = %v, want %v", got, tt.want)
			}
		})
	}
}
