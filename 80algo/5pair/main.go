package main

import (
	"fmt"
)

/*
// 1 Точно ли понял задачу?
// 2 Какова сложность алгоритма? Проговори это вслух
// 3 Какие расходы по памяти?
// 4 не уходи в сторону
// 5 Имеет ли смысл реализовавывать мое решение?
// 6 !!! Проверка готового решения

Example 1:
Input: nums = [2,1,1, 6, 5, 7,11,15], target = 9
Output: [0,5]
Explanation: Because nums[0] + nums[5] == 9, we return [0, 5].

Example 2:
Input: nums = [3,2,4], target = 6
Output: [1,2]

Example 3:
Input: nums = [3,3], target = 6
Output: [0,1]
*/
func main() {
	in1 := []int{2, 1, 1, 6, 5, 7, 11, 15}

	fmt.Println(findSum(in1, 9))

	in2 := []int{3, 2, 4}

	fmt.Println(findSum(in2, 6))

	in3 := []int{3, 4, 5, 9, 3}

	fmt.Println(findSum(in3, 6))
}

func findSum(arr []int, target int) []int {
	storage := make(map[int]int, len(arr))

	for idx, value := range arr {
		idxOfPair, ok := storage[target-value]
		if !ok {
			storage[value] = idx
			continue
		}

		return []int{idxOfPair, idx}
	}

	return nil
}
