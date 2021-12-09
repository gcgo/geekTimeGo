package main

import "fmt"

func main() {
	var ax1, ay1, ax2, ay2, bx1, by1, bx2, by2 float64
	fmt.Println("请输入第一条线的第一个点的x坐标：（整数）")
	fmt.Scanln(&ax1)
	fmt.Println("请输入第一条线的第一个点的y坐标：（整数）")
	fmt.Scanln(&ay1)
	fmt.Println("请输入第一条线第二个点的x坐标：（整数）")
	fmt.Scanln(&ax2)
	fmt.Println("请输入第一条线第二个点的y坐标：（整数）")
	fmt.Scanln(&ay2)
	fmt.Println("请输入第二条线第一个点的x坐标：（整数）")
	fmt.Scanln(&bx1)
	fmt.Println("请输入第二条线第一个点的y坐标：（整数）")
	fmt.Scanln(&by1)
	fmt.Println("请输入第二条线第二个点的x坐标：（整数）")
	fmt.Scanln(&bx2)
	fmt.Println("请输入第二条线第二个点的y坐标：（整数）")
	fmt.Scanln(&by2)

	boolPingXing := pingxing(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2)
	fmt.Println(boolPingXing)
}

func pingxing(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2 float64) string {
	if ay1 == ay2 {
		if by1 == by2 {
			return "平行"
		} else {
			return "不平行"
		}
	}
	if by1 == by2 {
		if ay1 == ay2 {
			return "平行"
		} else {
			return "不平行"
		}
	}
	k1 := (ax1 - ax2) / (ay1 - ay2)
	k2 := (bx1 - bx2) / (by1 - by2)
	if k1 == k2 {
		return "平行"
	}
	return "不平行"
}
