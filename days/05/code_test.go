package _5

import "testing"

func TestGetClosestSeedPlaningLocation(t *testing.T) {
	type args struct {
		inputFileName string
		seedsAsRange  bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		{"Example", args{"example.txt", false}, 35},
		{"Puzzle", args{"puzzle_work.txt", false}, 309796150},
		{"Example seeds are a range", args{"example.txt", true}, 46},
		{"Puzzle 02", args{"puzzle_work.txt", true}, 50716416},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetClosestSeedPlaningLocation(tt.args.inputFileName, tt.args.seedsAsRange); got != tt.want {
				t.Errorf("GetClosestSeedPlaningLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}