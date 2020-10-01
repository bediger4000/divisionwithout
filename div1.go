package main

/*
Implement integer division without using the division operator.
Your function should return a tuple of
(dividend,remainder)
and it should take two numbers,
the product and divisor.
*/

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	product, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	divisor, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	dividend, remainder := divide(product, divisor)

	fmt.Printf("%d/%d = %d rem %d\n", product, divisor, dividend, remainder)

	// Double-check using built-in integer arithmetic
	dividend2 := product / divisor
	remainder2 := product % divisor
	fmt.Printf("integer arithmetic: %d/%d = %d rem %d\n", product, divisor, dividend2, remainder2)
}

func divide(product, divisor int64) (int64, int64) {

	var count int64

	sign := int64(1)
	if product < 0 {
		sign = -1
		product = -product
	}
	if divisor < 0 {
		sign = -sign
		divisor = -divisor
	}

	for product >= divisor {
		product -= divisor
		count++
	}

	return sign * count, sign * product
}
