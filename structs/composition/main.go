package main

import "fmt"

type User struct {
	name     string
	passWord string
	age      int
	weight   float32
	height   float32
	sex      string
	BodyComposition
}

type BodyComposition struct {
	fatPerc float32
	muscle  float32
	water   float32
	imc     float32
}

func main() {
	maria := User{
		name:     "Maria",
		passWord: "******",
		age:      30,
		weight:   96,
		height:   172,
		sex:      "women",
	}

	// composition direct access
	maria.fatPerc = 25
	maria.muscle = 40
	maria.water = 20
	maria.imc = 25

	fmt.Printf("The body composition for %s is: FatPerc: %.2f, Muscle: %.2f, Water: %.2f, Imc: %.2f", maria.name, maria.fatPerc, maria.muscle, maria.water, maria.imc)

}
