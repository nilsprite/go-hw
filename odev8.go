package main

import "fmt"

func main() {
	var height float64
	var weight float64

	fmt.Println("weight (kg) : ")
	fmt.Scan(&weight)
	fmt.Println("height (m) : ")
	fmt.Scan(&height)
	var bmi float64 = weight / (height * height)
	//bmi= body mass index
	fmt.Println("BMI:")
	fmt.Println(bmi)

	if bmi < 18.5 {
		fmt.Println("you are skinny")
	} else if 18.5 <= bmi && bmi < 24.9 {
		fmt.Println("you are normal weight")
	} else if 24.9 <= bmi && bmi < 29.9 {
		fmt.Println("you are overweight")
	} else if 29.9 <= bmi && bmi < 34.9 {
		fmt.Println("you are first degree obese")
	} else if 34.9 <= bmi && bmi <= 39.9 {
		fmt.Println(" you are second degree obese")
	} else if 40 < bmi {
		fmt.Println("you are third degree obese")
	}

}
