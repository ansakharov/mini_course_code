package main

import (
	"fmt"
	"runtime"
)

// what happens here?
func main() {
	ch := make(chan int)

	fmt.Println(runtime.NumGoroutine())
	ch <- 1
}
