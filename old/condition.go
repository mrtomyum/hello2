package main

import (
	"fmt"
)

func main() {
	name := "tom"

	if num := 2; num == 1 {
		fmt.Println("One", num)
	} else {
		fmt.Println("Two", num)
	}

	if name != "Weerasak" {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not Equal")
	}
}
