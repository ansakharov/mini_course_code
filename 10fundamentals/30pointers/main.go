package main

import (
	"fmt"
	"unsafe"
)

func main() {
	in := []int{64, 8, 32, 16}

	out := make([]*int, len(in))

	// idx and value reused over all iterations.
	for idx, value := range in {
		// value := value
		out[idx] = &value
	}

	for idx, value := range out {
		fmt.Printf(
			"idx: %d, value %d, size of idx: %d, size of value: %d\n",
			idx,
			*value,
			unsafe.Sizeof(value),
			unsafe.Sizeof(value),
		)
	}
}
