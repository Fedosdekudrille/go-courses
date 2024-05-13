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
	ch := make(chan []T, 8)
	for i := 0; i < 8; i++ {
		if leftAdd > 0 { // Если не кратно 8, равномерно распределяется между горутинами
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
		}(elements[lBorder:lBorder+part+add], check, ch)
		lBorder += part + add
	}
	for i := 0; i < 8; i++ {
		res = append(res, <-ch...)
	}
	return
}

func where(nums []int, check func(int) bool) []int {
	if len(nums) == 0 || check == nil {
		return []int{}
	}
	length := len(nums)
	retNums := make([]int, 0, length)
	for i := 0; i < length; i++ {
		if check(nums[i]) {
			retNums = append(retNums, nums[i])
		}
	}

	return retNums
}

type Student struct {
	Iq   int
	Name string
	Mark int
}

func (s Student) Compare(other Student) int {
	sTotal := s.Iq + s.Mark*10
	otherTotal := other.Iq + other.Mark*10
	if s.Name == other.Name || sTotal == otherTotal {
		return 0
	} else if sTotal > otherTotal {
		return 1
	}
	return -1
}
func main() {
	defer func() {
		if err := recover(); err != nil {
			println(err)
		}
	}()
	fmt.Println(Where([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func(x int) bool { return x%2 == 1 }))
	sLen := 10000000 //Равенство достигается примерно при 500000 длине
	s := make([]int, sLen)
	for i := 0; i < sLen; i++ {
		s = append(s, i)
	}
	t := time.Now()
	Where(s, func(x int) bool { return x%2 == 1 })
	fmt.Println(time.Since(t))
	t = time.Now()
	where(s, func(x int) bool { return x%2 == 1 })
	fmt.Println(time.Since(t))

	avgStud := Student{Iq: 110, Name: "Ivan", Mark: 7}
	fmt.Println(Where([]Student{{Iq: 90, Name: "Volodya", Mark: 9}, {Iq: 125, Name: "Petya", Mark: 5}, avgStud}, func(s Student) bool { return s.Compare(avgStud) >= 0 }))

	Where([]int{1, 2, 3}, func(x int) bool { panic("Учебная тревога") }) // Обработка паники на внешнем коде
}
