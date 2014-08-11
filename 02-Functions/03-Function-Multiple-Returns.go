package main

import (
	"fmt"
)

type Employee struct {
	FirstName string
	LastName  string
	Language  string
}

func Register(emp Employee) {
	fullName, lang := GetFullnameAndLang(emp.FirstName, emp.LastName, emp.Language)
	fmt.Println("Fullname: ", fullName)
	fmt.Println("Language: ", lang)
}

func GetFullnameAndLang(firstName, lastName, language string) (string, string) {
	return firstName + " " + lastName, language + "lang !!!"
}

func main() {
	var emp = Employee{"John", "Doe", "Go"}

	Register(emp)
}
