package main

import (
	"bytes"
	"fmt"
	"io"
)

///Почему запись данных в буфер начинается с нулевого элемента,
//если создаётся слайс с указателем на четвёртый элемент?

func main() {
	s := "23458"
	sl := make([]byte, 10, 100)
	sl2 := sl[3:]

	for i := 0; i < len(sl); i++ {
		sl[i] = 1
	}

	n, err := io.ReadAtLeast(bytes.NewBuffer(([]byte(s))[3:5]), sl2, 1)

	fmt.Printf("n: %d    err: %s\n;   sl: %v\n ls->string: %s", n, err, sl, string(sl))
	fmt.Println()

	fmt.Println()
	fmt.Printf("n: %d    err: %s\n;   sl: %v\n ls->string: %s", n, err, sl2, string(sl2))
}
