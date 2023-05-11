package main

import "fmt"

type S struct {
	value string
}

func (s S) String() string {
	return fmt.Sprintf("my name is %s", s.value)
}

func main() {
	// def f(self,
	s := S{
		value: "hello",
	}

	sPtr := &S{
		value: "hello2",
	}

	fmt.Println(s, "|", sPtr)
}
