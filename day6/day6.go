package main

import(
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func get_input() []int {
	dat, err := ioutil.ReadFile("day6/day6.txt")
	check(err)
	input := string(dat)[:(len(string(dat))-1)]
	as_strings := strings.Split(input, "\t")
	blocks := []int{}
	for _, i := range(as_strings) {
		num, _ := strconv.Atoi(i)
		blocks = append(blocks, num)
	}
	return blocks
}

func max_left_index(numbers []int) int {
	// Returns the index of the max item in the array (ties go to the lowest tied index)
	max := 0
	max_index := 0
	for index, i := range(numbers){
		if i > max{
			max = i
			max_index = index
		}
	}
	return max_index
}

func distribute_blocks(blocks []int) []int {
	max_index := max_left_index(blocks)
	to_disperse := blocks[max_index]
	blocks[max_index] = 0
	location := (max_index + 1) % len(blocks)
	for i := 0; i < to_disperse; i++ {
		blocks[location] += 1
		location += 1
		location = location % len(blocks)
	} 
	return blocks
}

func slices_equal(a []int, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func contains(a [][]int, b []int) bool {
	for _, i := range(a){
		if slices_equal(b, i){
			return true
		}
	}
	return false
}

func part_1(blocks []int) int {
	blocks_copy := make([]int, len(blocks))
    copy(blocks_copy, blocks)
	seen := [][]int{blocks_copy}
	steps := 0
	for {
		blocks = distribute_blocks(blocks)
		steps += 1
		if contains(seen, blocks){
			return steps
		}
		blocks_copy := make([]int, len(blocks))
		copy(blocks_copy, blocks)
		seen = append(seen, blocks_copy)
	}
}

func part_2(blocks []int) int {
	// get into a known repeating state:
	blocks_copy := make([]int, len(blocks))
    copy(blocks_copy, blocks)
	seen := [][]int{blocks_copy}
	steps := 0
	for {
		blocks = distribute_blocks(blocks)
		steps += 1
		if contains(seen, blocks){
			break
		}
		blocks_copy := make([]int, len(blocks))
		copy(blocks_copy, blocks)
		seen = append(seen, blocks_copy)
	}

	steps = 0
	target := make([]int, len(blocks))
	copy(target, blocks)
	for {
		blocks = distribute_blocks(blocks)
		steps += 1
		if slices_equal(target, blocks){
			break
		}
	}
	return steps
}

func main() {
	fmt.Printf("Part 1: %d\n", part_1(get_input()))
	fmt.Printf("Part 2: %d\n", part_2(get_input()))
}