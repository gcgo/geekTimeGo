package main

import (
	"fmt"
	"math"
	"strconv"
)

/**
计算两个圆的面积之和并输出，保留三位小数
*/

func main() {
	var r1, r2 float64 = 3, 4
	s1 := computeS(r1)
	s2 := computeS(r2)
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.3f", s2-s1), 64)
	fmt.Println(value)
}

/**
计算圆面积
*/
func computeS(intR float64) float64 {
	return math.Pi * intR * intR
}
