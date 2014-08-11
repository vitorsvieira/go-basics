package main

import (
	"fmt"
)

type Member struct {
	FirstName string
	LastName  string
	Language  string
}

type Members []Member

func main() {

	members := Members{
		{"Rob", "Pike", "Go"},
		{"Andrew", "Gerrand", "Go"},
		{"Robert", "Griesemer", "Go"},
		{"Brad", "Fitzpatrick", "Go"},
	}

	fmt.Println("Adding member....")
	members.Add(Member{"Go", "Lang", "Golang"})

	fmt.Println("Final Result...")
	members.ShowAll()

}

func (members *Members) Add(member Member) {

	*members = append(*members, Member{member.FirstName, member.LastName, member.Language})
}

func (members Members) ShowAll() {

	for _, m := range members {

		fmt.Println("Mr.", m.FirstName, m.LastName)
	}
}
