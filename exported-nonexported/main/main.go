package main

import (
	"fmt"

	"go-practice/exported-nonexported/helper"
)

func main() {
	fmt.Println(helper.Exported)
	// fmt.Println(helper.notExported) doesnt work
}
