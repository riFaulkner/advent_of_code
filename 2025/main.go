package main

import (
	"fmt"
	"os"
	"strconv"

	aoc "github.com/rifaulkner/aoc/2025/days"
)

const (
	puzzlePath = "inputs/puzzles/"
)

func main() {
	fmt.Println("AoC!")
	day, err := strconv.ParseInt(os.Args[1], 10, 32)
	if err != nil {
		panic(fmt.Sprintf("Couldn't get a valid day got %s", os.Args[0]))
	}
	part, err := strconv.ParseInt(os.Args[2], 10, 32)
	if err != nil {
		panic(fmt.Sprintf("Couldn't get a valid part got %s", os.Args[0]))
	}

	result := -1
	switch day {
	case 1:
		result = aoc.Day01(puzzlePath+"01-p1.txt", int(part), false)
	case 2:
		result = aoc.Day02(puzzlePath+"02-p1.txt", int(part), false)
	case 3:
		result = aoc.Day03(puzzlePath+"03-p1.txt", int(part), false)
	}

	fmt.Printf("Result for Day %d Part %d: %v\n", day, part, result)
}
