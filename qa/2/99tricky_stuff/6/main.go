package main

import (
	"fmt"
	"sync"
	"time"
)

//Пример для разбора
//6.
//Опиши, что делает функция isCallAllowed?

var callCount = make(map[uint]uint)
var locker = &sync.Mutex{}

func isCallAllowed(allowedCount uint) bool {
	if allowedCount == 0 {
		return true
	}

	locker.Lock()
	defer locker.Unlock()
	// 21:10:00-21:10:30, 21:10:30-21:11:00, 21:11:01
	// 1322093490 / 30 -> 5585, 5586, 5587
	curTimeIndex := uint(time.Now().Unix() / 30)

	// предыдущие 30 секунд
	prevIndexVal, _ := callCount[curTimeIndex-1]

	// предыдущий бакет съел весь лимит
	if prevIndexVal >= allowedCount {
		return false
	}

	curIndexVal, ok := callCount[curTimeIndex]
	if !ok {
		callCount[curTimeIndex] = 1
		return true
	}
	if (curIndexVal + prevIndexVal) >= allowedCount {
		return false
	}
	callCount[curTimeIndex]++
	return true
}

func main() {
	fmt.Printf("%v\n", isCallAllowed(3)) // true
	fmt.Printf("%v\n", isCallAllowed(3)) // true
	fmt.Printf("%v\n", isCallAllowed(3)) // true
	time.Sleep(time.Second * 30)
	fmt.Printf("%v\n", isCallAllowed(3)) // false
	fmt.Printf("%v\n", isCallAllowed(3)) // false
}
