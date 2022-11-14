package main

import "fmt"

func main() {
	imcCustomer1 := calulateImc(80, 1.80)
	imc1, message1 := validateImc(imcCustomer1)
	fmt.Printf("\nThe imc of Customer1 is: %.2f, and message for Customer1 is: %s\n\n", imc1, message1)

	imcCustomer2 := calulateImc(90, 1.80)
	imc2, message2 := validateImc(imcCustomer2)
	fmt.Printf("The imc of Customer2 is: %.2f, and message for Customer2 is: %s", imc2, message2)
}

func validateImc(imc float32) (float32, string) {
	if imc >= 18.5 && imc <= 25 {
		return imc, "healthy"
	}

	if imc >= 25 && imc <= 30 {
		return imc, "overweight"
	}

	if imc >= 30 && imc <= 35 {
		return imc, "grade 1 obesity"
	}

	return 0, "Invalid IMC value"
}

func calulateImc(weight, height float32) float32 {
	return weight / (height * height)
}
