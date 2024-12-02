package main

import (
	"fmt"
	"reflect"
)

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
	fmt.Println(reflect.TypeOf(z)) // instanceof from java
}
