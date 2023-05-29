package main

import (
	"fmt"
)

func main() {
	var sl []int

	sl2 := make([]int, 0, 2)

	fmt.Println(sl == nil)
	fmt.Println(sl2 == nil)
	fmt.Println(len(sl2) == 0)
}
