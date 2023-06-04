package main

import (
	"fmt"
)

type MinStack struct {
	storage  []int
	minSlice []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(in int) {
	this.storage = append(this.storage, in)

	if len(this.minSlice) == 0 {
		this.minSlice = append(this.minSlice, in)
	} else {
		if in <= this.minSlice[len(this.minSlice)-1] {
			this.minSlice = append(this.minSlice, in)
		}
	}
}

func (this *MinStack) Pop() {
	if len(this.storage) == 0 {
		return
	}

	value := this.storage[len(this.storage)-1]

	this.storage = this.storage[:len(this.storage)-1]

	if this.minSlice[len(this.minSlice)-1] == value {
		this.minSlice = this.minSlice[:len(this.minSlice)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.storage) == 0 {
		return 0
	}

	value := this.storage[len(this.storage)-1]

	return value
}

func (this *MinStack) GetMin() int {
	if len(this.minSlice) == 0 {
		return 0
	}
	return this.minSlice[len(this.minSlice)-1]
}

func main() {
	minStack := MinStack{}
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	minStack.Push(-3)
	minStack.Push(-4)
	minStack.Push(-3)
	minStack.Push(-3)
	minStack.Push(-3)
	fmt.Println(minStack.GetMin()) // return -3
	minStack.Pop()
	fmt.Println(minStack.Top())    // return 0
	fmt.Println(minStack.GetMin()) // return -2
}
