package main

import "fmt"

func indexesMap(a []int) map[int][]int {
	idxsMap := map[int][]int{}
	for i := 0; i < len(a); i++ {
		if _, ok := idxsMap[a[i]]; ok {
			idxsMap[a[i]] = append(idxsMap[a[i]], i)
		} else {
			idxsMap[a[i]] = append([]int{}, i)
		}
	}
	return idxsMap
}

func main() {
	testCases := [][]int{{1, 1, 4, 5, 8, 8, 8, 2}, {1, 1, 1, 1, 0, 0}, {}, {1}}
	for i := 0; i < len(testCases); i++ {
		fmt.Println("test: ", testCases[i])
		fmt.Println("result: ", indexesMap(testCases[i]))
	}
}
