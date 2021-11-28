package main

import "fmt"

/**
加减乘除
*/

func main() {
	var a, b int = 30, 3
	fmt.Println("a+b= ", a+b)
	fmt.Println("a-b= ", a-b)
	fmt.Println("a*b= ", a*b)
	fmt.Println("a/b= ", a/b)
	fmt.Println(exchange(a, b))
}

func exchange(a, b int) (int, int) {
	a ^= b
	b ^= a
	a ^= b
	return a, b
}
