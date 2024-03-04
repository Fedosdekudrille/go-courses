package main

import "fmt"

func group(groupingMap map[byte]string) (res map[byte][]string) {
	res = make(map[byte][]string, 10)
	for i := 0; i < 10; i++ {
		str := make([]string, 0, len(groupingMap)/10)
		for key, value := range groupingMap {
			if key%10 == byte(i) {
				str = append(str, value)
			}
		}
		if len(str) > 0 {
			res[byte(i)] = str
		}
	}
	return res
}

func main() {
	fmt.Println(group(map[byte]string{32: "smth", 62: "smth too", 1: "another thing"}))
}
