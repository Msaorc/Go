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

func (u *User) CalcIMC() {
	u.imc = u.weight / (u.height * u.height)
}

func main() {
	user := User{
		name:     "Marcos",
		passWord: "******",
		age:      30,
		weight:   96,
		height:   1.72,
		sex:      "male",
	}

	user.CalcIMC()

	fmt.Printf("The imc for %s is: %0.2f", user.name, user.imc)
}
