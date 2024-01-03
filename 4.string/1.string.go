package main

import (
	"fmt"
	"strings"
)

/*  complexity analysis
    1. Acces at an index -> O(1)
	1. Slicing -> O(1)
    2. Search an substring -> O(n)
    3. Insertion at the end -> O(m) m is the size of the other string if m<<n => O(1)
    4. Insertion at another position -> O(n)
    5. Remotion at the end == slicing -> O(1)
    6. Remotion at another position -> O(n - idx)
    7. Replacing -> O(n) n is the length of new string
    8. Slicing -> O(n) n is the length o substring
*/

func main() {
	var str string
	str = "abcdefghij"
	fmt.Println(str)
	// 1. Acces at an index -> O(1)
	fmt.Println("---1---")
	fmt.Println(fmt.Sprintf("%c", str[2]))
	fmt.Println(string(str[2]))
	fmt.Println(str[2 : 2+1])
	fmt.Println(str[2:5])
	// 2. Search an element -> O(n)
	fmt.Println("---2---")
	if strings.Index(str, "def") != -1 {
		fmt.Printf("def is in %d\n", strings.Index(str, "def"))
	}
	if strings.Index(str, str[2:4]) != -1 {
		fmt.Printf("%d\n", strings.Index(str, str[2:4]))
	}
	// 3. Insertion at the end -> O(m) (another string), if m << n => O(1)
	fmt.Println("---3---")
	str += "AAA"
	fmt.Println(str)
	// 4. Insertion at another position -> O(m + n - ini), if m << n => O(n-ini)
	fmt.Println("---4---")
	idx := 4
	str = str[:idx] + "BBB" + str[idx:]
	fmt.Println(str)
	// 5. Remotion at the end == slicing -> O(1)
	fmt.Println("---5---")
	str = str[:len(str)-1]
	fmt.Println(str)
	// 6. Remotion at another position -> O(n-idx)
	fmt.Println("---6---")
	str = str[:3] + str[5:]
	fmt.Println(str)
	// 7. Replacing -> O(n)
	fmt.Println("---7---")
	str = str[:4] + "NNNN" + str[4+4:]
	fmt.Println(str)
}
