package main

import "fmt"

/**
加减乘除
*/

func main() {
	money := 20
	switch money {
	case 20:
		fmt.Println("20")
	case 30:
		fmt.Println("30")
	default:
		fmt.Println("没钱")
	}
}
