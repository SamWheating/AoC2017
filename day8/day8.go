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

func get_input() [][]string {
	dat, err := ioutil.ReadFile("day8/day8.txt")
	check(err)
	input := string(dat)[:(len(string(dat)) - 1)]
	lines := strings.Split(input, "\n")
	commands := [][]string{}
	for _, line := range lines {
		commands = append(commands, strings.Split(line, " "))
	}
	return commands
}

func init_registers(commands [][]string) map[string]int {
	registers := make(map[string]int)
	for _, command := range commands {
		registers[command[0]] = 0
	}
	return registers
}

func evaluate_conditional(command []string, registers map[string]int) bool {
	a := registers[command[4]]
	operator := command[5]
	b, err := strconv.Atoi(command[6])
	check(err)
	switch operator {
	case ">":
		return a > b
	case "<":
		return a < b
	case "==":
		return a == b
	case "!=":
		return a != b
	case ">=":
		return a >= b
	case "<=":
		return a <= b
	}
	return false
}

func execute_instruction(command []string, registers map[string]int) map[string]int {
	register := command[0]
	operator := command[1]
	value, err := strconv.Atoi(command[2])
	check(err)
	switch operator {
	case "inc":
		registers[register] += value
	case "dec":
		registers[register] -= value
	}
	return registers
}

func max_register(registers map[string]int) int {
	max_register := 0
	for _, value := range registers {
		if value > max_register {
			max_register = value
		}
	}
	return max_register
}

func part_1(commands [][]string) int {

	registers := init_registers(commands)
	for _, command := range commands {
		if evaluate_conditional(command, registers) {
			registers = execute_instruction(command, registers)
		}
	}
	return max_register(registers)
}

func part_2(commands [][]string) int {
	registers := init_registers(commands)
	max_seen := 0
	for _, command := range commands {
		if evaluate_conditional(command, registers) {
			registers = execute_instruction(command, registers)
			if max_register(registers) > max_seen {
				max_seen = max_register(registers)
			}
		}
	}
	return max_seen
}

func main() {
	fmt.Printf("Part 1: %d\n", part_1(get_input()))
	fmt.Printf("Part 2: %d\n", part_2(get_input()))
}
