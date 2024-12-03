package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readFile() [][]string {
	// alternative to var file *File, err error = os.Open()
	file, err := os.Open("./input.txt")
	if err != nil {
		// stop execution -- not best practice
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// make is preferred to var pairs [][]string because it can be treated differently
	pairs := make([][]string, 0)
	// while scanner.hasNext() equivalent
	for scanner.Scan() {
		line := scanner.Text()
		pair := strings.Split(line, "   ")
		pairs = append(pairs, pair)
	}
	return pairs
}

func splitCols(pairs [][]string) ([]int, []int) {
	var col1, col2 []int
	for i := 0; i < len(pairs); i++ {
		val1, err1 := strconv.Atoi(pairs[i][0])
		if err1 != nil {
			panic(err1)
		}
		val2, err2 := strconv.Atoi(pairs[i][1])
		if err2 != nil {
			panic(err2)
		}
		col1 = append(col1, val1)
		col2 = append(col2, val2)
	}
	return col1, col2
}

func main() {
	pairs := readFile()
	r, l := splitCols(pairs)
	// part1/part2 are lambdas for visual separation in main, side effects from sorting doesn't matter for part 2

	// Time Complexity: O(2(nlogn)) => O(nlogn)
	// Space Complexity: O(1), assumes that r and l would be given as inputs if we didn't have to read from file
	part1 := func() int {
		slices.Sort(r)
		slices.Sort(l)
		sum := 0
		for i := 0; i < len(pairs); i++ {
			// explicit conversions are required in go
			sum += int(math.Abs(float64(r[i] - l[i])))
		}
		return sum
	}

	// Time Complexity: O(2n) => O(n)
	// Space Complexity: O(n), assumes that r and l would be given as inputs if we didn't have to read from file
	// If l and r are sorted, I wonder if a dynamic sliding window on the right list would allow us to negate the need for a map
	part2 := func() int {
		rCount := make(map[int]int)
		for i := 0; i < len(r); i++ {
			rCount[r[i]] = rCount[r[i]] + 1
		}

		sum := 0
		for i := 0; i < len(l); i++ {
			sum += l[i] * rCount[l[i]]
		}
		return sum
	}

	fmt.Printf("Part1=%d\nPart2=%d\n", part1(), part2())
}
