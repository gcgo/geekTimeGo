package main

import "fmt"

func main() {
	a, b := 1, 2
	c := add(&a, &b)
	fmt.Println(c)
}

func add(a, b *int) (res int) {
	return *a + *b

}
