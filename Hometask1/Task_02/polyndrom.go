package main

import (
	"fmt"
)

func main() {
	fmt.Println("Insert number")
	var (
		num        uint
		trinaryNum uint = 0
	)
	fmt.Scanln(&num)

	var numPlace uint = 1
	for ; num > 0; numPlace *= 10 {
		trinaryNum += num % 3 * numPlace
		num /= 3
	}

	fmt.Println(trinaryNum)
	var startNumPlace uint = 1
	for {
		numPlace /= 10
		if startNumPlace >= numPlace {
			fmt.Println("Number is polyndrom")
			break
		} else if (trinaryNum/startNumPlace)%10 != (trinaryNum/numPlace)%10 {
			fmt.Println("Number isn't polyndrom")
			break
		}
		startNumPlace *= 10
	}
}
