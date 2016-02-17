package main

import (
    "fmt"
)

func main()  {
    var x int
    f := "Fizz"
    b := "Buzz"
    fmt.Print("Enter number: ")
    fmt.Scanln(&x)
    if x%3 == 0 && x%5 ==0 {
        fmt.Println(f+b)
    } else if x%3 == 0 && x%5 != 0 {
        fmt.Println(f)
    } else if x%3 !=0 && x%5 == 0 {
        fmt.Println(b)
    } else {
        fmt.Println(x)
    }
}
