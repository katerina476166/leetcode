package main

//если впервые, то
// добавляем в словарь, прибавляем к сумме
//если повтор, то
//удаляем то, что было раньше, считаем сумму удаленных и вычитаем их из суммы

func maximumUniqueSubarray(nums []int) int {
	sum := 0
	maxsum := sum
	uniqueHashSet := make(map[int]bool)
	left := 0

	for index := 0; index < len(nums); index++ {
		value := nums[index]

		if repeatedInfo, ok := uniqueHashSet[value]; ok && repeatedInfo {
			subtractionSum := 0
			for nums[left] != value {
				uniqueHashSet[nums[left]] = false
				subtractionSum += nums[left]
				left++
			}
			left++

			sum -= subtractionSum
		} else {
			sum += value
		}

		uniqueHashSet[value] = true

		if maxsum < sum {
			maxsum = sum
		}
	}
	return maxsum
}
