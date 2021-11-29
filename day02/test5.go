package main

import "fmt"

/**
加减乘除
*/

func main() {
	//money := 20
	//switch money {
	//case 20:
	//	fmt.Println("20")
	//	fallthrough
	//case 30:
	//	fmt.Println("30")
	//	fallthrough
	//default:
	//	fmt.Println("没钱")
	//}
	var money interface{} = 100.12
	switch newMoney := money.(type) {
	case int:
		tmpMoney := newMoney + 3.0 //3.1234报错,switch case在转类型过程中，那个type出来的值的类型会跟着case的类型自动转获取的。
		fmt.Println("int", tmpMoney)
	case float64:
		tmpMoney := newMoney + 3.0
		fmt.Println("float64", tmpMoney)
	case string:
		fmt.Println("string")
	}
}
