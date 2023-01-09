package main

import "fmt"

func main() {
	var values [5]int
	fmt.Println("enter 5 random integers")
	for i := 0; i < 5; i++ {
		fmt.Scan(&values[i])
	}
	fmt.Println(values)
	var min int = values[0]
	for i := 1; i < 5; i++ {
		if values[i] < min {
			min = values[i]
		}
	}
	fmt.Println(min)
}
