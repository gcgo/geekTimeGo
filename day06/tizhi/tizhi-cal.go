package tizhi

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

func Aa() {

}

func GiveAdvice(age, tizhi, sex float64) (string, string) {
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
