package aoc

import (
	"bytes"
	"os"
)

func Day03(filepath string, part int, debug bool) int {
	file, _ := os.ReadFile(filepath)

	size := 2
	if part == 2 {
		size = 12
	}

	sum := 0
	digits := make([]byte, size)

	for b := range bytes.SplitSeq(file, []byte("\n")) {
		if len(b) == 0 {
			break
		}

		currentStart := 0
		currentEnd := len(b) - (len(digits) - 1) // we're on the first digit to start, so add 1
		for i := range digits {
			// get larges digit from current start to end - num digits remaining
			in := getLargestDigitIndex(b[currentStart:currentEnd])
			// add to our digit list
			digits[i] = b[currentStart+in]

			// update current start
			currentStart += in + 1 // index + 1 since our search is inclusive of the first index

			// update current end
			remainingDigits := len(digits) - (i + 2)
			currentEnd = len(b) - remainingDigits
		}

		// Get the first largest value's index
		// Don't pass in the last spot, since nothing can come after it
		// i := getLargestDigitIndex(b[:(len(b) - 1)])

		// then get the next highest digit that comes after the highest digit
		// if i == len()
		// j := i + 1 + getLargestDigitIndex(b[i+1:])

		// r := make([]byte, 2)
		// r[0] = b[i]
		// r[1] = b[j]

		// sum += GetNumberFromBytes(r)
		sum += GetNumberFromBytes(digits)
	}

	return sum
}

func getLargestDigitIndex(b []byte) int {
	largest := b[0]
	index := 0

	for i := range b {
		if b[i] > largest {
			largest = b[i]
			index = i
		}
	}

	return index
}
