package main

import "fmt"

func main() {
	var myInterface interface{} = "M&A System"
	fmt.Println(myInterface.(string))

	res, ok := myInterface.(int)
	fmt.Printf("The Value of (res) is: %v and the result of ok: %v", res, ok)
}
