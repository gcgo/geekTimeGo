package main

import "log"

type Person struct {
	name   string
	sex    float64
	tall   float64
	weight float64
	age    int

	bmi     float64
	fatRate float64
}

func (p *Person) calcBmi() error {
	bmi, err := BMI(p.weight, p.tall)
	if err != nil {
		log.Printf("error when calculating BMP for Person[%s]: %v", p.name, err)
		return err
	}
	p.bmi = bmi
	return nil
}

func (p *Person) calcFatRate() {
	p.fatRate = Tizhi(p.bmi, float64(p.age), p.sex)
}

// Persons a set of person
type Persons []Person

// Len return count
func (p Persons) Len() int {
	return len(p)
}

// Less return bigger true
func (p Persons) Less(i, j int) bool {
	return p[i].fatRate < p[j].fatRate
}

// Swap swap items
func (p Persons) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
