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
	fmt.Println("Printing employee information")
	fmt.Println(jim)
}
