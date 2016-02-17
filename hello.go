package hello

import (
	"fmt"
)

func init() {
	fmt.Println("Init Hello")
}

func Say(name string) {
	fmt.Println("Hello", name)
}
