package main

import (
	"Go-Fundamentals/modules/module"
	"fmt"
)

func main() {
	user := module.User{
		Name:     "Marcos",
		PassWord: "******",
		Age:      30,
		Weight:   96,
		Height:   1.72,
		Sex:      "male",
	}

	user.CalcIMC()
	defer fmt.Printf("\nIMC with Defer: %.2f", user.Imc)
	fmt.Printf("IMC no Defer: %.2f", user.Imc)
}
