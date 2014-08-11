package main

import (
	"fmt"
)

type Student struct {
	FirstName, LastName string
}

type FullNamer interface {
	Fullname() string
}

func (student *Student) Fullname() string {

	return fmt.Sprintf("%s %s", student.FirstName, student.LastName)
}

func CallStudent(f FullNamer) string {

	return fmt.Sprintf("Welcome! %s", f.Fullname())
}

func main() {

	student := &Student{"Rob", "Pike"}

	fmt.Println(CallStudent(student))
}
