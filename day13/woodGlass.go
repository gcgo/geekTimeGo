package main

import "fmt"

type WoodGlass struct{}

func (*WoodGlass) Open() {
	fmt.Println("打开木门门")
}
func (*WoodGlass) Close() {
	fmt.Println("关闭木门")
}
