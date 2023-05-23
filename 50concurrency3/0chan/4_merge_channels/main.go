package main

import (
	"fmt"
	"sync"
)

// merge two channels
func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 20)

	ch1 <- 1
	ch2 <- 2
	ch2 <- 4
	close(ch1)
	close(ch2)

	ch3 := merge[int](ch1, ch2)

	for val := range ch3 {
		fmt.Println(val)
	}
	// errors:
	//close(ch3) // panic
	// ch3 <- 1
}

func merge[T any](chans ...chan T) chan T {
	result := make(chan T, 1024)
	wg := sync.WaitGroup{}

	for _, ch := range chans {
		ch := ch
		wg.Add(1)
		go func() {
			defer wg.Done()
			for value := range ch {
				result <- value
			}
		}()
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	return result
}

func syncMerge[T any](chans ...chan T) chan T {
	l := 0
	for _, singleCh := range chans {
		l += len(singleCh) // 3
	}

	result := make(chan T, l)
	for _, singleCh := range chans {
		for val := range singleCh {
			result <- val
		}
	}
	close(result)

	return result
}
