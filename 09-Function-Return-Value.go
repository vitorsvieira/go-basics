package main

import (
	"fmt"
)

type Employee struct {
	firstname string
	lastname  string
	language  string
}

func Register(emp Employee) {
	fmt.Println("Fullname: ", GetFullname(emp.firstname, emp.lastname))
	fmt.Println("Language: ", emp.language)
}

func GetFullname(firstname, lastname string) string {
	return firstname + " " + lastname
}

func main() {
	var emp = Employee{"John", "Doe", "Go"}

	Register(emp)
}
