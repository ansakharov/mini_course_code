package main

import (
	"fmt"
)

// AAAAABBBCCCDDDEEAA
// A5B3C3D3E2A2

// 1 Точно ли понял задачу?
// 2 Какова сложность алгоритма? Проговори это вслух
// 3 Какие расходы по памяти?
// 4 не уходи в сторону
// 5 Имеет ли смысл реализовавывать мое решение?
// 6 !!! Проверка готового решения
func main() {
	in := "AAAAABBBCCCDDDEEAA"
	fmt.Println(stringEncoding(in))
}

func stringEncoding(in string) string {
	var prev rune
	count := 1
	out := ""

	for i := 0; i < len(in); i++ {
		_ = in[i]
	}
	for idx, singleRune := range in {
		if idx == 0 {
			prev = singleRune
			continue
		}

		if singleRune == prev {
			count++
			continue
		}

		out += fmt.Sprintf("%c%d", prev, count)
		prev = singleRune
		count = 1
	}

	out += fmt.Sprintf("%c%d", prev, count)

	return out
}
