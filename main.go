package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {

	theMan := person{
		firstName: "Party",
		lastName:  "Pete",
		contactInfo: contactInfo{
			email:   "theman@party.com",
			zipCode: 42069,
		},
	}

	theManPointer := &theMan
	theManPointer.updateName("International Superstar Gunther")
	theMan.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

