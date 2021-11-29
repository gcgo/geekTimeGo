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
	var money interface{} = 100
	switch newMoney := money.(type) {
	case int:
		tmpMoney := newMoney + 3.0 //3.1234报错
		fmt.Println("int", tmpMoney)
	case float64:
		fmt.Println("float64")
	case string:
		fmt.Println("string")
	}
}
