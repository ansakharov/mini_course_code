package main

import (
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	ch := make(chan string, 1)

	ch2 := make(chan struct{})
	group := errgroup.Group{}

	go func() {
		for {
			select {
			case <-ch2:
				fmt.Println("timeToOut")
				break
			case ch <- "99":
				fmt.Println("written")
			}
		}
		time.Sleep(time.Second)
	}()
	//	close(ch)

	group.Go(func() error {
		k, v := <-ch
		fmt.Println(k, v)

		<-ch2
		close(ch)

		time.Sleep(time.Second * 2)
		return nil
	})

	err := group.Wait()
	fmt.Println(err)

	/*close(ch)
	k, v := <-ch
	fmt.Println(k, v)
	k, v = <-ch
	fmt.Println(k, v)

	fmt.Println("end")*/

}
