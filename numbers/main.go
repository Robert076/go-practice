package main

import "fmt"

func sameThing() {
	var x int = 5
	fmt.Print(x)

	y := 5
	fmt.Print(y) // same thing
}

func main() {
	sameThing()
}
