package main

import (
	"fmt"
	"io"
)

type OrderType []int

func (O OrderType) GetFirst() int {
	return O[0]
}

// Познание нераспространённых трюков языка Go - например,
// как заставить функцию имплементировать интерфейс, без пустой структуры. О чем речь?)
// Можно какой то пример?

type FuncThatWillBeWriterAndReader func(p []byte) (n int, err error)

func (f FuncThatWillBeWriterAndReader) Read(p []byte) (n int, err error) {
	return f(p)
}

func (f FuncThatWillBeWriterAndReader) Write(p []byte) (n int, err error) {
	return f(p)
}

func main() {
	f := FuncThatWillBeWriterAndReader(func(p []byte) (n int, err error) {
		copy(p, "Hahahaha!")
		return len(p), io.EOF
	})

	var r io.Reader = f

	buf := make([]byte, 1024)

	fmt.Println()
	fmt.Println("Calling the reader")

	n, err := r.Read(buf)
	fmt.Println(string(buf[:n]), err)

	/*	f = func(p []byte) (n int, err error) {
		copy(p, "Hahahaha i write here!")

		return len(p), nil
	}*/
	var wr io.Writer = f

	newBuf := make([]byte, 1024)

	fmt.Println()
	fmt.Println("Calling the writer")
	nWr, errWr := wr.Write(newBuf)
	fmt.Println(string(newBuf[:nWr]), errWr)

}
