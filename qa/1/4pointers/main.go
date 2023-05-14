package main

import (
	"fmt"
	"unsafe"
)

/*
Можно подробнее пояснить про работу со слайсом поинтеров на int?
не получилось пока понять пример с выводом в цикле значений из слайса
*/

/*
на 16:13 вы упоминаете популярный пример ошибки с указателями,
и который раз упоминаете, что на php по другому! да так же на php! просто мы там такое редко делаем.
вот ссылка на такой же код https://3v4l.org/bTkQJ - полный эквивалент и вывод такой же - все 16!!!
*/
func main() {
	in := []int{64, 8, 32, 16}

	out := make([]*int, len(in))

	// idx and value reused over all iterations.
	for idx, value := range in {
		value := value
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
