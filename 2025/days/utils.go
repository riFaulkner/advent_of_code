package aoc

import (
	"fmt"
	"strconv"
)

func GetNumberFromBytes(b []byte) int {
	str := string(b)
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return num
}

func GetNumberFromString(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return num
}

func GetFileName(day int, puzzle bool) string {
	f := "tests"
	if puzzle {
		f = "puzzles"
	}

	return fmt.Sprintf("../inputs/%s/%d.txt", f, day)
}
