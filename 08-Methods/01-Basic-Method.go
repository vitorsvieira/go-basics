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

	fmt.Println("Adding new member....\n")
	members.Add(Member{"Go", "Lang", "Golang"})

	fmt.Println("Done!\n====UPDATED LIST====")
	members.ShowAll()

	fmt.Println("\nRenaming Members...\n")
	members.ClearNames("--", "--")
	fmt.Println("Done!\n====UPDATED LIST====")
	members.ShowAll()

}

func (members *Members) Add(member Member) {

	*members = append(*members, Member{member.FirstName, member.LastName, member.Language})
}

func (members *Members) ClearNames(firstName, lastName string) {

	for i := 0; i < len(*members); i++ {
		(*members)[i].FirstName = firstName
		(*members)[i].LastName = lastName
	}
}

func (members Members) ShowAll() {

	for _, m := range members {
		fmt.Println("Mr.", m.FirstName, m.LastName)
	}
}
