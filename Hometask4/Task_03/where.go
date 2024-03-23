package main

import (
	"fmt"
	"time"
)

func Where[T interface{}](elements []T, check func(T) bool) (res []T) {
	if len(elements) == 0 || check == nil {
		return
	}
	l := len(elements)
	part, leftAdd, add := l/8, l%8, 0
	var lBorder int
	ch := [8]chan []T{} // Получают результат, обеспечивают ожидание выполнения
	for i := 0; i < 8; i++ {
		ch[i] = make(chan []T)
		if leftAdd > 0 {
			add = 1
			leftAdd--
		} else {
			add = 0
		}
		go func(elements []T, check func(T) bool, ch chan<- []T) {
			var res []T
			for i := 0; i < len(elements); i++ {
				if check(elements[i]) {
					res = append(res, elements[i])
				}
			}
			ch <- res
		}(elements[lBorder:lBorder+part+add], check, ch[i])
		lBorder += part + add
	}
	for i := 0; i < 8; i++ {
		if v := ch[i]; v != nil {
			res = append(res, <-ch[i]...)
		} else {
			fmt.Println(i)
		}
	}
	return
}

func where(nums []int, check func(int) bool) []int {
	length := len(nums)
	retNums := make([]int, 0, length)
	for i := 0; i < length; i++ {
		if check(nums[i]) {
			retNums = append(retNums, nums[i])
		}
	}

	return retNums
}

func main() {
	fmt.Println(Where([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func(x int) bool { return x%2 == 1 }))
	s := make([]int, 100000000)
	for i := 0; i < 100000000; i++ {
		s = append(s, i)
	}
	t := time.Now()
	Where(s, func(x int) bool { return x%2 == 1 })
	fmt.Println(time.Since(t))
	t = time.Now()
	where(s, func(x int) bool { return x%2 == 1 })
	fmt.Println(time.Since(t))
}
