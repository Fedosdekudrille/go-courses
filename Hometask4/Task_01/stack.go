package main

import (
	"errors"
)

type Stack[T interface{}] struct { //Мог обойтись одной структурой, но решил делать "по правилам"
	current *node[T]
}

type node[T interface{}] struct {
	Value T
	next  *node[T]
}

func (s *Stack[T]) IsEmpty() bool {
	if s == nil || s.current == nil {
		return true
	}
	return false
}

func (s *Stack[T]) Push(value T) {
	if s == nil {
		return
	}
	s.current = &node[T]{
		Value: value,
		next:  s.current,
	}
}

func (s *Stack[T]) Pop() (val T, err error) {
	if s.IsEmpty() {
		return val, errors.New("stack is empty")
	}
	val = s.current.Value
	s.current = s.current.next
	return val, nil
}

func main() {
	st := Stack[int]{}
	st.Push(1)
	st.Push(2)
	st.Push(3)
	st.Push(4)
	for !st.IsEmpty() {
		if val, err := st.Pop(); err == nil {
			println(val)
		}
	}
	_, err := st.Pop()
	println(err.Error())
	s := Stack[error]{}
	s.Push(nil)
}
