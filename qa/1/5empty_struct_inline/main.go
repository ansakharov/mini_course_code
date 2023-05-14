package main

import (
	"fmt"
	"unsafe"
)

// Почему пустая структура занимает 0 байт, может есть какие-то нюансы?

type EmptyFirst struct {
	B struct{}
	A int32
}

type EmptyLast struct {
	A int32
	B struct{}
}

func main() {
	first := EmptyFirst{}
	last := EmptyLast{}
	fmt.Println(unsafe.Sizeof(first))
	fmt.Println(unsafe.Sizeof(last))
}

func Unique(in []int64) []int64 {
	out := make([]int64, 0, len(in))
	unique := make(map[int64]struct{})

	for _, value := range in {
		if _, ok := unique[value]; !ok {
			unique[value] = struct{}{}
			out = append(out, value)
		}
	}

	return out
}
