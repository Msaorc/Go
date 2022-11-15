package main

import "fmt"

func main() {
	total, message := func() (int, string) {
		return multiplication(587, 75), func() string {
			if multiplication(587, 75) > 10000 {
				return "Number is greater than 10000"
			}
			return "Number is not greater than 10000"
		}()
	}()

	fmt.Printf("The result the multiplication is: %d, the of message is: %s", total, message)
}

func multiplication(firtValue, secondValue int) int {
	return firtValue * secondValue
}
