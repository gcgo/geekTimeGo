package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input string
	fmt.Scanln(&input)
	params := strings.Split(input, ",")
	switch params[1] {
	case "+":
		a, _ := strconv.Atoi(params[0])
		b, _ := strconv.Atoi(params[2])
		fmt.Println(a + b)
	case "-":
		a, _ := strconv.Atoi(params[0])
		b, _ := strconv.Atoi(params[2])
		fmt.Println(a - b)
	case "*":
		a, _ := strconv.Atoi(params[0])
		b, _ := strconv.Atoi(params[2])
		fmt.Println(a * b)
	case "/":
		a, _ := strconv.Atoi(params[0])
		b, _ := strconv.Atoi(params[2])
		fmt.Println(a / b)
	default:
		fmt.Println("不支持")

	}
}
