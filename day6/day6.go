package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"

	"github.com/evanmpollack/day6/internal/set"
	"github.com/evanmpollack/day6/internal/state"
)

type position struct {
	row int
	col int
}

func readFile() []string {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	maze := make([]string, 0)
	for scanner.Scan() {
		maze = append(maze, scanner.Text())
	}
	return maze
}

func getStartPosition(grid []string) position {
	row := 0
	col := 0
	for r, s := range grid {
		for co, ch := range s {
			if ch == '^' {
				row = r
				col = co
				break
			}
		}
	}
	return position{row: row, col: col}
}

// a struct's internals, if primitive, are able to be compared using ==. Therefore, we do not want pointers, we want the actual struct in the set. Otherwise, we would have to implement and impose an Equals method for the generic set
// pos is the underlying struct because we don't want the values to be altered between frames (e.g. if we have to backtrack, even when switching directions, we will be on top of a #, breaking the algortthm)
func countStepsInRoute(maze []string, pos position, count int, dir state.Direction, visited *set.Set[position]) (bool, int) {
	if !inBounds(maze, pos) {
		return true, count
	}

	if maze[pos.row][pos.col] == '#' {
		return false, count
	}

	if !set.Contains(visited, pos) {
		count++
		set.Add(visited, pos)
	}

	switch dir {
	case state.Up:
		pos.row--
	case state.Right:
		pos.col++
	case state.Down:
		pos.row++
	case state.Left:
		pos.col--
	}

	exit, count := countStepsInRoute(maze, pos, count, dir, visited)

	if !exit {
		exit, count = countStepsInRoute(maze, pos, count, state.Next(dir), visited)
	}

	return exit, count
}

func inBounds(grid []string, pos position) bool {
	return pos.row >= 0 && pos.row < len(grid) &&
		pos.col >= 0 && pos.col < utf8.RuneCountInString(grid[pos.row])
}

func main() {
	maze := readFile()
	startingPoint := getStartPosition(maze)
	_, part1 := countStepsInRoute(maze, startingPoint, 0, state.First(), set.NewSet[position]())

	part2 := 0

	/**

	Plan:
	1. Run variation of dfs on maze
	2. Increment a counter for each step (after pre check)
	3. return bool to ensure we can swtich directions if we reach an obstacle

	**/

	fmt.Printf("Part1=%d\nPart2=%d\n", part1, part2)
}
