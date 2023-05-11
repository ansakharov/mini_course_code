package main

import (
	"fmt"
	"unsafe"
)

// interface and two structs that implement it
type Notifier interface {
	Notify() string
}

type Order struct {
	UserID int64
	ItemID int64
	Price  float64
	Status string
}

func (o Order) Notify() string {
	return fmt.Sprintf("Order with ID %d for user %d has status: %s", o.ItemID, o.UserID, o.Status)
}

type Order2 struct {
	UserID      int64
	ItemID      int64
	Price       float64
	Status      string
	customField bool
}

func (o Order2) Notify() string {
	return fmt.Sprintf("hahahaha")
}

func PrintNotification(n Notifier) {
	fmt.Println(n.Notify())
}

func ReceiveInterface(order Notifier) {
	// without ok can panic
	//res := order.(Order)

	res, ok := order.(*Order) //problem

	fmt.Println(ok)
	fmt.Println(res.UserID)
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
	var empty any
	fmt.Printf("size of empty interface is %d bytes\n\n", unsafe.Sizeof(empty))
	// struct {}{}
	order := Order{
		UserID: 42,
		ItemID: 23,
		Price:  19.99,
		Status: "processing",
	}

	//PrintNotification(order)

	ReceiveInterface(&order)

	order2 := Order2{
		UserID: 42,
		ItemID: 23,
		Price:  19.99,
		Status: "processing",
	}

	//PrintNotification(order2)

	ReceiveInterface(&order2)

	// res := order.(Order2)
	// _ = res
}
