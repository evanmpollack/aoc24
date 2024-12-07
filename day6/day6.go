package main

import (
	"os"
	"bufio"
	"fmt"
)

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

func main() {
	maze := readFile()
	fmt.Println(maze)

	/**

	Plan:
	1. Run dfs on maze
	2. Increment a counter for each step (after pre check)
	3. return bool to ensure we can break out early if the guard exits the maze

	**/
}