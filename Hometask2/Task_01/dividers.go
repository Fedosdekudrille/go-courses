package main

import "fmt"

func countSimpleDividers(num int) (string, bool) {

	if num < 2 {
		return "", false
	}

	counter := 0
	for i := 2; i <= num; i++ {
		for num%i == 0 {
			num /= i
			counter++
		}
	}

	if counter%2 == 0 {
		return "even", true
	} else {
		return "odd", true
	}
}

func main() {
	fmt.Println(countSimpleDividers(1))
	fmt.Println(countSimpleDividers(4))
	fmt.Println(countSimpleDividers(5))
	fmt.Println(countSimpleDividers(1000))
}
