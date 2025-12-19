package aoc

import "testing"

func TestDay07(t *testing.T) {
	testFilepath := GetFileName(7, false)
	puzzleFilepath := GetFileName(7, true)

	type args struct {
		filepath string
		part     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Test Input Part 1", args: args{filepath: testFilepath, part: 1}, want: 21},
		{name: "Puzzle Input Part 1", args: args{filepath: puzzleFilepath, part: 1}, want: 1590},
		{name: "Test Input Part 2", args: args{filepath: testFilepath, part: 2}, want: 40},
		{name: "Puzzle Input Part 2", args: args{filepath: puzzleFilepath, part: 2}, want: 20571740188555},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day07(tt.args.filepath, tt.args.part); got != tt.want {
				t.Errorf("Day07() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkDay07(b *testing.B) {
	for b.Loop() {
		Day07(GetFileName(7, true), 1)
	}
}
