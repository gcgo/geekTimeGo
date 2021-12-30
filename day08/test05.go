package main

import (
	"fmt"
	_ "github.com/spf13/cobra"
	"os"
)

func main() {
	//录入
	var name string
	var sex, weight, height, age string
	//var next int
	args := os.Args
	fmt.Println(args)
	name = args[1]
	sex = args[2]
	height = args[3]
	weight = args[4]
	age = args[4]
	//计算

	//评估结果

}
