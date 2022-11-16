package main

import "fmt"

type NewType int

var (
	a NewType = 2
	b string  = "Hello World"
	d bool    = false
)

func main() {
	c := "Assigning new value"
	c = fmt.Sprintf("The type of the variable (C) is: %T and value %v", c, c)
	fmt.Printf("The type of the variable (A) is: %T and value %v \n\n", a, a)
	fmt.Printf("The type of the variable (B) is: %T and value %v \n\n", b, b)
	fmt.Println(c)
	fmt.Printf("\nThe type of the variable (D) is: %T and value %v", d, d)
}
