package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		fmt.Print(e)
		panic(e)
	}
}

func has_duplicates(phrase []string) bool {
	for i, _ := range phrase {
		for j, _ := range phrase {
			if phrase[i] == phrase[j] && i != j {
				return true
			}
		}
	}
	return false
}

func count_letters(word string, letter string) int {
	count := 0
	for i := 0; i < len(word); i++ {
		if word[i] == letter[0] {
			count += 1
		}
	}
	return count
}

func are_anagrams(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if count_letters(a, string(a[i])) != count_letters(b, string(a[i])) {
			return false
		}
	}

	return true
}

func has_anagrams(phrase []string) bool {
	for i, _ := range phrase {
		for j, _ := range phrase {
			if are_anagrams(phrase[i], phrase[j]) && i != j {
				return true
			}
		}
	}
	return false
}

func main() {

	// Load + parse the input into slices of strings.
	dat, err := ioutil.ReadFile("day4/day4.txt")
	check(err)
	lines := strings.Split(string(dat)[:len(string(dat))-1], "\n")
	words := [][]string{}

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		line = strings.Replace(line, "\n", "", 1)
		var passphrase = strings.Split(line, " ")
		words = append(words, passphrase)
	}

	// Part 1 - Count the number of passphrases without duplicates
	valid := 0
	for _, i := range words {
		if !has_duplicates(i) {
			valid += 1
		}
	}
	fmt.Printf("Part 1: %d\n", valid)

	// Part 2 - Count the number of passphrases without duplicates or anagrams
	valid = 0
	for _, i := range words {
		if !has_anagrams(i) {
			valid += 1
		}
	}
	fmt.Printf("Part 2	: %d\n", valid)
}
