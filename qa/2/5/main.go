package main

import "fmt"

func main() {
	//testSlices1()
	//testSlices2()
	//testSlices3()
	testSlices4()
}

func testSlices1() {
	a := []string{"a", "b", "c", "d", "e"} // l:5 c:5
	b := a[1:2]                            // l:? c:?
	// "b" - len 1, cap 4
	b[0] = "q"
	fmt.Println(a, len(a), cap(a)) // ? "a", "q", "c", "d", "e"
	fmt.Println(b, len(b), cap(b)) // ? "q"
}

func testSlices2() {
	a := []byte{'a', 'b', 'c'}
	// a[1:2] len 1, cap 2
	b := append(a[1:2], 'd') // 'b', 'd'
	b[0] = 'z'               // 'z', 'd'
	fmt.Printf("%s\n", a)    // azd
	fmt.Printf("%s\n", b)    // zd
}

func testSlices3() {
	a := []byte{'a', 'b', 'c'} // abd
	// a[1:2] len 1, cap 2
	b := append(a[1:2], 'd', 'x') // bdx
	b[0] = 'z'                    // zdx
	fmt.Printf("%s\n", a)         // abd
	fmt.Printf("%s\n", b)         // zdx
}

func testSlices4() {
	a := []byte{'a', 'b', 'c'}
	b := string(a)
	a[0] = 'z'
	fmt.Printf("%s\n", b)
}
