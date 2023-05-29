package main

import (
	"fmt"
)

// type switch удобно выбираем по типу переменной
type customInt int

func tSw(in any) {
	////	val, ok := in.(string)
	//	fmt.Println(val, ok)

	switch val := in.(type) {
	case string:
		fmt.Println("its string!", val)
	case int:
		fmt.Println("its int!", val)
	default:
		panic("unsupported type")
	}
}

func main() {
	tSw("string")
	tSw(123)

	// custom type, bool throws panic
	tSw(customInt(1))
	//	tSw(false)

}
