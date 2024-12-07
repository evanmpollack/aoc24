package main

import (
	"os"
	"bufio"
	"strings"
	"strconv"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile() (map[int][]int, [][]int) {
	file, err := os.Open("./input.txt")
	checkErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	adj := make(map[int][]int)
	for scanner.Scan() && scanner.Text() != "" {
		mapping := strings.Split(scanner.Text(), "|")
		key, errK := strconv.Atoi(mapping[0])
		checkErr(errK)
		value, errV := strconv.Atoi(mapping[1])
		checkErr(errV)

		if adj[key] == nil {
			adj[key] = make([]int, 0)
		}

		adj[key] = append(adj[key], value)
	}

	lines := make([][]int, 0)
	for scanner.Scan() {
		line := make([]int, 0)
		for _, s := range strings.Split(scanner.Text(), ",") {
			i, err := strconv.Atoi(s)
			checkErr(err)
			line = append(line, i)
		}
		lines = append(lines, line)
	}
	
	return adj, lines
}

func topologicalSort(graph map[int][]int, curr int, out []int) []int {
	return out
}

func main() {
	adjList, lines := readFile()

	/**

	Plan:
	1. Create dependency graph
	2. Topological sort
	3. Back sorted list with value -> index lookups
	4. See if I can find a way to know whether or not the value at the input index is in the proper order

	**/
}