package main

import "fmt"

//闭包
var count int

func main() {
	fmt.Println(calcSum(1, 2, 3, 4, 5, 6))
	fmt.Println(calcSum(1, 2, 3, 4, 5, 6, 7))
	fmt.Println(calcSum(1, 2, 3, 4, 5, 6, 8))
	fmt.Println(calcSum(1, 2, 3, 4, 5, 6, 9))
	fmt.Println(calcSum(1, 2, 3, 4, 5, 6, 10))
	fmt.Println(calcSum(1, 2, 3, 4, 5, 6, 11))
	fmt.Println(calcSum(1, 2, 3, 4, 5, 6, 12))

	showCount()
}

func showCount() {
	fmt.Println("sum: ", count)
}

func calcSum(nums ...int) (sum int) {
	for _, num := range nums {
		sum += num
	}
	count++
	return sum
}
