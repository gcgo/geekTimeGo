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
				tall:   rand.Float64()*0.5 + 1.5,
				sex:    float64(rand.Intn(2)),
				age:    rand.Intn(40) + 18,
			}
			p.calcBmi()
			p.calcFatRate()
			log.Printf("姓名%s: ,性别：%f ,年龄：%d ,身高：%f ,体重:%f ,BMI：%f ,体脂率: %f \n", p.name, p.sex, p.age, p.tall, p.weight, p.bmi, p.fatRate)
			chP <- p
		}(i)
	}
	time.Sleep(5 * time.Second)
}
