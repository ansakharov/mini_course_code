package main

import (
	"fmt"
	"sort"
)

type User struct {
	Son  *User
	Sl   []int
	ID   int
	Name string
}

func main() {
	sl := make([]User, 3)
	sl[0].ID = 1
	sl[1].ID = 2
	sl[2].ID = 90

	sl[0].Name = "Bruno"
	sl[1].Name = "Bruno"
	sl[2].Name = "Alex"

	fmt.Println(sl)

	sort.Slice(sl, func(i, j int) bool {
		return sl[i].Name < sl[j].Name
	})

	fmt.Println(sl)
}
