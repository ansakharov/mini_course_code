package main

import "fmt"

func main() {
	wrong()
	correct()
}

func wrong() {
	m := map[string]int{
		"one": 1,
		"two": 2,
		//	"three": 0,
	}

	value := m["three"]
	fmt.Println("value: ", value)
}

func correct() {
	m := map[string]int{
		"one": 1,
		"two": 2,
	}

	value, ok := m["three"]
	if !ok {
		fmt.Println("value: !ok")

		return
	}
	fmt.Println("value: ", value)
}
