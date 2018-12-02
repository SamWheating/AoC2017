package main

import(
	"fmt"
	"strings"
	"io/ioutil"
	"strconv"
)

type program struct {
	name string
	weight int
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

	for _, line := range(lines) {
		children := []string{}
		line = strings.Replace(line, "\n", "", 1)
		name := strings.Split(line, " ")[0]
		weight, _ := strconv.Atoi(strings.Split(strings.Split(line, "(")[1], ")")[0])
		if len(strings.Split(line, " ")) >= 3 {
			children = strings.Split(strings.Replace(strings.Split(line, "->")[1], " ", "", -1),",")
		}
		programs = append(programs, program{name, weight, children})
	}
	return programs
}


func part_1(programs []program) string {
	// find the one program not in any other program's children
	has_parent := false
	for _, candidate := range(programs){
		has_parent = false
		for _, prospective_parent := range(programs){
			for _, child := range(prospective_parent.children){
				if child == candidate.name{
					has_parent =true
				}
			}
		}
		if !has_parent {
			return candidate.name
		}
	}
	return "not successful"
}

func main(){
	fmt.Print(part_1(get_input()))
}