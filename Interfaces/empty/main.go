package main

import "fmt"

func main() {
	var x interface{} = "Marcos Test"
	var y interface{} = 100.00
	showType(x)
	showType(y)
}

func showType(i interface{}) {
	fmt.Printf("O value is %v and type is: %T\n", i, i)
}
