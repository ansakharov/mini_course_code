package main

import "fmt"

/*
вопрос ещё есть слайсам - почему мы слайс резервируем - но обращение к этим ячейкам вызывает панику ?
в чём был смысл так делать ? ведь ведь заразервированная память всё равно не доступна другим.
или мы экономим процессорное время до тех пор, пока не произойдёт аллокация?
а по сути экономии памяти - нет ?
*/
func main() {
	s1 := make([]int, 0, 1024)

	fmt.Printf("s1: len = %d, cap = %d\n", len(s1), cap(s1))

	// panic
	//fmt.Println(s1[6])

	s2 := make([]int, 5)
	fmt.Printf("s2: len = %d, cap = %d\n", len(s2), cap(s2))

	// also panic
	// fmt.Println(s2[6_write_in_slice])

	// overflow cap -> len now 2x
	s2 = append(s2, 10)

	fmt.Printf("s2 new len-cap: len = %d, cap = %d\n", len(s2), cap(s2))

	// it's ok now
	fmt.Println(s2[5])
}
