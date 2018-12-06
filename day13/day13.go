package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getInputs() map[int]map[string]int {
	dat, err := ioutil.ReadFile("day13/day13.txt")
	//dat, err := ioutil.ReadFile("day13/day13_sample.txt")
	check(err)
	input := strings.Split(string(dat)[:(len(string(dat))-1)], "\n")

	// Each firewall has a corresponding {height: int, position: int}
	firewalls := make(map[int]map[string]int)
	for _, wall := range input {
		depth, err := strconv.Atoi(strings.Split(wall, ":")[0])
		check(err)
		height, err := strconv.Atoi(strings.Split(wall, " ")[1])
		firewalls[depth] = map[string]int{"height": height, "position": 1, "direction": 1}
	}
	return firewalls
}

func sendPacket(firewalls map[int]map[string]int) int {
	// Calculates the severity of a trip through the firewalls starting at t=0
	severity := 0
	for i, firewall := range firewalls {
		if i%(2*(firewall["height"]-1)) == 0 {
			severity += i * firewall["height"]
		}
	}
	return severity
}

func clearPacket(firewalls map[int]map[string]int, delay int) bool {
	// Tests whether a packet makes it through the firewalls unscathed
	for i, firewall := range firewalls {
		if (delay+i)%(2*(firewall["height"]-1)) == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Part 1: ", sendPacket(getInputs()))
	delay := 0
	firewalls := getInputs()
	for {
		if clearPacket(firewalls, delay) {
			fmt.Println("Part 2: ", delay)
			break
		}
		delay++
	}

}
