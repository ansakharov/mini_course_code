package main

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"
)

// в каких случаях стоит использовать буферизованные, а в каких небуферизованные каналы?

/// buffer async stop blocking due different speed of goroutines

// unbuffer strict sync of goroutines

func main() {

	// var wg sync.WaitGroup
	wg := sync.WaitGroup{}

	// <- 1123 <- 1124

	// 1123 <-

	// 1124 <-
	background, cancel := context.WithCancel(context.Background())
	fmt.Println("background created")
	go func() {
		select {
		case <-background.Done():
		case <-time.After(time.Millisecond):
		}

		fmt.Println("gor100 finished")
	}()
	go func() {
		select {
		case <-background.Done():
		case <-time.After(time.Millisecond):
		}

		fmt.Println("gor101 finished")
	}()
	go func() {
		select {
		case <-background.Done():
		case <-time.After(time.Second):
		}

		fmt.Println("gor102 finished")
	}()

	time.Sleep(time.Second * 3)
	cancel()

	time.Sleep(time.Second * 3)

	os.Exit(1)

	ch := make(chan int, 10)
	go func() {
		for i := 0; i < 2; i++ {
			ch <- i
		}
		close(ch)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for value := range ch {
			fmt.Printf("value %d, gor num 1\n", value)
		}
		// http get order
		// go to db
		// produce msg to kafka

	}()
	wg.Add(1)

	go func() {
		defer wg.Done()
		for value := range ch {
			fmt.Printf("value %d, gor num 2\n", value)
		}
	}()
	wg.Add(1)

	go func() {
		defer wg.Done()
		for value := range ch {
			fmt.Printf("value %d, gor num 3\n", value)
		}
	}()

	wg.Wait()
}

func main1() {

	ch := make(chan struct{})

	go func() {
		time.Sleep(time.Second * 3)
		<-ch
		fmt.Println("time to work")
	}()

	ch <- struct{}{}
	fmt.Println("in main")

	time.Sleep(time.Second)
}
