package main

import "fmt"

type Employee struct {
	firstname string
	lastname  string
	language  string
}

func Register(emp Employee) {
	fmt.Println("Fullname: ", emp.firstname, emp.lastname)
	fmt.Println("Language: ", emp.language)
}

func main() {

	var emp = Employee{"John", "Doe", "Go"}

	Register(emp)

}
