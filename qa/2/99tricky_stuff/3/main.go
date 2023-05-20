package main

import (
	"fmt"
)

// 2. // Будет ли напечатан “ok” ?
func main() {
	defer func() {
		recover()
	}()

	if true {
		panic("test panic")
	}
	fmt.Println("ok")
}
