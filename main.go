package main

import (
	"fmt"
	"reflect"
)

func test() {
	var output string
	fmt.Scan(&output) // just like in c/c++
	fmt.Print(output)
	fmt.Print(&output) // just like in c/c++
}

func main() {
	// go is statically typed
	// meaning that variables always have a specific type and cannot change
	var x string = "Hello world!"
	fmt.Println(x)

	var y string
	y = "Test"
	fmt.Println(y)

	var z = "Hello"
	z += " World!"

	// find the type of a variable in 2 ways
	fmt.Println(reflect.TypeOf(z))
	fmt.Printf("%T\n", z)

	const xx string = "test"
	// xx = "test" compilation error
	fmt.Print(xx)
	test()
}
