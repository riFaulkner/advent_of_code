package aoc

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"strconv"
)

func Day01(filepath string, part int, debug bool) int {
	fmt.Printf("-- Day 1 Part %d--\n", part)

	file, err := os.ReadFile(filepath)
	if err != nil {
		panic(fmt.Sprintf("Bad data, bad developer %v", err))
	}

	// Clean the data
	cleanFile := bytes.ReplaceAll(file, []byte("L"), []byte("-"))
	cleanFile = bytes.ReplaceAll(cleanFile, []byte("R"), []byte("+"))

	lines := bytes.Split(cleanFile, []byte("\n"))
	currentPosition := 50
	zeroCount := 0
	for _, v := range lines {
		turns, _ := strconv.ParseInt(string(v), 10, 32)

		if debug {
			fmt.Printf("Turning %d from position %d\n", turns, currentPosition)
		}
		zeroCount += processLine(int(turns), currentPosition, part, debug)

		currentPosition = currentPosition + int(turns)
	}

	return zeroCount
}

func processLine(turns, currentPosition, part int, debug bool) int {
	zeroCount := 0
	if part == 2 {
		fullRotations := int(math.Abs(float64(turns / 100)))
		partialRotaion := turns % 100

		absCurrentPos := currentPosition % 100

		if absCurrentPos < 0 {
			absCurrentPos = 100 + absCurrentPos // subtract the value from 100 to get the actual current position
		}

		newPosition := absCurrentPos + int(partialRotaion)
		if absCurrentPos != 0 && newPosition <= 0 || newPosition >= 100 {
			zeroCount++
		}

		zeroCount += int(fullRotations)

		if debug {
			fmt.Printf("\tVal of: %d had %d full rotations and %d partialRotations - newPosition is %d - zero count %d\n ", turns, fullRotations, partialRotaion, newPosition, zeroCount)
		}
	}
	if part == 1 {
		if currentPosition%100 == 0 {
			if debug {
				fmt.Println("Zero!")
			}
			zeroCount++
		}
	}
	return zeroCount
}
