package main

import "fmt"

func freqMap(a []int) map[int]int {
	freq := map[int]int{}
	for i := 0; i < len(a); i++ {
		freq[a[i]]++
	}
	return freq
}

func main() {
	testCases := [][]int{{1, 1, 4, 5, 8, 8, 8, 2}, {1, 1, 1, 1, 0, 0}, {}, {1}, {0}}
	for i := 0; i < len(testCases); i++ {
		fmt.Println("test: ", testCases[i])
		freqs := freqMap(testCases[i])
		fmt.Println(freqs)
	}
}
