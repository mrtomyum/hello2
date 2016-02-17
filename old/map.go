package main

import (
    "fmt"
    "encoding/json"
)

func main() {
    var m []map[string]string
    reqJson := `[
        {"firstName": "เกษม", "lastName": "อานนทวิลาศ"},
        {"firstName": "สมชาย", "lastName": "เด็กดี"}
    ]`
    json.Unmarshal([]byte(reqJson), &m)
    for _, v := range m {
        fmt.Println(v["firstName"])
        fmt.Println(v["lastName"])

    }
}
