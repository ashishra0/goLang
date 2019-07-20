package main

import "fmt"

type contactInfo struct {
	phone int
	email string
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Halpert",
		contactInfo: contactInfo{
			phone: 8806011009,
			email: "jim@office.com",
		},
	}

	contact := contactInfo{
		phone: 8806011009,
		email: "ashisrao@email.com",
	}

	contactPointer := &contact
	contactPointer.updatePhone(9822147009)

	jimPointer := &jim
	jimPointer.updateName("Phil")

	fmt.Println(jim)
	fmt.Println(contact)
}
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (pointerToContact *contactInfo) updatePhone(newNo int) {
	(*pointerToContact).phone = newNo
}
