package main

import "fmt"

type contactInfo struct {
	phoneNumber string
	address     string
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	me := person{
		firstName: "Dennis",
		lastName:  "Chege",
		contactInfo: contactInfo{
			phoneNumber: "+254704733211",
			address:     "220 - 10202",
		},
	}
	me.print()
	me.newFName("Robert")
	me.print()
}

func (user *person) newFName(name string) {
	(*user).firstName = name
}

func (user person) print() {
	fmt.Println(user)
}
