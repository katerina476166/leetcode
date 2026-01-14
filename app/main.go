package main

import (
	"fmt"
)

//go mod init app
//запуск в консоли go run main.go
//после любого измениния Ctrl+S, иначе собирает старое

func main() {

	fmt.Printf(" L=  %d\n", lengthOfLongestSubstring("pwwkew"))
	fmt.Printf(" L=  %d\n", lengthOfLongestSubstring("bbbbb"))
	fmt.Printf(" L=  %d\n", lengthOfLongestSubstring("abcabcbb"))

}
