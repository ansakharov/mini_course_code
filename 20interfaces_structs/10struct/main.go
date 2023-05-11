package main

import (
	"fmt"
	"unsafe"
)

// Order struct with fields userID, itemID, price, and status
type Order struct {
	// indentation of bools
	Premium bool    // 8
	NewFlow bool    // 8
	UserID  int64   // 8
	ItemID  int64   // 8
	Price   float64 // 8
	Status  string  // 16
	// 42
	// 48
	// 56
}

// NewOrder is a constructor function that creates a new order
func NewOrder(userID int64, itemID int64, price float64, status string) *Order {
	return &Order{
		UserID: userID,
		ItemID: itemID,
		Price:  price,
		Status: status,
	}
}

// PrintOrderDetails prints the order details
func (o *Order) PrintOrderDetails() {
	// def f1(self, args1...)

	fmt.Printf("Order details:\n")
	fmt.Printf("User ID: %d\n", o.UserID)
	fmt.Printf("Item ID: %d\n", o.ItemID)
	fmt.Printf("Price: %.2f\n", o.Price)
	fmt.Printf("Status: %s\n", o.Status)
	fmt.Println("size", unsafe.Sizeof(*o))
}

func main() {
	order := NewOrder(1, 666, 19.99, "processing")
	order.PrintOrderDetails()

	fmt.Println(order)
}
