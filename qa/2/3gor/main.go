//go:build integration

package main

import (
	"fmt"
	"sync"
	"sync/atomic"

	"golang.org/x/sync/errgroup"
)

// Вопросы по горутинам:
// что мы увидим в stdout?

func main() {
	// testGoroutines3()
	// testGoroutines4()
	testGoroutines5()
}

/*
func testGoroutines1() {
	var ch chan int
	for i := 0; i < 3; i++ {
		func(idx int) {
			ch <- (idx + 1) * 2
		}(i)
	}
	fmt.Println("result:", <-ch)
}*/

/*
	func testGoroutines2() {
		ch := make(chan string)
		go func() {
			for m := range ch {
				fmt.Println("processed:", m)
			}
		}()
		ch <- "cmd.1"
		ch <- "cmd.2"
		close(ch)
	}
*/
func testGoroutines3() {
	data := []string{"one", "two", "three"}
	group := sync.WaitGroup{}
	for _, v := range data {
		group.Add(1)
		v := v
		go func() {
			defer group.Done()
			fmt.Println(v)
		}()
	}
	group.Wait()
	//time.Sleep(3 * time.Second)
}

func testGoroutines4() {
	var num int64
	group := sync.WaitGroup{}
	for i := 0; i < 10000; i++ {
		group.Add(1)
		go func(i int64) {
			defer group.Done()
			atomic.SwapInt64(&num, i)
			//num = i
		}(int64(i))
	}

	group.Wait()
	fmt.Printf("NUM is %d", atomic.LoadInt64(&num))
}

func testGoroutines5() {
	dataMap := make(map[string]int, 10000)
	// var d map[string]int
	mu := sync.Mutex{}
	group := errgroup.Group{}
	for i := 0; i < 10000; i++ {
		i := i

		group.Go(func() error {
			mu.Lock()
			defer mu.Unlock()
			dataMap[fmt.Sprintf("%d", i)] = i

			return nil
		})

	}

	_ = group.Wait()
	//time.Sleep(5 * time.Second)

	fmt.Println(len(dataMap))
}
