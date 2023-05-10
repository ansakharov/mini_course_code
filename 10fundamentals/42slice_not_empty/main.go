package main

import (
	"fmt"
)

func main() {
	// very common mistake
	sl := make([]int64, 10)

	for i := int64(0); i <= 9; i++ {
		sl[i] = i + 1
		//sl = append(sl, i)
	}

	// What output here????
	fmt.Println(sl)
}
