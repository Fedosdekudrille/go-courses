package main

import "fmt"

func main() {
	var num int
	fmt.Scanln(&num)
	if num >= 1_000_000_000 {
		fmt.Println("Input error")
	} else {
		var biggestNum int = 0
		var currentNum int
		for num > 0 {
			currentNum = num % 10
			if currentNum > biggestNum {
				biggestNum = currentNum
			}
			num /= 10
		}
		fmt.Println(biggestNum)
	}
}
