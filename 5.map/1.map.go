package main

import "fmt"

func main() {
	colors := map[string]int{
		"red":   12,
		"green": 20,
	}
	fmt.Println(colors)

	// O(n)
	for c, n := range colors {
		fmt.Println("c: ", c, "n: ", n)
	}

	// O(1)
	fmt.Println(colors["red"])
	colors["white"] = 10
	fmt.Println(colors)

	// O(1)
	delete(colors, "red")
	fmt.Println(colors)

	// finding O(1)
	value, ok := colors["red"]
	if ok == true {
		fmt.Println(value)
	}
	value, ok = colors["blue"]
	if ok == false {
		fmt.Println("there is not blue key")
	}

}
