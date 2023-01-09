package main

import "fmt"

func main() {
	var i int = 0
	for i = 0; i < 100; i++ {
		if i%2 == 1 {
			fmt.Println(i)
		}
	}
}
