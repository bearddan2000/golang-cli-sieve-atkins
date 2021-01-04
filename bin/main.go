package main

import (
	"fmt"
)

func SieveOfAtkin(limit int) {
	// 2 and 3 are known to be prime
	if limit > 2 {
		fmt.Printf("[OUTPUT] %d ", 2)
	}
	if limit > 3 {
		fmt.Printf("%d ", 3)
	}

	// Initialise the sieve array with
	// false values
	sieve := make([]bool, limit, limit*2)

	for i := 0; i < limit-1; i++ {
		sieve[i] = false
	}

	/* Mark siev[n] is true if one of the
	   following is true:
	   a) n = (4*x*x)+(y*y) has odd number
	      of solutions, i.e., there exist
	      odd number of distinct pairs
	      (x, y) that satisfy the equation
	      and    n % 12 = 1 or n % 12 = 5.
	   b) n = (3*x*x)+(y*y) has odd number
	      of solutions and n % 12 = 7
	   c) n = (3*x*x)-(y*y) has odd number
	      of solutions, x > y and n % 12 = 11 */

	for x := 1; x*x < limit; x++ {
		for y := 1; y*y < limit; y++ {

			// Main part of Sieve of Atkin
			n := (4 * x * x) + (y * y)
			if n <= limit && (n%12 == 1 || n%12 == 5) {
				sieve[n] = sieve[n] != true
			}

			n = (3 * x * x) + (y * y)
			if n <= limit && n%12 == 7 {
				sieve[n] = sieve[n] != true
			}

			n = (3 * x * x) - (y * y)
			if x > y && n <= limit && n%12 == 11 {
				sieve[n] = sieve[n] != true
			}
		}
	}

	// Mark all multiples of squares as
	// non-prime
	for r := 5; r*r < limit; r++ {
		if sieve[r] == true {
			for i := r * r; i < limit-1; i += r * r {
				sieve[i] = false
			}
		}
	}

	// Print primes using sieve[]
	for a := 5; a < limit; a++ {
		if sieve[a] == true {
			fmt.Printf("%d ", a)
		}
	}
	fmt.Println("")
}
func main() {
	input := 20
	fmt.Printf("[INPUT] %d\n", input)
	SieveOfAtkin(input)
}
