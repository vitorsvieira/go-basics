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
	fullname, lang := GetFullnameAndLang(emp.firstname, emp.lastname, emp.language)
	fmt.Println("Fullname: ", fullname)
	fmt.Println("Language: ", lang)
}

func GetFullnameAndLang(firstname, lastname, language string) (fullname string, lang string) {

	fullname = firstname + " " + lastname
	lang = language + "lang !!!"

	return
}

func main() {
	var emp = Employee{"John", "Doe", "Go"}

	Register(emp)
}
