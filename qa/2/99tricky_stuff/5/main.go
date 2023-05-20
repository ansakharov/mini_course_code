package main

import (
	"fmt"
	"time"
)

// Мы пытаемся подсчитать количество выполненных параллельно операций,
//что может пойти не так?

var callCounter uint

func main() {
	for i := 0; i < 10000; i++ {
		go func() {
			// Ходим в базу, делаем долгую работу
			time.Sleep(time.Second)
			//Увеличиваем счетчик
			callCounter++
		}()
	}
	fmt.Println("Call counter value = ", callCounter)
}
