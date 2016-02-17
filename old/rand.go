package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().Unix())
    ans := rand.Intn(10)+1
    count := GameLoop(ans)
    fmt.Println("Counter = ",count)
}

func GameLoop(ans int) int{
    var x, count int
    LOOP:
    for {
        fmt.Print("Please enter guess int(x): ")
        fmt.Scan(&x)
        switch {
        case ans == x:
            fmt.Println("Corrected! answer = ", ans)
            break LOOP
        case ans < x:
            fmt.Println("Answer < x")
        case ans > x:
            fmt.Println("Answer > x")
        }
        count++
    }
    return count
}
