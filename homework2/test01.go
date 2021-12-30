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
		bmi := githubbmi.BMI(weight, height)
		tizhi := githubbmi.Tizhi(bmi, age, sex)
		fmt.Println(name, "体脂是：", tizhi)
		level, advice := giveAdvice(age, tizhi, sex)
		ziliao[name] = fmt.Sprintf("bmi是：%f,体脂是：%f,级别是：%s，建议是：%s", BMI(weight, height), tizhi, level, advice)
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

/**
BMI=体重(公斤)÷(身高×身高)(米)

*/
func BMI(weight, height float64) float64 {
	return weight / height / height
}

/**
Tizhi
计算体脂：1.2×BMI+0.23×年龄-5.4-10.8×性别(男为1，女为0)
*/
func Tizhi(bmi, age, sex float64) float64 {
	if sex == 1 {
		return 1.2*bmi + 0.23*age - 5.4 - 10.8
	} else if sex == 0 {
		return 1.2*bmi + 0.23*age - 5.4
	}
	return 0
}

func giveAdvice(age, tizhi, sex float64) (string, string) {
	if sex == 1 {
		if age >= 18 && age <= 39 {
			if tizhi <= 10 {
				return "偏瘦", "多吃"
			} else if tizhi > 10 && tizhi <= 16 {
				return "标准", "保持"
			} else if tizhi > 16 && tizhi <= 21 {
				return "偏重", "运动"
			} else if tizhi > 21 && tizhi <= 26 {
				return "肥胖", "减肥"
			} else {
				return "严重肥胖", "没救了"
			}
		} else if age >= 40 && age <= 59 {
			if tizhi <= 11 {
				return "偏瘦", "多吃"
			} else if tizhi > 11 && tizhi <= 17 {
				return "标准", "保持"
			} else if tizhi > 17 && tizhi <= 22 {
				return "偏重", "运动"
			} else if tizhi > 22 && tizhi <= 27 {
				return "肥胖", "减肥"
			} else {
				return "严重肥胖", "没救了"
			}
		} else {
			if tizhi <= 13 {
				return "偏瘦", "多吃"
			} else if tizhi > 13 && tizhi <= 19 {
				return "标准", "保持"
			} else if tizhi > 19 && tizhi <= 24 {
				return "偏重", "运动"
			} else if tizhi > 24 && tizhi <= 29 {
				return "肥胖", "减肥"
			} else {
				return "严重肥胖", "没救了"
			}
		}
		//如果是女性
	} else {
		if age >= 18 && age <= 39 {
			if tizhi <= 20 {
				return "偏瘦", "多吃"
			} else if tizhi > 20 && tizhi <= 27 {
				return "标准", "保持"
			} else if tizhi > 27 && tizhi <= 34 {
				return "偏重", "运动"
			} else if tizhi > 34 && tizhi <= 39 {
				return "肥胖", "减肥"
			} else {
				return "严重肥胖", "没救了"
			}
		} else if age >= 40 && age <= 59 {
			if tizhi <= 21 {
				return "偏瘦", "多吃"
			} else if tizhi > 21 && tizhi <= 28 {
				return "标准", "保持"
			} else if tizhi > 28 && tizhi <= 35 {
				return "偏重", "运动"
			} else if tizhi > 35 && tizhi <= 40 {
				return "肥胖", "减肥"
			} else {
				return "严重肥胖", "没救了"
			}
		} else {
			if tizhi <= 22 {
				return "偏瘦", "多吃"
			} else if tizhi > 22 && tizhi <= 29 {
				return "标准", "保持"
			} else if tizhi > 29 && tizhi <= 36 {
				return "偏重", "运动"
			} else if tizhi > 36 && tizhi <= 41 {
				return "肥胖", "减肥"
			} else {
				return "严重肥胖", "没救了"
			}
		}
	}
}
