package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"

	"github.com/evanmpollack/day4/internal"
)

var XMAS = []rune{'X', 'M', 'A', 'S'}
var MAS = XMAS[1:]

func readFile() []string {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func inBounds(grid []string, row int, col int) bool {
	return row < len(grid) && row >= 0 && col < utf8.RuneCountInString(grid[row]) && col >= 0
}

func searchDirectionsXMAS(grid []string, row int, col int) int {
	// 1. Allocate slices for directions

	up := make([]rune, 0)
	upRight := make([]rune, 0)
	right := make([]rune, 0)
	downRight := make([]rune, 0)
	down := make([]rune, 0)
	downLeft := make([]rune, 0)
	left := make([]rune, 0)
	upLeft := make([]rune, 0)

	// 2. Collect characters in each direction

	for i := range XMAS {
		// is conversion to rune necessary or can I just use bytes in trie?
		if inBounds(grid, row-i, col) {
			up = append(up, rune(grid[row-i][col]))
		}

		if inBounds(grid, row-i, col+i) {
			upRight = append(upRight, rune(grid[row-i][col+i]))
		}

		if inBounds(grid, row, col+i) {
			right = append(right, rune(grid[row][col+i]))
		}

		if inBounds(grid, row+i, col+i) {
			downRight = append(downRight, rune(grid[row+i][col+i]))
		}

		if inBounds(grid, row+i, col) {
			down = append(down, rune(grid[row+i][col]))
		}

		if inBounds(grid, row+i, col-i) {
			downLeft = append(downLeft, rune(grid[row+i][col-i]))
		}

		if inBounds(grid, row, col-i) {
			left = append(left, rune(grid[row][col-i]))
		}

		if inBounds(grid, row-i, col-i) {
			upLeft = append(upLeft, rune(grid[row-i][col-i]))
		}
	}

	// 3. Build trie from directions -- figure out how to properly import local packages

	trie := internal.NewTrie()
	internal.Insert(trie, up)
	internal.Insert(trie, upRight)
	internal.Insert(trie, right)
	internal.Insert(trie, downRight)
	internal.Insert(trie, down)
	internal.Insert(trie, downLeft)
	internal.Insert(trie, left)
	internal.Insert(trie, upLeft)

	// 4. Count the times the S in XMAS is "inserted" when building trie, return 0 otherwise

	return internal.CountEntries(trie, XMAS)
}

func searchDirectionsX_MAS(grid []string, row int, col int) int {
	valid := inBounds(grid, row - 1, col - 1) &&
		inBounds(grid, row - 1, col + 1) &&
		inBounds(grid, row + 1, col + 1) &&
		inBounds(grid, row + 1, col - 1)
	
	if !valid {
		return 0
	}

	count := 0


	if ((grid[row - 1][col - 1] == 'M' && grid[row + 1][col + 1] == 'S') || (grid[row - 1][col - 1] == 'S' && grid[row + 1][col + 1] == 'M')) &&
		((grid[row - 1][col + 1] == 'M' && grid[row + 1][col - 1] == 'S') || (grid[row - 1][col + 1] == 'S' && grid[row + 1][col - 1] == 'M')) {
		count++
	}
	
	return count
}

func main() {
	lines := readFile()

	// off by 3 in test input, something is being counted more than necessary
	part1 := func() int {
		totalXMAS := 0
		for row, s := range lines {
			for col, c := range s {
				if c == 'X' {
					totalXMAS += searchDirectionsXMAS(lines, row, col)
				}
			}
		}
		return totalXMAS
	}

	part2 := func() int {
		totalX_MAS := 0
		for row, s := range lines {
			for col, c := range s {
				if c == 'A' {
					totalX_MAS += searchDirectionsX_MAS(lines, row, col)
				}
			}
		}
		return totalX_MAS
	}

	fmt.Printf("Part1=%d\nPart2=%d\n", part1(), part2())
}
