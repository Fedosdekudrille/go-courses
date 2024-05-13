package main

import (
	"Task_03/pkg/cache"
	"fmt"
	"math/rand/v2"
	"strconv"
)

const (
	USER = iota
	INT
	STRING
)
const size = 10000

type User struct {
	Name string
	Age  int
}

func (u User) String() string {
	return fmt.Sprintf("Name: %s, Age: %d", u.Name, u.Age)
}

func main() {
	c := cache.NewCache(size - 1000) // it'll cause lots of nils at the end, cause data is deleted
	keys := make([]string, size)

	for i := 0; i < size; i++ {
		switch rand.IntN(3) {
		case USER:
			name := make([]rune, rand.IntN(6)+3)
			for j := 0; j < len(name); j++ {
				name[j] = 'a' + rune(rand.IntN(26))
			}
			keys[i] = string(name)
			c.Set(keys[i], User{Name: keys[i], Age: rand.IntN(100)})
		case INT:
			key := strconv.Itoa(rand.Int())
			keys[i] = key
			c.Set(key, rand.Int())
		case STRING:
			s := make([]rune, rand.IntN(20)+5)
			for j := 0; j < len(s); j++ {
				s[j] = 'a' + rune(rand.IntN(26))
			}
			keys[i] = string(s)
			c.Set(keys[i], keys[i])
		}
	}

	for _, key := range keys {
		obj := c.Get(key)
		switch obj.(type) {
		case User:
			user := obj.(User)
			fmt.Println("User:", user)
		case int:
			i := obj.(int)
			fmt.Println("Int:", i)
		case string:
			s := obj.(string)
			fmt.Println("String:", s)
		case nil:
			fmt.Println("nil")
		}
	}
}
