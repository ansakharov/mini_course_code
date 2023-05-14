package main

import (
	"fmt"
)

func main() {
	str := "ABCDEF"
	// str[0] = "B"
	fmt.Println(&str)

	str += "GH"

	fmt.Println(&str)
	fmt.Println(len(inStr(str)))
}

func inStr(in string) string {
	in += "sdsfsdfdsdsfsdf"
	return in
}
