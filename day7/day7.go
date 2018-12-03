package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type program struct {
	name     string
	weight   int
	children []string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func get_input() []program {
	dat, err := ioutil.ReadFile("day7/day7.txt")
	check(err)
	lines := strings.Split(string(dat)[:len(string(dat))-1], "\n")
	programs := []program{}

	for _, line := range lines {
		children := []string{}
		line = strings.Replace(line, "\n", "", 1)
		name := strings.Split(line, " ")[0]
		weight, _ := strconv.Atoi(strings.Split(strings.Split(line, "(")[1], ")")[0])
		if len(strings.Split(line, " ")) >= 3 {
			children = strings.Split(strings.Replace(strings.Split(line, "->")[1], " ", "", -1), ",")
		}
		programs = append(programs, program{name, weight, children})
	}
	return programs
}

func get_parent(child program, programs []program) program {
	for _, i := range programs {
		for _, candidate := range i.children {
			if child.name == candidate {
				return child
			}
		}
	}
	return child
}

func get_program(name string, programs []program) program {
	program_source := program{"a", 0, []string{}}
	for _, i := range programs {
		if i.name == name {
			program_source = i
		}
	}
	return program_source
}

func get_weight(subject program, programs []program) int {
	total_weight := subject.weight
	if len(subject.children) != 0 {
		total_weight := subject.weight
		for _, child := range subject.children {
			total_weight += get_weight(get_program(child, programs), programs)
		}
	}
	return total_weight
}

func part_1(programs []program) string {
	// find the one program not in any other program's children
	has_parent := false
	for _, candidate := range programs {
		has_parent = false
		for _, prospective_parent := range programs {
			for _, child := range prospective_parent.children {
				if child == candidate.name {
					has_parent = true
				}
			}
		}
		if !has_parent {
			return candidate.name
		}
	}
	return "not successful"
}

func part_2(programs []program) int {
	return 1
}

func main() {
	fmt.Printf("Part 1: %s\n", part_1(get_input()))
}
