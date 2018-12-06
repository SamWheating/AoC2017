package main

import "fmt"

// Checking if the 16 lowest bits of a number match is equivalent to
// checking if the numbers mod 2^16 match...

func main() {

	// Puzzle inputs / Magic numbers
	A := 591
	AFactor := 16807
	B := 393
	BFactor := 48271

	count := 0
	agreed := 0

	for {

		A = (A * AFactor) % 2147483647
		B = (B * BFactor) % 2147483647
		if A%65536 == B%65536 {
			agreed++
		}
		count++
		if count > 40000000 {
			fmt.Println("Part 1: ", agreed)
			break
		}
	}

	A = 591
	B = 393

	count = 0
	agreed = 0

	for {

		for {
			A = (A * AFactor) % 2147483647
			if A%4 == 0 {
				break
			}
		}

		for {
			B = (B * BFactor) % 2147483647
			if B%8 == 0 {
				break
			}
		}

		if A%65536 == B%65536 {
			agreed++
		}
		count++
		if count > 5000000 {
			fmt.Println("Part 2: ", agreed)
			break
		}
	}

}
