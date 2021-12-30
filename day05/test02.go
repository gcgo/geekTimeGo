package main

import "fmt"

func main() {
	fmt.Print("身高（米）：")
	var weight float64
	fmt.Scanln(&weight)
	var height float64
	fmt.Scanln(&height)
	var age int
	fmt.Scanln(&age)
	var sex string
	fmt.Scanln(&sex)
	var bmi float64 = weight / (height * height)
	var male int = 1
	if sex == "女" {
		male = 0
	}
	var test float64 = 1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(male)
	fmt.Println("体脂率：", test)

}
