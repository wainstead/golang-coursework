package main

import "fmt"
// bleh tabs. must fix.
type contactInfo struct {
	email   string
	zipCode string
}

// note the simplicity of the declaration syntax: no commas, colons, etc.
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
	// can also just say "contactInfo" which will be important later on
}

func main() {
	// instructor hates this syntax. First way:
	// alex := person{"Alex", "Anderson"}
	// Second way: specify the field names
	alex := person{firstName: "Alex", lastName: "Anderson"}
	//fmt.Println(alex)
	alex.printPerson()
	var alex2 person
	// I can do this because the fields are assigned default values (empty strings here)
	fmt.Println("just alex2:")
	fmt.Println(alex2)
	fmt.Println("printPerson alex2:")
	alex2.printPerson()
	fmt.Println("Printf with 'v'") // prints the struct structure, it appears
	fmt.Printf("%+v\n", alex2)
	// Assigning values to struct fields
	alex2.firstName = "Steve"
	alex2.lastName = "Wainstead"
	alex2.contact.email = "nobody@nowhere.com"
	alex2.contact.zipCode = "10001"
	alex2.printPerson()

	bob := person{
		firstName: "J.R.",
		lastName:  "Dobbs",
		contact: contactInfo{
			email:   "dobbs@slack.com",
			zipCode: "94000",
		},
	}
	bob.printPerson()
	fmt.Printf("%+v\n", bob)
	bob.updateNamePassByValue("Jimmy") // pass by value, so does not affect our "bob"

	// Long form: get the address and call the function
	bobPointer := &bob
	bobPointer.updateName("Jimmy")
	bob.printPerson()

	// or just:
	bob.updateName("Raphael")
	bob.printPerson()
}

func (p person) printPerson() {
	fmt.Println(p.firstName, p.lastName)
	fmt.Println(p.contact.email, p.contact.zipCode)
}

// Go is pass by value; the name change here has no effect from the caller's perspective
func (p person) updateNamePassByValue(newFirstName string) {
	p.firstName = newFirstName
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

// Turn address into value with *address
// Turn value into address with &value
// Declare a variable to point to a type with *
