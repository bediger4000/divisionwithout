package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	m, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("n %d\t%16b\n", n, n)
	fmt.Printf("m %d\t%16b\n", m, m)

	q := divide(uint16(n), uint16(m))

	fmt.Printf("\nq %d\t%16b\n", q, q)
}

func divide(numerator uint16, denominator uint16) uint16 {

	if denominator > numerator {
		return 0
	}

	if denominator == numerator {
		return 1
	}

	var position uint16

	position = 1

	for denominator <= numerator {
		denominator <<= 1
		position <<= 1
	}

	denominator >>= 1
	position >>= 1

	var answer uint16

	for position != 0 {
		if numerator >= denominator {
			numerator -= denominator
			answer |= position
		}
		position >>= 1
		denominator >>= 1
	}

	return answer
}
