package main

import (
	"fmt"
	"sort"
)

func main() {
	l := []int{10, 28, 9, 2, 1, 0, 11, 12, 9, 2}
	fmt.Println(l)
	sort.Slice(l, func(i, j int) bool { return l[i] < l[j] })
	fmt.Println(l)
}
