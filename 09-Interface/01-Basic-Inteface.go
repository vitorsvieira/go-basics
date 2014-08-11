package main

import (
	"fmt"
)

type Student struct {
	FirstName, LastName string
}

type Instructor struct {
	FirstName, Course string
}

type FullNamer interface {
	Fullname() string
}

func (student *Student) Fullname() string {

	return fmt.Sprintf("%s %s", student.FirstName, student.LastName)
}

func (instructor *Instructor) Fullname() string {

	return fmt.Sprintf("%s from %s course", instructor.FirstName, instructor.Course)
}

func Welcome(f FullNamer) string {

	return fmt.Sprintf("Welcome! %s", f.Fullname())
}

func main() {

	student := &Student{"John", "Doe"}
	instructor := &Instructor{"Rob", "Golang"}

	fmt.Println(Welcome(student))
	fmt.Println(Welcome(instructor))
}
