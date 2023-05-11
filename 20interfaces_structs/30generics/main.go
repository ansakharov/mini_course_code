package main

import (
	"fmt"
)

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
}

func (s *Stack[T]) Pop() T {
	if len(s.data) == 0 {
		panic("stack is empty")
	}
	last := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return last
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

// not method but func
func UniqueSlice[T comparable](arr []T) []T {
	if arr == nil {
		return nil
	}

	arrMap := make(map[T]struct{}, len(arr))
	result := make([]T, 0, len(arr))

	for _, value := range arr {
		if _, ok := arrMap[value]; !ok {
			arrMap[value] = struct{}{}
			result = append(result, value)
		}
	}

	return result
}

func main() {
	// must declare type
	intStack := &Stack[int]{}
	intStack.Push(1)
	intStack.Push(2)

	fmt.Println(intStack.Pop(), intStack.IsEmpty())
	fmt.Println(intStack.Pop(), intStack.IsEmpty())

	fmt.Println()

	stringStack := &Stack[string]{}
	stringStack.Push("hello")
	stringStack.Push("world")

	fmt.Println(stringStack.Pop(), stringStack.IsEmpty())
	fmt.Println(stringStack.Pop(), stringStack.IsEmpty())

	fmt.Println()

	anyStack := &Stack[any]{}
	anyStack.Push([]int64{3, 2, 1})
	anyStack.Push([]string{"three", "two", "one"})

	fmt.Println(anyStack.Pop(), anyStack.IsEmpty())
	fmt.Println(anyStack.Pop(), anyStack.IsEmpty())

	// show unique after
	fmt.Println(UniqueSlice([]int64{1, 2, 3, 4, 5, 1, 2, 3, 6, 4, 3, 7, 2, 34}))
	// fmt.Println(UniqueSlice([]func() int {}))

}
