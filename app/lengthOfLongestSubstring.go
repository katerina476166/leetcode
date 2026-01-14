package main

import (
	"unicode/utf8"
)

func lengthOfLongestSubstring(s string) int {

	l := 0
	maxl := l
	uniqueHashSet := make(map[rune]bool)
	leftIndex := 0
	for _, value := range s {
		if isrepeated, ok := uniqueHashSet[value]; ok && isrepeated {

			for runeValue, width := utf8.DecodeRuneInString(s[leftIndex:]); runeValue != value; {
				leftIndex += width
				uniqueHashSet[runeValue] = false
				l--
				runeValue, width = utf8.DecodeRuneInString(s[leftIndex:])
			}
			leftIndex += utf8.RuneLen(value)
		} else {
			l += 1
			uniqueHashSet[value] = true
		}
		if maxl < l {
			maxl = l
		}

	}
	return maxl

	/*	for index := 0; index < len(nums); index++ {
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
		return maxsum*/

}
