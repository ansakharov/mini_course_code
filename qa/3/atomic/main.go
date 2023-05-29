package main

import (
	"fmt"
	"sync/atomic"
)

// atomic.SwapInt64. Разве операция присваивания в go не атомарная?
func main() {
	var val int64 = 10
	newVal := int64(20)

	// volatile
	prevVal := atomic.SwapInt64(&val, newVal)

	fmt.Println(prevVal) // Выводит: 10
	fmt.Println(val)     // Выводит: 20

}
