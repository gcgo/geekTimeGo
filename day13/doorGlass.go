package main

import "fmt"

var _ Door = &DoorGlass{}

type DoorGlass struct{}

func (DoorGlass) Lock() {
	fmt.Println("锁上玻璃门")
}

func (DoorGlass) Open() {
	fmt.Println("打开玻璃门")
}
func (DoorGlass) Close() {
	fmt.Println("关闭玻璃门")
}
