package main

import "fmt"

func isPrime(x *int) bool {
	is := true

	if *x == 0 || *x == 1 {
		is = false
	}
	if *x%2 == 0 && *x > 2 {
		is = false
	}
	for d := 3; d*d <= *x; d += 2 {
		if *x%d == 0 {
			is = false
		}
	}

	return is
}

func main() {
	x := 5
	fmt.Print(isPrime(&x))

	var y int = 6
	fmt.Println(isPrime(&y))
}
