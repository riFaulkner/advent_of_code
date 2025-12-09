package aoc

import "strconv"

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
