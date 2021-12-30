package main

import "fmt"

//闭包

func main() {
	testFunc := calcSum1()
	fmt.Println(testFunc(1, 2, 3, 4, 5))
	fmt.Println(testFunc(1, 2, 3, 4, 5))
	fmt.Println(testFunc(1, 2, 3, 4, 5))
	fmt.Println(testFunc(1, 2, 3, 4, 5))
	fmt.Println(testFunc(1, 2, 3, 4, 5))
	fmt.Println(testFunc(1, 2, 3, 4, 5))
}

func calcSum1() func(...int) (int, int) {
	var i int
	return func(nums ...int) (int, int) {
		sum := 0
		for _, num := range nums {
			sum += num
		}
		i++
		return sum, i
	}

}
