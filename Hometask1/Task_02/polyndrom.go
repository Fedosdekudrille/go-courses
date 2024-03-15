package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func main() {
	start := time.Now()
	for i := 0; i < 100_000; i++ {
		pain(rand.Int(rand.Reader, big.NewInt(2)))
	}
	fmt.Println(time.Now().Sub(start))
}

func pain(bigNum *big.Int, _ error) {
	num := bigNum.Int64()
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
