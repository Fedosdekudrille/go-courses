package main

import "fmt"

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

func foreach(nums []int, action func(int)) {
	for _, value := range nums {
		action(value)
	}
}

func main() {
	nums := []int{1, 2, 16}

	fmt.Println(where(nums, func(num int) bool { return num%2 == 0 }))

	foreach(nums, func(num int) { fmt.Print(num, " ") })
}
