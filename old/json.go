package main

import (
    "fmt"
    "encoding/json"
)

func main() {
    var data interface{}
    reqJSON := "{\"name\": \"Kasem\"}"
    err := json.Unmarshal([]byte(reqJSON), &data)
    if err != nil {
        fmt.Println("Error ", err)
    }
    fmt.Println(data)
}
