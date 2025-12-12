package aoc

import (
	"bytes"
	"os"
)

func Day04(filepath string, part int) int {
	file, _ := os.ReadFile(filepath)
	grid := bytes.Split(file, []byte("\n"))
	roll := byte('@')
	dot := byte('.')
	sum := 0
	startSum := -1

	for startSum < sum {
		startSum = sum

		for r := range grid {
			for c := range grid[r] {
				if grid[r][c] == roll {
					if checkSurrounding(grid, r, c) {
						sum++
						if part == 2 {
							grid[r][c] = dot
						}
					}
				}
			}
		}
		if part == 1 {
			startSum = sum + 1
		}
	}

	return sum
}

func checkSurrounding(grid [][]byte, row, col int) bool {
	sum := 0
	for i := row - 1; i < row+2; i++ {
		if i < 0 || i >= len(grid) { // out of bounds
			continue
		}
		for j := col - 1; j < col+2; j++ {
			if j == col && row == i {
				continue // that our current value
			}
			if j < 0 || j >= len(grid[i]) { // out of bounds!
				continue
			}

			if grid[i][j] == []byte("@")[0] {
				sum++
			}
		}
	}

	return sum <= 3
}
