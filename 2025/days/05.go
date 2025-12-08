package aoc

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
)

type freshnessRange struct {
	min int
	max int
}

func Day05(filepath string, part int, debug bool) int {
	// first get the list of fresh ranges
	f, _ := os.Open(filepath)
	defer f.Close()
	file := bufio.NewScanner(f)

	freshnessKeys := make([]freshnessRange, 0)
	phase := 1
	count := 0
	for file.Scan() {
		l := file.Text()

		if l == "" {
			phase++
		}

		if phase == 1 {
			freshnessKeys = append(freshnessKeys, newFreshnessRange(l))
			continue
		}

		if part == 2 {
			count := getCountOfFreshIds(freshnessKeys)
			return count
		}

		i, _ := strconv.Atoi(l)

		for _, k := range freshnessKeys {
			if k.isInRange(i) {
				count++
				break
			}
		}
	}

	return count
}
func getCountOfFreshIds(ranges []freshnessRange) int {
	//First sort the ranges by min
	slices.SortFunc(ranges, func(a, b freshnessRange) int {
		return a.min - b.min
	})

	// iterate through all the ranges, counting from the lowest known value to the high value of the ranges
	// then keep track of the lowest value checked
	lowestChecked := 0 // 0 ok?
	sum := 0
	for _, r := range ranges {
		if r.min > lowestChecked { // bring lowest check to highest of two values
			lowestChecked = r.min
		}

		// verify that we haven't jumped the entire range, if we have skip it
		if lowestChecked > r.max {
			continue
		}

		sum += r.max - (lowestChecked - 1) // account for the ranges being inclusive
		lowestChecked = r.max + 1
	}

	return sum
}

func newFreshnessRange(line string) freshnessRange {
	vs := strings.Split(line, "-")
	min, _ := strconv.Atoi(vs[0])
	max, _ := strconv.Atoi(vs[1])
	return freshnessRange{
		min: min,
		max: max,
	}
}

func (r freshnessRange) isInRange(val int) bool {
	if val >= r.min && val <= r.max {
		return true
	}
	return false
}
