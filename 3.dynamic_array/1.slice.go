package main

import "fmt"

/*  complexity analysis
    1. Acces at an index -> O(1)
	1. Slicing -> O(1)
    2. Seach an element -> O(n)
    3. Insertion at the end -> O(1)
    4. Insertion at another position -> O(n)
    5. Remotion at the end -> O(1)
    6. Remotion at another position -> O(n)

*/

func main() {
	var vec []int
	for i := 1; i < 6; i++ {
		vec = append(vec, i)
	}
	fmt.Println(vec)
	// 1. Acces at an index -> O(1)
	fmt.Println(vec[3])
	fmt.Println(vec[1:3])
	// 2. Seach an element -> O(n)
	for i := 0; i < len(vec); i++ {
		if vec[i] == 4 {
			fmt.Println(i)
		}
	}
	// 3. Insertion at the end -> O(1)
	vec = append(vec, 0)
	fmt.Println(vec)
	// 4. Insertion at another position -> O(n)
	idx := 2
	vec = append(vec[:idx], append([]int{100}, vec[idx:]...)...)
	fmt.Println(vec)
	// 5. Remotion at the end -> O(1)
	vec = vec[:len(vec)-1]
	fmt.Println(vec)
	// 6. Remotion at another position -> O(n)
	ini := 2
	end := 4
	vec = append(vec[:ini], vec[end+1:]...)
	fmt.Println(vec)
}
