package main

import "fmt"

func main() {
	var values [5]float32
	fmt.Println("enter the values to get the avarage result: ")
	for i := 0; i < 5; i++ {
		fmt.Scan(&values[i])
	}
	fmt.Println(values)
	var average = (values[0] + values[1] + values[2] + values[3] + values[4]) / 5
	fmt.Println(average)

}
