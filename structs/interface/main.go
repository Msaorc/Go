package main

import "fmt"

type User struct {
	name   string
	active bool
	cpf    string
}

type Company struct {
	name   string
	active bool
	cnpj   string
}

type ActivateDisable interface {
	Activate()
	Disable()
}

func (c *Company) Activate() {
	c.active = true
}

func (c *Company) Disable() {
	c.active = false
}

func (c *User) Activate() {
	c.active = true
}

func (c *User) Disable() {
	c.active = false
}

func main() {
	user := User{
		name: "Marcos Augusto",
	}

	company := Company{
		name: "M&A System",
	}

	ActivateDisable.Activate(&user)
	fmt.Printf("The User %s is active? %t\n\n", user.name, user.active)

	ActivateDisable.Disable(&user)
	fmt.Printf("The User %s is active? %t\n\n", user.name, user.active)

	ActivateDisable.Activate(&company)
	fmt.Printf("The Company %s is active? %t\n\n", company.name, company.active)

	ActivateDisable.Disable(&company)
	fmt.Printf("The Company %s is active? %t\n\n", company.name, company.active)
}
