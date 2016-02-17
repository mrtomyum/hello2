package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Firstname string
	Lastname  string
}

func (u User) Fullname() string {
	return u.Firstname + " " + u.Lastname
}

var users []User

func main() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		bJson, err := json.Marshal(users)
		if err != nil {
			fmt.Println(err)
		} else {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			fmt.Fprintf(w, string(bJson))
		}
		// fmt.Fprintf(w, "Hello")
	})
	go http.ListenAndServe(":8080", nil)

	choice := 1
	menuList := `Menu
    1) Add User
    2) JSON
    3) ออกโปรแกรม
    Enter: `
	for {
		fmt.Print(menuList)
		fmt.Scan(&choice)
		menu(choice)
	}
}

func menu(choice int) {
	switch choice {
	case 1:
		var f, l string
		fmt.Print("Please enter User.Firstname: ")
		fmt.Scan(&f)
		fmt.Print("Please enter User.Lastname: ")
		fmt.Scan(&l)
		// เก็บลง struct
		u := User{
			Firstname: f,
			Lastname:  l,
		}
		users = append(users, u)
		fmt.Println("Full name: ", u.Fullname())
		fmt.Println(users)
	case 2:
		bJson, err := json.Marshal(users)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(bJson))

		}
	case 3:
		fmt.Println("Exit Program...")
		break
	}
}
