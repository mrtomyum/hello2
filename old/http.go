package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var url string = "http://www.google.com/finance/info?q=INDEXBKK:"
	var a string
	fmt.Print("Enter parameter>> ")
	fmt.Scanf("%s", &a)
	res, _ := http.Get(url + a)
	// if err != nil {
	// 	// fmt.Errorf("Error>> %s", err)
	// 	fmt.Println(err)
	// }
	body, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

	fmt.Printf("%s\n", body)
}
