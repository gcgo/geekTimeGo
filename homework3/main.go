package main

import (
	"fmt"
	"time"
)

/**
【说明】：
	模拟电梯一次运行的结果
【使用】：
	当被询问"有人坐电梯吗？"时，可以选则有或者没有；
	选择有，输入要去的楼层；
	所有人输入完毕后，电梯运行。
*/
func main() {
	e := &Elevator{
		CurrentFloor: 3,    //电梯初始位置，在3层
		Direction:    "up", //默认向上行驶，向下为：down
	}
	floor := 0
	next := "y"
	for {
		fmt.Println("有人坐电梯吗？\t有：【y】\t没有:【任意输入】")
		fmt.Scanln(&next)
		if next != "y" {
			break
		}
		fmt.Println("请输入你要去几层：【1-5】")
		fmt.Scanln(&floor)
		if floor <= 0 || floor > MAXFLOOR {
			fmt.Println("没有这个楼层")
			continue
		}
		e.Destinations = append(e.Destinations, floor)

	}
	//电梯运行
	path := e.LetsGo()
	time.Sleep(time.Second)
	fmt.Println("电梯运行轨迹：", path)
}
