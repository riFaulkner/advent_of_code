package aoc

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
)

const (
	bar      = byte('|')
	splitter = byte('^')
	dot      = byte('.')
)

func Day07(filepath string, part int) int {
	// Get the grid
	file, _ := os.ReadFile(filepath)
	grid := bytes.Split(file, []byte("\n"))
	grid[0] = bytes.Replace(grid[0], []byte("S"), []byte("|"), 1)
	sum := 0

	if part == 2 {
		return d7part2(grid)
	}

	for r := range grid {
		if r == 0 {
			continue
		}
		for c := range grid[r] {
			if grid[r-1][c] == bar { // grid section above has a | bring the bar down to this level
				switch grid[r][c] {
				case bar:
					continue
				case splitter:
					if right := findNextBeamSpot(grid, c, r, 1); right != -1 {
						grid[r][right] = bar
					}
					if left := findNextBeamSpot(grid, c, r, -1); left != -1 {
						grid[r][left] = bar
					}
					sum++
				default:
					// make the current location a | and move on
					grid[r][c] = bar
				}
			}
		}
	}

	return sum
}

// find the column value of the next available location to place your split
func findNextBeamSpot(grid [][]byte, c, r, direction int) int {
	for i := c + direction; c <= 0 || c < len(grid[r]); i += direction {
		switch grid[r][i] {
		case bar:
			return -1
		case dot:
			return i
		}
	}
	return -1
}

func d7part2(grid [][]byte) int {
	// sum := 1
	re := regexp.MustCompile(`\|`)
	startC := re.FindIndex(grid[0])

	cache := make(map[string]int)
	sum := findRealities(grid[:len(grid)-1], cache, startC[0], 0)

	return sum
}

func findRealities(grid [][]byte, cache map[string]int, locc, locr int) int {
	for r := locr + 1; r < len(grid); r++ { // can we actually look at row + 1?
		if locc < 0 || locc >= len(grid[r]) {
			return 0
		}

		switch grid[r][locc] {
		case dot:
			// we're good to keep c in the same spot, move one row down
			continue
		case splitter:
			cacheKey := fmt.Sprintf("%d-%d", locc, r)
			if v, ok := cache[cacheKey]; ok { // we've been here before, return the value quickly /* instead */
				return v
			}
			// its a splitter, if possible move to the left -- and add a new value to the stack of proceess
			l := findRealities(grid, cache, locc-1, r) + findRealities(grid, cache, locc+1, r)

			cache[cacheKey] = l
			return l
		}
	}

	return 1
}
