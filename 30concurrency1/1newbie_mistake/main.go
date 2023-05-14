package main

import (
	"fmt"
	"sync"
)

func main() {
	//channelSync()
	wgSync()
}

func channelSync() {
	length := 100
	sl := make([]struct{}, length)
	ch := make(chan struct{})

	for key, value := range sl {
		key, value := int64(key), value
		go func() {
			fmt.Println(key, value)
			ch <- struct{}{}
		}()
	}

	//counter := 0
	/*for range ch {
		counter++
		if counter == length {
			close(ch)
		}
	}*/

	for i := 0; i < length-1; i++ {
		<-ch
	}

	fmt.Println("end")
}

func wgSync() {
	sl := make([]struct{}, 100)
	wg := sync.WaitGroup{}

	for key, value := range sl {
		key, value := int64(key), value

		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(key, value)
		}()
	}

	wg.Wait()
	fmt.Println("end")
}
