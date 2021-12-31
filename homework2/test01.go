package main

import (
	"fmt"
	githubbmi "github.com/armstrongli/go-bmi"
)

func main() {
	var name string
	var sex, weight, height, age float64
	var next int
	ziliao := make(map[string]string)
	var tizhiToal []float64

	for {
		fmt.Println("姓名：")
		fmt.Scanln(&name)

		fmt.Println("性别：(男：1，女：0)")
		fmt.Scanln(&sex)

		fmt.Println("体重(公斤)：")
		fmt.Scanln(&weight)

		fmt.Println("身高（米）：")
		fmt.Scanln(&height)

		fmt.Println("年龄：")
		fmt.Scanln(&age)
		bmi, _ := githubbmi.BMI(weight, height)
		tizhi := githubbmi.Tizhi(bmi, age, sex)
		fmt.Println(name, "体脂是：", tizhi)
		level, advice, _ := githubbmi.GiveAdvice(age, tizhi, sex)
		ziliao[name] = fmt.Sprintf("bmi是：%f,体脂是：%f,级别是：%s，建议是：%s", bmi, tizhi, level, advice)
		tizhiToal = append(tizhiToal, tizhi)
		fmt.Println("是否继续下一轮计算：是：1，否:任意输入")
		fmt.Scanln(&next)
		if next != 1 {
			break
		}
	}
	for name, val := range ziliao {
		fmt.Println(name, "资料为：", val)
	}
	var tizhiAll float64
	for _, v := range tizhiToal {
		tizhiAll += v
	}
	fmt.Println(len(tizhiToal), "个人的平均体脂为：", tizhiAll/float64(len(tizhiToal)))
}
