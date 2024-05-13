package main

import (
	"Task_01/internal/ufile"
	"Task_01/pkg/udb"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	err := udb.EnsureCreatedDB()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = udb.ClearUsers()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	str, err := ufile.ReadFile("./assets/persons.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	users, err := ufile.ParseUsers(str)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = udb.AddUsers(users)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	users, err = udb.GetUsers()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, user := range users {
		fmt.Println(user.String())
	}
}
