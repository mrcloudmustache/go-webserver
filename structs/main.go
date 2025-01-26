package main

import "fmt"

type contact struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contact
}

func main() {
	jim := person{
		firstName: "James",
		lastName:  "Thomas",
		contact: contact{
			email:   "abc@gmail.com",
			zipCode: 55124,
		},
	}
	// Print memory address
	// fmt.Printf("%p\n", &jim)
	jim.updateName("Andy")
	jim.print()

}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

// Prints struct properties and their values
func (p person) print() {
	fmt.Printf("%+v\n", p)
}
