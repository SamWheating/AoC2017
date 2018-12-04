package main

import (
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		fmt.Print(e)
		panic(e)
	}
}

func get_input() string {
	dat, err := ioutil.ReadFile("day9/day9.txt")
	check(err)
	input := string(dat)[:(len(string(dat)) - 1)]
	return input
}

func cancel_characters(stream string) string {
	// removes characters following ! from the string
	i := 0
	cleaned_string := ""
	for {
		if string(stream[i]) == "!"{
			i += 2
		} else {
			cleaned_string += string(stream[i])
			i +=1 
		}
		if i == len(stream) {
			return cleaned_string
		}
	}
}

func remove_garbage(stream string) string {
	// removes everything between sets of <> from the string
	i := 0
	in_garbage := false
	garbage_free := ""
	for {
		if string(stream[i]) == "<" && !in_garbage{
			in_garbage = true
		} else if string(stream[i]) == ">" && in_garbage {
			in_garbage = false
		} else if !in_garbage {
			garbage_free += string(stream[i])
		}
		i += 1
		if i == len(stream){
			return garbage_free
		}
	}
}

func total_points(stream string) int {
	// sum the value of the trash (part 1)
	score := 0
	current_depth := 1
	for _, i := range(stream){
		if string(i) == "{" {
			score += current_depth
			current_depth += 1
		} else if string(i) == "}" {
			current_depth -= 1
		}
	}
	return score
}

func count_garbage (stream string) int {
	// removes everything between sets of <> from the string (part 2)
	i := 0
	in_garbage := false
	garbage_count := 0
	for {
		if string(stream[i]) == "<" && !in_garbage{
			in_garbage = true
		} else if string(stream[i]) == ">" && in_garbage {
			in_garbage = false
		} else if in_garbage {
			garbage_count += 1
		}
		i += 1
		if i == len(stream){
			return garbage_count
		}
	}
}

func main() {
	//fmt.Print(get_input())
	//fmt.Println(remove_garbage(cancel_characters(get_input())))
	fmt.Printf("Part 1: %d\n", total_points(remove_garbage(cancel_characters(get_input()))))
	fmt.Printf("Part 2: %d\n", count_garbage(cancel_characters(get_input())))
}