package main

import "fmt"

func main() {
	var midterm float32
	var final float32
	fmt.Println("midterm: ")
	fmt.Scan(&midterm)
	fmt.Println("final: ")
	fmt.Scan(&final)

	var weightOfMidterm float32 = midterm * 30 / 100
	var weightOfFinal float32 = final * 70 / 100
	var average float32 = (weightOfMidterm + weightOfFinal) / 2
	fmt.Println("average: ")
	fmt.Println(average)

}
