package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var mem runtime.MemStats

	runtime.ReadMemStats(&mem)
	fmt.Printf("Before get(): Alloc = %v MiB\n", mem.Alloc/1024/1024)

	resOfRes := make([][]int, 0, 100)

	for i := 0; i < 50; i++ {
		res := getLast()

		runtime.GC()
		runtime.ReadMemStats(&mem)
		fmt.Printf("After getAll(): Alloc = %v MiB, slice: %v\n", mem.Alloc/1024/1024, res)

		resOfRes = append(resOfRes, res)
	}

	fmt.Println()

	_ = resOfRes

	// Now go into loop and see that pointers are alive
	for i := 0; i < 20000; i++ {
		runtime.GC()
		runtime.ReadMemStats(&mem)
		fmt.Printf(
			"GC can't clear underlaying array, because slice points on it elem\nAlloc = %v MiB\n\n",
			mem.Alloc/1024/1024,
		)
		time.Sleep(time.Second * 2)
	}

	// remove this print and then GC can remove all slices way before
	fmt.Println(resOfRes)
}

func getLast() []int {
	sl := make([]int, 0, 100000)

	for i := 0; i < 100000; i++ {
		sl = append(sl, i)
	}

	// return last element
	return sl[99999:]

	/*
		Instead copy elem to new slice and old array will be removed by GC eventually.
		lastElem := make([]*forTest, 1)
		lastElem[0] = sl[99999]

		return lastElem
	*/

}
