package main

import "fmt"

func takeInput1st(input string) (result string) {
	fmt.Println("Give me your 1st input: ")
	fmt.Scan(&input)
	result = input
	return
}

func takeInput2nd(input string) string {
	fmt.Println("Give me your 2nd input: ")
	fmt.Scan(&input)
	fmt.Println(&input) // incorrect since we make a pass by value it makes a new addr here
	return input
}

func takeInput3rd(input *string) string {
	fmt.Println("Give me your 3rd input: ")
	fmt.Scan(input)
	fmt.Println(input) // this is correct. we pass an address so its passed by reference not value
	return *input      // therefore we work with the same address we have in the main
}

func main() {
	var input string
	input = takeInput1st(input)
	fmt.Print(input)

	fmt.Println("\n\n	")

	var input2 string
	input2 = takeInput2nd(input2)
	fmt.Println(&input2)

	fmt.Println("\n\n	")

	var input3 string
	input3 = takeInput3rd(&input3)
	fmt.Println(&input3)
}
