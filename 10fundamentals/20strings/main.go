package main

import (
	"fmt"
	"unsafe"
)

func main() {
	/*
		type StringHeader struct {
			Data uintptr
			Len  int
		}
	*/
	// Part 1: Iterating Over Bytes in a String
	str := "abcdefqweqwqweqweqwewqe"

	fmt.Println("Part 1: Iterating Over Bytes in a String")
	for i := 0; i < len(str); i++ {
		fmt.Printf("Character: %c, Byte: %x\n", str[i], str[i])
	}
	fmt.Println()

	// Part 2: Iterating Over Bytes in a Multibyte String
	strMultiBytes := "абвгде"

	fmt.Println("Part 2: Iterating Over Bytes in a Multibyte String")
	for i := 0; i < len(strMultiBytes); i++ {
		fmt.Printf("Byte at offset %d: %x, trying print: %c\n", i, strMultiBytes[i], strMultiBytes[i])
	}
	fmt.Println()

	// Part 3: Iterating Over Runes in a Multibyte String
	fmt.Println("Part 3: Iterating Over Runes in a Multibyte String")
	for i, r := range strMultiBytes {
		fmt.Printf("Character: %c, Rune: %U, Byte offset: %d, Size: %d\n", r, r, i, unsafe.Sizeof(r))
	}

	// error
	// strMultiBytes[0] = "i"

	// Part 4: Convert []byte to string with unsafe without allocations
	bytes := []byte{'a', 'b', 'c', 'd', 'e', 'f'}
	strFromBytes := *(*string)(unsafe.Pointer(&bytes))
	bytes[5] = 'a'
	bytes[4] = 'b'
	bytes[3] = 'c'
	bytes[2] = 'd'
	bytes[1] = 'e'
	bytes[0] = 'f'

	fmt.Printf("String from []byte without allocations: %s\n", strFromBytes)
	bytes = []byte{'a', 'a', 'a', 'a', 'a', 'a'}
	fmt.Printf("Other storages for string and []byte: %s\n", strFromBytes)
}
