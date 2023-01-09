package main

import "fmt"

func main() {
	var inputValue int
	fmt.Print("enter an integer: ")
	fmt.Scan(&inputValue)
	if inputValue%2 == 0 {
		fmt.Println("its even")
	} else {
		fmt.Println("its odd")
	}
}
