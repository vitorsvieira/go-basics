package main

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
	Language  string
}

func Register(emp Employee) {
	fmt.Println("Fullname: ", emp.FirstName, emp.LastName)
	fmt.Println("Language: ", emp.Language)

}

func main() {

	var emp = Employee{"John", "Doe", "Go"}

	Register(emp)

}
