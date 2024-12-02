package main

import "fmt"

func takeInput1st(input string) (result string) {
	fmt.Println("Give me your input: ")
	fmt.Scan(&input)
	result = input
	return
}

func takeInput2nd(input string) string {
	fmt.Println("Give me your input: ")
	fmt.Scan(&input)
	return input
}

func main() {
	var input string
	input = takeInput1st(input)
	fmt.Print(input)

	fmt.Println("\n\n")

	var input2 string
	input2 = takeInput2nd(input2)
	fmt.Print(input2)
}
