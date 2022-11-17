package main

import (
	"Go-Fundamentals/modules/module"
	"fmt"
)

func main() {
	fmt.Printf("The result of sum is: %v \n\n", module.Sum(20, 20))
	fmt.Printf("The result of subtraction is: %v\n\n", module.Subtraction(40, 20))

	myUser := module.User{
		Name:     "Marcos",
		PassWord: "******",
		Age:      30,
		Weight:   96,
		Height:   1.72,
		Sex:      "male",
	}

	myUser.CalcIMC()
	fmt.Printf("The IMC for :%s id %.2f", myUser.Name, myUser.Imc)
}
