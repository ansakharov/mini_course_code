package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	mutex1 := &sync.Mutex{}
	mutex2 := &sync.Mutex{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		mutex1.Lock()
		defer mutex1.Unlock()

		fmt.Println("g1 взяла первый мьютекс")

		mutex2.Lock()
		defer mutex2.Unlock()

		fmt.Println("g1  взяла второй мьютекс")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		time.Sleep(0 * time.Second * time.Duration(rand.Intn(2)))
		mutex2.Lock()
		defer mutex2.Unlock()

		fmt.Println("g2 взяла второй мьютекс")

		mutex1.Lock()
		defer mutex1.Unlock()

		fmt.Println("g2 взяла первый мьютекс")
	}()

	wg.Wait()

	fmt.Println("Дедлока нет!")
}
