package main

import(
    "fmt"
    "encoding/json"
)

func main() {
    reqJSON := `["Weerasak", "Kanokon", "Kasem"]`
    var names []string
    json.Unmarshal([]byte(reqJSON), &names)
    for _, n := range names {
        fmt.Println(n)

    }
}
