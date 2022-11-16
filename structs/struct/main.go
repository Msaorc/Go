package main

import "fmt"

type User struct {
	name     string
	passWord string
	age      int
	weight   float32
	height   float32
	fatPerc  float32
	muscle   float32
	water    float32
	imc      float32
}

func main() {
	richard := User{
		name:     "Ricard",
		passWord: "******",
		age:      30,
		weight:   96,
		height:   172,
		fatPerc:  33.5,
		muscle:   61,
		water:    40,
	}

	richard.imc = 34

	fmt.Printf("The IMC for %s is: %.2f", richard.name, richard.imc)
}
