package main

import (
	"fmt"
	"os"
)

const BLANK = -1

func readFile() string {
	data, err := os.ReadFile("./test_input.txt")
	if err != nil {
		panic(err)
	}
	return string(data)
}

func fill(s []int, v int, times int) []int {
	for i := 0; i < times; i++ {
		s = append(s, v)
	}
	return s
}

// Because we are working with the underlying ascii values and know every character is a valid int, we can get the real int value by subtracting the input code point by the start of the range
// This is the same idea behind getting the index of the child array in a trie of the english alphabet (c - 'a')
func runeToInt(r rune) int {
	return int(r - '0')
}

// Expands dense format based on rules outlined in instructions (odd -> id, even -> .)
func expand(s string) []int {
	out := make([]int, 0)
	id := 0
	for i, c := range s {
		if i%2 == 0 {
			out = fill(out, id, runeToInt(c))
			id++
		} else {
			out = fill(out, BLANK, runeToInt(c))
		}
	}
	return out
}

func swap(s []int, i int, j int) {
	s[i], s[j] = s[j], s[i]
}

// Compresses disk by moving the left most data to the right most empty space
func compressBlockParts(disk []int) []int {
	// first blank
	l := 0
	// last value
	r := len(disk) - 1
	for l <= r {
		if disk[l] != BLANK {
			l++
			continue
		}

		if disk[r] == BLANK {
			r--
			continue
		}

		swap(disk, l, r)
	}
	return disk[:l]
}

// // inclusive
// func gapWindow(disk []int, start int) (int, int) {
// 	for start < len(disk) && disk[start] != BLANK {
// 		start++
// 	}

// 	end := start

// 	for end < len(disk) && disk[end] == BLANK {
// 		end++
// 	}

// 	return start, end - 1
// }

// // inclusive
// func fileWindow(disk []int, start int) (int, int) {
// 	for start >= 0 && disk[start] == BLANK {
// 		start--
// 	}

// 	end := start

// 	for end >= 0 && disk[end] != BLANK {
// 		end--
// 	}

// 	return end + 1, start
// }

// func compressBlocks(disk []int) {
// 	gapStart, gapEnd := 0, 0
// 	fileStart, fileEnd := len(disk)-1, len(disk)-1
// 	// remaining space -> index
// 	smallGaps := make(map[int]int)

// 	// change to "do-while"
// 	for {
// 		gapStart, gapEnd = gapWindow(disk, gapEnd)
// 		fileStart, fileEnd = fileWindow(disk, fileStart)

// 		gapSize := (gapEnd - gapStart) + 1
// 		fileSize := (fileEnd - fileStart) + 1

// 		if gapEnd > fileStart {
// 			break
// 		}

// 		// left off with swapping logic
// 		fmt.Println(gapSize)
// 		fmt.Println(fileSize)
// 		fmt.Printf("(%d, %d)\t(%d, %d)\n", gapStart, gapEnd, fileStart, fileEnd)

// 		gapEnd++
// 		fileStart--
// 	}

// }

func calculateCheckSum(disk []int) int {
	checkSum := 0
	for i, v := range disk {
		if v != BLANK {
			checkSum += i * v
		}
	}
	return checkSum
}

func main() {
	input := readFile()
	expandedDisk := expand(input)
	originalExpandedDisk := make([]int, len(expandedDisk))
	copy(originalExpandedDisk, expandedDisk)

	compressedDisk1 := compressBlockParts(expandedDisk)
	part1 := calculateCheckSum(compressedDisk1)

	// compressBlocks(originalExpandedDisk)
	part2 := 0

	fmt.Printf("Part1=%d\nPart2=%d\n", part1, part2)
}
