package main

import (
	"Task_02/internal/server"
	"Task_02/pkg/udb"
	"math/rand/v2"
)

func main() {
	udb.EnsureCreatedDB()
	users := make([]udb.User, 10)
	for i := 0; i < 10; i++ {
		name, surname := make([]rune, rand.IntN(5)+4), make([]rune, rand.IntN(7)+4)
		for j := 0; j < len(name); j++ {
			name[j] = 'a' + rune(rand.IntN(26))
		}
		for j := 0; j < len(surname); j++ {
			surname[j] = 'a' + rune(rand.IntN(26))
		}
		users[i].Parse(string(name) + " " + string(surname))
	}
	udb.AddUsers(users)
	server.MustStartServer()
}
