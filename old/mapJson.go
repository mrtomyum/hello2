package main

import (
    "fmt"
    "encoding/json"
    "io/ioutil"
    "net/http"
)

func main() {
    // var url string = "www.google.com/finance/info?q=INDEXBKK:SET"
    var index string
    url := "http://www.google.com/finance/info?q="

    fmt.Scanln(&index)
    res, err := http.Get(url+ index)
    if err != nil {
        fmt.Println(err)
    }

    reqJson, _ := ioutil.ReadAll(res.Body)
    res.Body.Close()

    var data []map[string]string
    json.Unmarshal(reqJson[3:], &data)
    fmt.Println("Index of", index, ":", data[0]["l_fix"])
}
