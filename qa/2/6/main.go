package main

import (
	"fmt"
)

func main() {
	s := make([]int, 0, 10) // Заполняем слайс

	for i := 0; i < 10; i++ {
		s = append(s, i)
	}
	// Что выведет этот принт

	// 0-9
	fmt.Printf("address: %p, length: %d, capacity: %d\n", s, len(s), cap(s))

	s = s[:0] // Что произойдет ? Что будет представлять собой слайс s ?

	// Что выведет этот принт
	fmt.Printf("address: %p, length: %d, capacity: %d\n", s, len(s), cap(s)) // [], 0, 10
	for i := 0; i < 10; i++ {
		s = append(s, i)
	}
	fmt.Printf("address: %p, length: %d, capacity: %d\n", s, len(s), cap(s)) //0-9, 10, 10
}
