package main

import "fmt"

func main() {
	var a []int
	fmt.Println(a)
	a = append(a, 1, 2, 3, 4, 5, 6)
	fmt.Println(a)
	c := append(a[:2], a[3:]...)
	c[0] = 9
	fmt.Println(a)

	b := "的否认不认同北方人"
	bb := []rune(b)
	bb[0] = '你'
	fmt.Println(b)

	d := make([]int, 1)
	fmt.Println(d)
}
