package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	var intVar int
	var int8Var int8
	var int16Var int16
	var int32Var int32
	var int64Var int64
	/*
		int8        the set of all signed  8-bit integers (-128 to 127)
		int16       the set of all signed 16-bit integers (-32768 to 32767)
		int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
		int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
	*/
	fmt.Printf("Size of int:   %d bytes\n", unsafe.Sizeof(intVar))
	fmt.Printf("Size of int8:  %d bytes\n", unsafe.Sizeof(int8Var))
	fmt.Printf("Size of int16: %d bytes\n", unsafe.Sizeof(int16Var))
	fmt.Printf("Size of int32: %d bytes\n", unsafe.Sizeof(int32Var))
	fmt.Printf("Size of int64: %d bytes\n", unsafe.Sizeof(int64Var))
	fmt.Println()

	var uintVar uint
	var uint8Var uint8
	var uint16Var uint16
	var uint32Var uint32
	var uint64Var uint64
	/*
		uint8       the set of all unsigned  8-bit integers (0 to 255)
		uint16      the set of all unsigned 16-bit integers (0 to 65,535)
		uint32      the set of all unsigned 32-bit integers (0 to 4,294,967,295)
		uint64      the set of all unsigned 64-bit integers (0 to 18,446,744,073,709,551,615)
	*/
	fmt.Printf("Size of uint:   %d bytes\n", unsafe.Sizeof(uintVar))
	fmt.Printf("Size of uint8:  %d bytes\n", unsafe.Sizeof(uint8Var))
	fmt.Printf("Size of uint16: %d bytes\n", unsafe.Sizeof(uint16Var))
	fmt.Printf("Size of uint32: %d bytes\n", unsafe.Sizeof(uint32Var))
	fmt.Printf("Size of uint64: %d bytes\n", unsafe.Sizeof(uint64Var))
	fmt.Println()

	var float32Var float32
	var float64Var float64
	/*
		float32     the set of all IEEE-754 32-bit floating-point numbers (approx. ±1.18e-38 to ±3.4e+38 with 7 decimal precision)
		float64     the set of all IEEE-754 64-bit floating-point numbers (approx. ±2.23e-308 to ±1.80e+308 with 15 decimal precision)
	*/

	var ordinaryNumber float32 = 1e2
	var largeNumber float32 = 1e7
	var smallNumber float32 = 0.000123
	var sumOrdinary float32 = ordinaryNumber + smallNumber
	var sumFloat32 = largeNumber + smallNumber

	var ordinaryNumber64 float64 = 1e2
	var largeNumber64 float64 = 1e7
	var smallNumber64 float64 = 0.000123
	var sumOrdinar64 = ordinaryNumber64 + smallNumber64
	var sumFloat64 = largeNumber64 + smallNumber64

	fmt.Printf("Float32: %.36f\n", sumOrdinary)
	fmt.Printf("Float64: %.36f\n", sumOrdinar64)

	fmt.Printf("Float32: %.36f\n", sumFloat32)
	fmt.Printf("Float64: %.36f\n", sumFloat64)

	fmt.Printf("Size of float32: %d bytes\n", unsafe.Sizeof(float32Var))
	fmt.Printf("Size of float64: %d bytes\n", unsafe.Sizeof(float64Var))
	fmt.Println()

	// Overflow of int64
	int64Var = math.MaxInt64 // 9223372036854775807 max value for int64
	fmt.Println("Max int64 value:", int64Var)

	int64Var++ // Incrementing the value causes an overflow
	fmt.Println("Overflowed int64 value:", int64Var)
}
