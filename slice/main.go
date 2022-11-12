package main

import "fmt"

func main() {
	slice := []int{10, 20, 30, 40, 50, 60}

	fmt.Printf("Length for slice is: %d, capacity is: %d, value: %v \n", len(slice), cap(slice), slice)

	fmt.Printf("Length for slice is: %d, capacity is: %d, value: %v \n", len(slice[0:5]), cap(slice), slice)

	fmt.Printf("Length for slice is: %d, capacity is: %d, value: %v \n", len(slice), cap(slice[1:2]), slice)

	fmt.Printf("Length for slice is: %d, capacity is: %d, value: %v \n", len(slice), cap(slice), slice[3:6])
}
