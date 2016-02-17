package main

import(
    "fmt"
)

func main() {
    var nums [5]int = [5]int {1,2,3,4,5}
    fmt.Println(nums, len(nums), nums[0])
    for i, n := range nums {
        fmt.Println(i, ":", n)
    }
}
