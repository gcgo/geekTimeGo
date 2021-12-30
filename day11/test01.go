package main

import (
	"fmt"
	"learn.go/day06/tizhi"
)

type Person struct {
	name   string
	sex    string
	tall   float64
	weight float64
	age    int

	bmi     float64
	fatRate float64
}

func (p *Person) calBMI() {
	bmi := tizhi.BMI(p.weight, p.tall)
	p.bmi = bmi
}

func main() {
	p := Person{name: "liuyang105", tall: 1.95, weight: 100, age: 31}
	p.calBMI()
	fmt.Println(p.bmi)
}
