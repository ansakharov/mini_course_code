package main

import (
	"fmt"
	// "reflect"
)

/*
	type SliceHeader struct {
		Data uintptr
		Len  int
		Cap  int
	}
*/
func main() {
	nums := make([]int, 0, 2) // Len 0, cap 2
	fmt.Println(nums, len(nums), cap(nums))

	nums = appendSlice(nums, 1024)
	fmt.Println(nums, len(nums), cap(nums)) // 1024

	mutateSlice(nums, 0, 512)
	fmt.Println(nums, len(nums), cap(nums))

}
func appendSlice(sl []int, val int) []int {
	sl = append(sl, val)

	return sl
}
func mutateSlice(sl []int, idx, val int) {
	sl2 := make([]int, 2, 2)
	copy(sl2, sl)
	//sl2 := sl
	sl2[idx] = val
	fmt.Println("sl2", sl2)
}
