package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	//	alex := person {
	//		firstName: "Alex",
	//		lastName: "Smith",
	//		contactInfo: {
	//			email: "blah@email.com",
	//			zipCode: 12345,
	//		},
	//	}

	// var allows use of initializing with zero value
	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Smith"
	alex.contactInfo = contactInfo{}
	alex.updateName("replacement")
	alex.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(name string) {
	p.firstName = name
}
