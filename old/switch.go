package main

import (
	"fmt"
)

func main() {
	var x int
	for {
		fmt.Scan(&x)
		switch {
        default:
            fmt.Println(x)
		case x %15 == 0:
			fmt.Println("FizzBuzz")
		case x %3 == 0:
			fmt.Println("Fizz")
		case x %5 == 0:
			fmt.Println("Buzz")
		}
	}
}
