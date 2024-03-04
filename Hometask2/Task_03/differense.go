package main

import "fmt"

func makeSliceBetween(first int, sec int) []int {
	if first > sec {
		first, sec = sec, first
	}
	slice := make([]int, 0, sec-first+1)
	for ; first <= sec; first++ {
		slice = append(slice, first)
	}
	return slice
}

func sequense(nums ...int) []int {
	var res []int
	switch len(nums) {
	case 0:
		res = make([]int, 0)
	case 1:
		res = makeSliceBetween(0, nums[0])
	case 2:
		res = makeSliceBetween(nums[0], nums[1])
	default:
		res = nums
	}
	return res
}

func main() {
	fmt.Println(sequense())
	fmt.Println(sequense(-6))
	fmt.Println(sequense(6, 9))
	fmt.Println(sequense(3, 5, 5, 2))
}
