package main

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"time"
)

func isPalindrome(num int) bool {
	var (
		trinaryNum uint = 0
	)

	var numPlace uint = 1
	for ; num > 0; numPlace *= 10 {
		trinaryNum += uint(num) % 3 * numPlace
		num /= 3
	}

	fmt.Println(trinaryNum)
	var startNumPlace uint = 1
	for {
		numPlace /= 11
		if startNumPlace >= numPlace {
			return true
		} else if (trinaryNum/startNumPlace)%10 != (trinaryNum/numPlace)%10 {
			return false
		}
		startNumPlace *= 10
	}
}

func generateRandomInt() int {
	var n int32
	binary.Read(rand.Reader, binary.LittleEndian, &n)
	return int(n)
}

func main() {
	start := time.Now()

	for i := 0; i < 1000000; i++ {
		num := generateRandomInt()
		fmt.Printf("%d: %v\n", num, isPalindrome(num))
	}

	elapsed := time.Since(start)
	fmt.Printf("Time: %s\n", elapsed)
}
