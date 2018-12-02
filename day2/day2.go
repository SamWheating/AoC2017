package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		fmt.Print(e)
		panic(e)
	}
}

func max(v []int) int {
	max := 0
	for i := 0; i < len(v); i++ {
		if v[i] > max {
			max = v[i]
		}
	}
	return max
}

func min(v []int) int {
	min := 100000000
	for i := 0; i < len(v); i++ {
		if v[i] < min {
			min = v[i]
		}
	}
	return min
}

func main() {

	// Parsing input file into a 2D array of integers
	dat, err := ioutil.ReadFile("day2/day2.txt")
	check(err)
	input := string(dat)[:(len(string(dat)) - 1)]
	lines := strings.Split(input, "\n")
	int_lists := [][]int{}
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		line = strings.Replace(line, "\n", "", 1)
		var ints = []int{}
		numbers := strings.Split(line, "\t")

		for _, i := range numbers {
			j, err := strconv.Atoi(i)
			check(err)
			ints = append(ints, j)
		}
		int_lists = append(int_lists, ints)
	}

	// Part 1 - Get the sum of (max - min) for every row:
	checksum := 0
	for i := 0; i < len(int_lists); i++ {
		checksum += (max(int_lists[i]) - min(int_lists[i]))
	}

	fmt.Printf("Part 1: %d\n", checksum)

	// Part 2 - The sum of the only clean division which can be made in each row
	checksum = 0
	found := false
	for i := 0; i < len(int_lists); i++ {
		found = false
		for j := 0; j < len(int_lists[i]); j++ {
			for k := 0; k < len(int_lists[i]); k++ {
				if (int_lists[i][j]%int_lists[i][k]) == 0 && !found && j != k {
					checksum += int_lists[i][j] / int_lists[i][k]
					found = true
				}
			}
		}
	}

	fmt.Printf("Part 2: %d\n", checksum)

}
