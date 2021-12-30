package main

import "fmt"

func main() {
	i := 1
	for {
		fmt.Println("计数器: ", i)
		if i >= 10 {
			break
		}
		i++
	}
	i = 0
	for i < 10 {
		fmt.Println("计数器: ", i)
		i++
	}
}
