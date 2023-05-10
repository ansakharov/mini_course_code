package main

import (
	"fmt"
)

func main() {
	//var m = map[string]int{}
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
	}
	m["s"] = 1

	for i := 0; i < 5; i++ {
		for key, value := range m {
			fmt.Printf("%s: %d | ", key, value)
		}
		fmt.Println()
	}
}
