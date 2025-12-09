package aoc

import (
	"bufio"
	"bytes"
	"os"
	"regexp"
)

func Day06(filepath string, part int) int {
	if part == 2 { // Its literally this different
		return part2(filepath)
	}

	file, _ := os.Open(filepath)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// For each line, scan it, if its numbers load it into the slice of slices. if its the final row, perform the arithmetic
	table := make([][]int, 0)
	for scanner.Scan() {
		l := scanner.Text()

		re := regexp.MustCompile(`\d+`)

		cols := re.FindAllString(l, -1)

		if len(cols) == 0 { // we're at the line of symbols
			// for each symbol, get what its operation is, and then for each line we have, pass in the symbol go row by row

			re = regexp.MustCompile(`[+*]`)

			cols = re.FindAllString(l, -1)
			sum := 0
			for i, c := range cols {
				columnSum := 0
				for j := range table {
					if j == 0 {
						columnSum = table[j][i]
						continue
					}
					columnSum = performOperation(columnSum, table[j][i], c)
				}
				sum += columnSum
			}
			return sum
		}

		// check for numbers b4 this
		line := make([]int, len(cols))

		for i := range len(cols) {
			line[i] = GetNumberFromString(cols[i])
		}

		table = append(table, line)
	}

	return -1
}

func performOperation(sum int, newValue int, opperation string) int {
	switch opperation {
	case "*":
		return sum * newValue
		// multiply
	case "+":
		// add
		return sum + newValue
	}

	panic("AHHHH")
}

func part2(filepath string) int {
	file, _ := os.ReadFile(filepath)
	// First get a slice of bytes this should make the file a grid
	grid := bytes.Split(file, []byte("\n"))
	re := regexp.MustCompile(`\d+|[+*]`)

	proDigits := make([]int, 0)
	probOpperator := ""
	sum := 0

	for c := range grid[0] { // Assume that the grid is a prefect square
		v := columnAsRow(c, grid)
		vals := re.FindAll(v, -1)

		if len(vals) == 0 {
			// set up the new problem
			colSum := proDigits[0]
			for n := range proDigits {
				if n == 0 {
					continue
				}
				colSum = performOperation(colSum, proDigits[n], probOpperator)
			}
			sum += colSum

			// reset
			probOpperator = ""
			proDigits = make([]int, 0)
			continue
		}
		if len(vals) == 2 {
			probOpperator = string(vals[1])
		}

		proDigits = append(proDigits, GetNumberFromBytes(vals[0]))
	}

	colSum := proDigits[0]
	for n := range proDigits {
		if n == 0 {
			continue
		}
		colSum = performOperation(colSum, proDigits[n], probOpperator)
	}
	sum += colSum

	return sum
}

func columnAsRow(col int, grid [][]byte) []byte {
	res := make([]byte, len(grid)-1)
	for i := range grid {
		if len(grid[i]) < col+1 {
			break
		}
		res[i] = grid[i][col]
	}

	return res
}
