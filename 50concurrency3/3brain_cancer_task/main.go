package main

import (
	"time"
)

func Process(i int) {
	println("processing ", i)

	time.Sleep(time.Second) // simulating a workload
}

func main() {
	var input = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, i := range input {
		i := i

		go Process(i)
	}

	time.Sleep(time.Second * 2)
}
