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
	dat, err := ioutil.ReadFile("day5/day5.txt")
	check(err)
	input := string(dat)[:(len(string(dat))-1)]
	as_strings := strings.Split(input, "\n")
	jumps := []int{}
	for _, i := range(as_strings) {
		num, _ := strconv.Atoi(i)
		jumps = append(jumps, num)
	}
	return jumps
}

func part_1(jumps []int) int {
	position := 0
	prev_position := 0
	jumps_taken := 0
	for {
		if position < 0 || position >= len(jumps){
			return jumps_taken
		}
		prev_position = position
		position += jumps[position]
		jumps[prev_position] += 1
		jumps_taken += 1
	}
}

func part_2(jumps []int) int {

	position := 0
	prev_position := 0
	jumps_taken := 0
	for {
		if position < 0 || position >= len(jumps){
			return jumps_taken
		}
		prev_position = position
		position += jumps[position]
		if jumps[prev_position] >= 3{
			jumps[prev_position] -= 1
		} else {
			jumps[prev_position] += 1
		}
		jumps_taken += 1
	}
}

func main(){
	fmt.Printf("Part 1: %d\n", part_1(get_input()))
	fmt.Printf("Part 2: %d\n", part_2(get_input()))
}