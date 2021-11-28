package main

import (
	"fmt"
	"math"
	"strconv"
)

/**
测试变量与常量
*/

var v2 = 111

const pi = math.Pi

func main() {
	var f1 float64 = 1.234
	var i1 int = int(f1)
	fmt.Println("f1:", f1, "i1:", i1)
	var a6 uint = math.MaxUint64
	var a7 int = int(a6)
	fmt.Println(a6, a7)
	fmt.Println(v2)

	i, err := strconv.Atoi("12345")
	if err != nil {
		panic(err)
	}
	i += 3
	println(i)

	s := strconv.Itoa(12345)
	s += "3"
	println(s)

	println(pi)
}
