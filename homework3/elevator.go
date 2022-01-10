package main

import (
	"fmt"
	"sort"
	"time"
)

/**
这是一个能装无数人的电梯
*/
const (
	MAXFLOOR = 5
)

type Elevator struct {
	CurrentFloor int
	Direction    string
	Destinations []int
}

/**
电梯运行
*/
func (elevator *Elevator) LetsGo() []int {
	var tmpUp []int
	var tmpDown []int
	var tmpPath []int
	for _, destination := range elevator.Destinations {
		if destination < elevator.CurrentFloor {
			tmpDown = append(tmpDown, destination)
		}
		if destination > elevator.CurrentFloor {
			tmpUp = append(tmpUp, destination)
		}
	}
	sort.Ints(tmpUp)
	sort.Ints(tmpDown)
	//fmt.Println(tmpUp)
	//电梯运行
	run, cur := 2, 0
	for run > 0 {
		if elevator.Direction == "down" {
			for i := len(tmpDown) - 1; i >= 0; i-- {
				cur = tmpDown[i]
				if cur == elevator.CurrentFloor {
					continue
				}
				time.Sleep(time.Second)
				tmpPath = append(tmpPath, cur)
				elevator.CurrentFloor = cur
				fmt.Printf("电梯朝【%s】行进，到达【%d】层\n", elevator.Direction, elevator.CurrentFloor)
			}
			elevator.Direction = "up"
			run--
		} else if elevator.Direction == "up" {
			for i := 0; i < len(tmpUp); i++ {
				cur = tmpUp[i]
				if cur == elevator.CurrentFloor {
					continue
				}
				time.Sleep(time.Second)
				tmpPath = append(tmpPath, cur)
				elevator.CurrentFloor = cur
				fmt.Printf("电梯朝【%s】行进，到达【%d】层\n", elevator.Direction, elevator.CurrentFloor)
			}
			elevator.Direction = "down"
			run--
		}
	}
	if len(elevator.Destinations) == 0 {
		fmt.Println("没人用电梯，电梯没有动")
	}

	return tmpPath
}
