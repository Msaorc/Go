package main

import "fmt"

func main() {
	salarys := map[string]float32{}
	salarys["master"] = 15.000
	salarys["manager"] = 10.000
	salarys["floorwalker"] = 5.000

	customers := make(map[string]string)
	customers["0"] = "Joao"
	customers["1"] = "Marcos"
	customers["2"] = "Thiago"

	responsibilitys := map[string]string{"Joao": "master", "Marcos": "manager", "Thiago": "floorwalker"}

	for _, nameCustomer := range customers {
		var customerResponsibility string
		var customerSalary float32

		for name, resposibilityCustomer := range responsibilitys {
			if nameCustomer == name {
				customerResponsibility = resposibilityCustomer

				for responsibility, salary := range salarys {
					if responsibility == customerResponsibility {
						customerSalary = salary
						break
					}
				}
			}
		}

		fmt.Printf("The Contributor %s, have the job %s with the salary %.3f \n\n", nameCustomer, customerResponsibility, customerSalary)
	}
}
