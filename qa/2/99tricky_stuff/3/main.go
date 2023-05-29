package main

import (
	"fmt"
	"sync"
)

// Функция должна напечатать:
// one // two // three // (в любом порядке и в конце обязательно)
//done!
//Но это не так, исправь код

func main() {
	fmt.Println("done!")
	data := []string{"one", "two", "three"}
	printText(data)
}

func printText(data []string) {
	wg := sync.WaitGroup{}

	for _, v := range data {

		wg.Add(1)
		go func(v string) {
			fmt.Println(v)
			wg.Done()
		}(v)

	}

	wg.Wait()
}
