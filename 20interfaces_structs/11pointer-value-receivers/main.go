package main

import (
	"fmt"
	"unsafe"
)

type MyStruct struct {
	Value int64 `json:"value"`
}

func (i *MyStruct) IncrementPointerReceiver() {
	i.Value++
	fmt.Printf("in method IncrementPointerReceiver, mem address: %p\n", &i)
	fmt.Printf("in method IncrementPointerReceiver: %v\n", i)
}

func main() {
	myS := MyStruct{Value: 665}

	fmt.Printf("MyStruct size: %d\n", unsafe.Sizeof(MyStruct{}))
	fmt.Printf("Empty struct size: %d\n", unsafe.Sizeof(struct{}{}))
	fmt.Println()

	myS.IncrementPointerReceiver()

	fmt.Println()
	fmt.Printf("in main current state: %v\n", myS) // 666
	fmt.Printf("in main, mem address: %p\n", &myS)

	fmt.Println()
}
