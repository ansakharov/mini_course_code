package main

import "fmt"

func main() {
	// 1.Modifiying slices
	/*	nums := []int{1, 2, 3, 4, 5}
		fmt.Println(len(nums), cap(nums)) // len 5 cap 5
		slice := nums[1:3]                // [2, 3], len 2, cap 4
		slice2 := nums[0:1]               // [1], len 1, cap 5
		slice3 := nums[2:2]               // [], len 0, cap 3
		slice4 := nums[4:5]               // [5], len 1, cap 1
		fmt.Println(slice)
		fmt.Println("len:", len(slice), "cap: ", cap(slice))
		fmt.Println("len:", len(slice2), "cap: ", cap(slice2))
		fmt.Println("len:", len(slice3), "cap: ", cap(slice3))
		fmt.Println("len:", len(slice4), "cap: ", cap(slice4))
		slice[0] = 10
		fmt.Println(nums) // 1, 10, 3, 4, 5
		fmt.Println("len: ", len(nums), "cap: ", cap(nums))*/

	fmt.Println()
	//2. Capacity and length
	/*	nums := []int{10, 20, 30, 40, 50, 60, 70, 80}
		slice := nums[1:3:4]
		fmt.Println(slice, len(slice), cap(slice)) // [20, 30], len 2, cap 3

		slice = append(slice, 6)

		// 10, 20, 30, 6, 50, 60, 70, 80
		fmt.Println(nums, len(nums), cap(nums))
		// 20, 30, 6, l3, c3
		fmt.Println(slice, len(slice), cap(slice))*/

	// 3. Copying slices
	/*	nums := []int{1, 2, 3, 4, 5}
		slice := make([]int, 4)
		copy(slice, nums[2:])
		fmt.Println(slice) //3, 4, 5, 0*/

	//	с аппендом слайса bool и byte есть особенности
	a := []byte{'a', 'b'}
	fmt.Println(cap(a)) // 8

	a = append(a, 'c')
	fmt.Println(cap(a)) // 8
}
