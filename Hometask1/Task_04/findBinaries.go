package main

import "fmt"

func main() {
	var a, b int

	fmt.Println("Enter first variable")
	fmt.Scanln(&a)
	fmt.Println("Enter second variable")
	fmt.Scanln(&b)

	if a > b {
		a, b = b, a
	}

	for ; a <= b; a++ {
		curNum := a
		onesNum := 0
		for curNum != 0 {
			onesNum += curNum & 1
			curNum >>= 1
		}
		if onesNum == 4 {
			fmt.Println(a)
		}
	}
}
