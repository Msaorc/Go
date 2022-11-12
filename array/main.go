package main

import "fmt"

func main() {
	var myArrayIntegers [3]int
	myArrayIntegers[0] = 10
	myArrayIntegers[1] = 20
	myArrayIntegers[2] = 33

	var myArrayStrings [3]string
	myArrayStrings[0] = "First Position of my array strings"
	myArrayStrings[1] = "Second Position of my array strings"
	myArrayStrings[2] = "Third Position of my array strings"

	println("Cycling through an array of Integers")

	for index, value := range myArrayIntegers {
		fmt.Printf("The index is: %d and value is: %d \n", index, value)
	}

	println("\nCycling throgh an array of Strings")

	for index, value := range myArrayStrings {
		fmt.Printf("The index is: %d and value is: (%s) \n", index, value)
	}
}
