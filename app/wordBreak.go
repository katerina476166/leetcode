package main

func wordBreak(s string, wordDict []string) bool {
	return wordBreakdp(s, wordDict)
}

func wordBreakdp(s string, wordDict []string) bool {

	dp := make([]bool, len(s)+1)
	dp[0] = true
	l := len(s)

	for i := 0; i <= l; i++ {
		if !dp[i] {
			continue
		}

		for _, w := range wordDict {
			end := i + len(w)
			if end <= l && s[i:end] == w {
				dp[end] = true

			}
		}

	}
	return dp[l]
}

//ar[0..len]
//ar[i]=true if wordtrue
//
