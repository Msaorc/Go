package main

import "fmt"

type NewType int

var (
	a NewType = 2
	b string  = "Hello World"
)

func main() {
	c := "Assigning new value"
	fmt.Printf("The type of the variable (C) is: %T \n\n", c)
	c = fmt.Sprintf("The type of the variable (A) is: %T", a)
	fmt.Println(c)
}
