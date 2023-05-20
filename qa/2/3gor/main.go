package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
Добрый день! на первой Q&A сессии задавал вопрос вопрос насчет утечек при ре-слайсинге массива, договорились,
что подробнее распишу в форме. Вопрос возник исходя из https://github.com/golang/go/wiki/SliceTricks:
"NOTE If the type of the element is a pointer or a struct with pointer fields, which need to be garbage collected,
the above implementations of Cut and Delete have a potential memory leak problem: some elements with values are still
referenced by slice a and thus can not be collected."
Не совсем понятно реально ли будет утечка, и почему это происходит только с указателями, а не с обычными знаениями?
Есть еще подобный вопросик на stackoverflow - https://stackoverflow.com/questions/55045402/memory-leak-in-golang-slice.
Русская версия если лень переводить) - https://question-it.com/questions/15884449/utechka-pamjati-v-sreze-golanga
*/
type forTest struct {
	in int64
}

func main() {
	var mem runtime.MemStats

	runtime.ReadMemStats(&mem)
	fmt.Printf("Before get(): Alloc = %v MiB\n", mem.Alloc/1024/1024)

	resOfRes := make([][]*forTest, 0, 100)

	for i := 0; i < 50; i++ {
		res := getLast()

		runtime.GC()
		runtime.ReadMemStats(&mem)
		fmt.Printf("After getAll(): Alloc = %v MiB, slice: %v\n", mem.Alloc/1024/1024, res)

		resOfRes = append(resOfRes, res)
	}

	fmt.Println()

	_ = resOfRes

	// Now go into loop and see that pointers are alive
	for i := 0; i < 20000; i++ {
		runtime.GC()
		runtime.ReadMemStats(&mem)
		fmt.Printf(
			"GC can't clear underlaying array, because slice points on it elem\nAlloc = %v MiB\n\n",
			mem.Alloc/1024/1024,
		)
		time.Sleep(time.Second * 2)
	}

	// remove this print and then GC can remove all slices way before
	//	fmt.Println(resOfRes)
}

func getLast() []*forTest {
	sl := make([]*forTest, 0, 100000)

	for i := 0; i < 100000; i++ {
		sl = append(sl, &forTest{in: int64(i)})
	}

	// return last element
	return sl[99999:]

	/*
		Instead copy elem to new slice and old array will be removed by GC eventually.
		lastElem := make([]*forTest, 1)
		lastElem[0] = sl[99999]

		return lastElem
	*/

}
