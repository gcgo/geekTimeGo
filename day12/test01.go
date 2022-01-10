package main

import "fmt"

type IOIterface interface {
	Read() string
}
type Soft struct {
}

func (Soft) Read() string {
	return "soft....."
}

func main() {
	equipment := Soft{}
	data := equipment.Read()
	fmt.Println(data)
}
