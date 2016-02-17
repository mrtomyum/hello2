package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Firstname string
}

func main() {
	pUser :=
		&User{
			Firstname: "Kasem",
		}

	users := []User{
		User{
			Firstname: "Kittithat",
		},
	}
	fmt.Println(pUser)
	fmt.Println(users)
	u, err := json.Marshal(pUser)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("JSON output: ", string(u))
	}
	jsonSlice, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("JSON output: ", string(jsonSlice))
	}

}
