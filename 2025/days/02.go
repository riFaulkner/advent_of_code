package aoc

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func Day02(filepath string, part int, debug bool) int {
	fmt.Printf("-- Day 2 Part %d--\n", part)

	file, _ := os.ReadFile(filepath)
	sum := 0

	for v := range bytes.SplitSeq(file, []byte(",")) {
		if debug {
			fmt.Printf("bytes: %v: \n", string(v))
		}
		r := regexp.MustCompile(`\d+`)

		vals := r.FindAll(v, 2)

		start := GetNumberFromBytes(vals[0])
		end := GetNumberFromBytes(vals[1])

		for i := start; i < end+1; i++ {
			iStr := strconv.FormatInt(int64(i), 10)

			if part == 1 {
				if findRepeating([]byte(iStr)) {
					sum += i
				}
			} else {
				if findRepeatingHard([]byte(iStr)) {
					sum += i
				}
			}
		}
		if debug {

			fmt.Printf("\t\t Sum: %v\n", sum)
		}
	}

	return sum
}

func findRepeatingHard(b []byte) bool {
	for i := 1; i <= len(b)/2; i++ {
		size := len(b) / i

		if size == len(b) {
			size = 1
		}
		if len(b)%i != 0 { // If this number of chars doesn't exactly fit the number of total characters in the ID, then it can't be split this way. e.g. can't have a 4 char sequence for a 9 char number
			continue
		}

		// check through each of the splits to see if all of the splits match
		subset := string(b[:i])

		re := fmt.Sprintf("%s", subset)
		r := regexp.MustCompile(re)
		count := len(r.FindAll(b, -1))
		if count*len(subset) == len(b) {
			// match found
			return true
		}
	}

	return false
}

func findRepeating(b []byte) bool {
	if len(b)%2 == 1 {
		return false
	}
	half := len(b) / 2

	return slices.Equal(b[:half], b[half:])
}
