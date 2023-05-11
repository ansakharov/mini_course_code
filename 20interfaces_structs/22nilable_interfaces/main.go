package main

import (
	"fmt"
)

type Order interface {
	GetID() int
}

type orderStruct struct {
	ID int
}

func (o *orderStruct) GetID() int {
	return o.ID
}

// receive interface return concrete types
func createOrder(shouldCreate bool) Order {
	var order *orderStruct
	fmt.Println("in createOrder: ", order == nil)

	if shouldCreate {
		order = &orderStruct{ID: 1}
	}
	return order
}

func main() {
	/*
		from runtime package

		type iface struct {
			tab  *itab
			data unsafe.Pointer
		}

		type eface struct {
			_type *_type
			data  unsafe.Pointer
		}
	*/

	order1 := createOrder(true)
	if order1 != nil {
		fmt.Printf("Order1 has ID: %d\n", order1.GetID())
	} else {
		fmt.Println("Order1 is nil")
	}

	order2 := createOrder(false)
	if order2 != nil {
		fmt.Printf("Order2 has ID: %d\n", order2.GetID())
	} else {
		fmt.Println("Order2 is nil")
	}
}
