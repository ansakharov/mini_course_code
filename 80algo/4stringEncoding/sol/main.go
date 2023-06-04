package main

import (
	"fmt"
)

// in "AAAAABBBCCCDDDEEAA"
// out "A5B3C3D3E2A2"

func main() {
	in := "AAAAABBBCCCDDDEEAA"
	out := ""
	var curr rune
	var prev rune
	count := 0

	for idx, value := range in {
		curr = value
		if idx == 0 {
			prev = curr
			count = 1
			continue
		}

		if curr == prev {
			count++
			continue
		}

		out += fmt.Sprintf("%c%d", prev, count)
		count = 1
		prev = curr
	}

	out += fmt.Sprintf("%c%d", prev, count)

	fmt.Println(out)
}
