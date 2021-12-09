package main

import "fmt"

func main() {
	var array1 [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(array1)
	var array2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(array2)
	array3 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array3)
	array4 := [3]int{}
	fmt.Println(array4)
	for key, val := range array3 {
		fmt.Println("第", key, "个：", val)
	}
}
