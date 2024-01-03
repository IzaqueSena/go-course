package main

import "fmt"

func digitsNumber(n int) int {
	var d int
	for d = 0; n > 0; d++ {
		n /= 10
	}
	return d
}

func main() {
	testCases := []int{123456789, 100, 1, 10021}
	for i := 0; i < len(testCases); i++ {
		fmt.Print(testCases[i], " ")
		fmt.Println(digitsNumber(testCases[i]))
	}
}
