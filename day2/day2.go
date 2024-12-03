package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func readFile() [][]int {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	reports := make([][]int, 0)
	for scanner.Scan() {
		levels := make([]int, 0)
		strLevels := strings.Split(scanner.Text(), " ")
		for i := 0; i < len(strLevels); i++ {
			num, err := strconv.Atoi(strLevels[i])
			if err != nil {
				panic(err)
			}
			levels = append(levels, num)
		}
		reports = append(reports, levels)
	}
	return reports
}

func isDecreasing(levels []int) bool {
	decreasing := true
	for i := 1; i < len(levels); i++ {
		if levels[i-1] > levels[i] {
			decreasing = false
			break
		}
	}
	return decreasing
}

func isIncreasing(levels []int) bool {
	increasing := true
	for i := 1; i < len(levels); i++ {
		if levels[i-1] < levels[i] {
			increasing = false
			break
		}
	}
	return increasing
}

func isSafeDifference(levels []int) bool {
	safe := true
	for i := 1; i < len(levels); i++ {
		absDiff := int(math.Abs(float64(levels[i] - levels[i-1])))
		if absDiff < 1 || absDiff > 3 {
			safe = false
			break
		}
	}
	return safe
}

func main() {
	reports := readFile()

	// Time Complexity O(r * l), where r is the number of reports and l is the number of levels in a report
	// Space Complextity: O(1), no extra space used
	part1 := func() int {
		safe := 0

		for i := 0; i < len(reports); i++ {
			levels := reports[i]
			if (isDecreasing(levels) || isIncreasing(levels)) && isSafeDifference(levels) {
				safe++
			}
		}

		return safe
	}

	part2 := func() int {
		return 0
	}

	fmt.Printf("Part1=%d\nPart2=%d\n", part1(), part2())
}
