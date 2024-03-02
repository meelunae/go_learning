package main

import "fmt"

type ContactInfo struct {
	email   string
	zipCode int
}

type Person struct {
	firstName string
	lastName  string
	contacts  ContactInfo
}

func main() {
	rin := Person{
		firstName: "Rin",
		lastName:  "Asunaro",
		contacts:  ContactInfo{"rin@ac.me", 12345},
	}
	rin.updateName("Asuka")
	rin.print()
}

func (p *Person) updateName(newName string) {
	p.firstName = newName
}

func (p Person) print() {
	fmt.Printf("%+v", p)
}
