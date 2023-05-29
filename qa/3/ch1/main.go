package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 1)

	ch <- ""

	close(ch)
	k, v := <-ch
	fmt.Println(k, v)
	k, v = <-ch
	fmt.Println(k, v)

	fmt.Println("end")

}
