package main

import (
	"fmt"
)

func main() {
	taskList := []string{"Проснуться", "Покушать", "Поработать"}

	wakeup := taskList[0:2]                       // Какой len/cap
	work := taskList[2:3]                         // Какой len/cap
	fmt.Println(wakeup, len(wakeup), cap(wakeup)) // "Проснуться", "Покушать", l2, c3
	fmt.Println(work, len(work), cap(work))       //  Поработать, l1, c1

	wakeup = append(wakeup, "Погулять с собакой")
	fmt.Println("TaskList: ", taskList)   //"Проснуться", "Покушать","Погулять с собакой"
	fmt.Println("Wakeup staff: ", wakeup) //"Проснуться", "Покушать","Погулять с собакой"
	fmt.Println("Workstaff:", work)       // "Погулять с собакой"
}
