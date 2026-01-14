package main

import (
	"sort"
)

///algo
//1.sort word by len
//2.find words concatenated and mark it
//conc[bool]
//завести массив макс длины и переиспольоватт его [maxlen] int dp
// его при использовании иниц -1
//в кадом индексе запоминаем id слова
//если последний элт заоплнен, то восстанавливаем из массива путь и отмечаем в concat

func findAllConcatenatedWordsInADict(words []string) []string {
	sort.Slice(words, func(i, j int) bool {
		// This function defines the sorting logic.
		// It should return true if the element at index i should come before the element at index j.
		return len(words[i]) < len(words[j])
	})

	dp := make([]int, len(words[len(words)-1])+1)
	conc := make([]bool, len(words))

	for index, w := range words {
		fillDP(w, words, dp, index)
		if dp[len(w)] != 0 {
			//	findpath(w, words, dp, conc)
			conc[index] = true
		}

		for j := len(dp) - 1; j >= 0; j-- {
			dp[j] = 0
		}
	}

	var mySlice []string

	for j := len(conc) - 1; j >= 0; j-- {
		if conc[j] {
			mySlice = append(mySlice, words[j])
		}
	}

	return mySlice
}

func findpath(s string, wordDict []string, dp []int, conc []bool) {

	for j := len(s); j > 0; {
		index := dp[j] - 1
		word := wordDict[index]
		j = j - len(word)
		conc[index] = true
	}
}

func fillDP(s string, wordDict []string, dp []int, exc int) {
	l := len(s)

	for i := 0; i <= l; i++ {
		if dp[i] == 0 && i != 0 {
			continue
		}

		//
		for index := 0; index < exc; index++ {
			w := wordDict[index]
			end := i + len(w)
			if end <= l && s[i:end] == w {
				dp[end] = index + 1
			}
		}
	}
}

//ar[0..len]
//ar[i]=true if wordtrue
//
