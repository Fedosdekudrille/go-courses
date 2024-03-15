package main

import "fmt"

func group(groupingMap map[byte]string) (res map[byte][]string) {
	res = make(map[byte][]string, 10)
	for index, value := range groupingMap {
		i := index % 10
		res[i] = append(res[i], value)
	}
	return res
}

func main() {
	fmt.Println(group(map[byte]string{32: "smth", 62: "smth too", 1: "another thing"}))
}
