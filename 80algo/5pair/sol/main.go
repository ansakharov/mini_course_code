package sol

func twoSum(nums []int, target int) []int {
	mapOfVals := make(map[int]int, len(nums))

	for idx, val := range nums {
		pair := target - val
		idxFirst, ok := mapOfVals[pair]
		if ok {
			return []int{idxFirst, idx}
		}

		mapOfVals[val] = idx
	}

	return nil
}

func findSum(arr []int, target int) []int {
	storage := make(map[int]int, len(arr))

	for idx, value := range arr {
		pairIdx, ok := storage[target-value]
		if !ok {
			storage[value] = idx
			continue
		}
		return []int{pairIdx, idx}
	}

	return nil
}
