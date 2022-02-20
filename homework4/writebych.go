package main

import (
	"log"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

func main() {
	var Ps Persons
	chP := make(chan Person, 1)
	go func(chP chan Person, Ps Persons) {
		for {
			P := <-chP
			Ps = append(Ps, P)
			sort.Sort(Ps)
			log.Println("打印最新排名：")
			for _, p := range Ps {
				log.Println(p.name)
			}
		}
	}(chP, Ps)
	for i := 0; i < 5; i++ {
		go func(i int) {
			rand.Seed(time.Now().UnixNano())
			p := Person{
				name:   "person+" + strconv.Itoa(i),
				weight: float64(rand.Intn(65) + 45),
				tall:   rand.Float64()*0.4 + 1.6,
				sex:    float64(rand.Intn(2)),
				age:    rand.Intn(40) + 18,
			}
			p.calcBmi()
			p.calcFatRate()
			var sex string
			if p.sex == 1 {
				sex = "男"
			} else {
				sex = "女"
			}
			log.Printf("新成员加入！！姓名%s: ,性别：%s ,年龄：%d ,身高：%f米 ,体重:%f公斤 ,BMI：%f ,体脂率: %f \n", p.name, sex, p.age, p.tall, p.weight, p.bmi, p.fatRate)
			chP <- p
		}(i)
		time.Sleep(500 * time.Microsecond)
	}
	time.Sleep(1 * time.Second)
}
