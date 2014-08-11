package main

import (
	"fmt"
)

type Member struct {
	FirstName string
	LastName  string
	Language  string
}

type Renamable interface {
	Rename(firstName, lastName string)
}

func main() {

	members := Members{
		{"Rob", "Pike", "Go"},
		{"Andrew", "Gerrand", "Go"},
		{"Robert", "Griesemer", "Go"},
		{"Brad", "Fitzpatrick", "Go"},
	}

}

func RenameToEmpty(r Renamable) {
	r.Rename("Empty", "Empty")
}

func (member *Member) Rename(firstName, lastName string){
	member.FirstName = firstName
	member.LastName = lastName
}
