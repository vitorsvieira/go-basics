package main

import (
	"fmt"
)

type Member struct {
	FirstName string
	LastName  string
	Language  string
}

func main() {

	slice := []Member{
		{"Rob", "Pike", "Go"},
		{"Andrew", "Gerrand", "Go"},
		{"Robert", "Griesemer", "Go"},
		{"Brad", "Fitzpatrick", "Go"},
	}
	ShowMembers(slice)
}

func ShowMembers(members []Member) {

	// the 'i' gives the loop need to use index value inside the block
	// if the block doesn't use the 'i' it will throw an error saying that a variable is not being used
	for i, m := range members {
		fmt.Println("Member nº", i+1, "-", m.Language+" developed by "+m.FirstName+" "+m.LastName)
	}

	for i := range members {
		fmt.Println("Member nº", i+1, "-", members[i].Language+" developed by "+members[i].FirstName+" "+members[i].LastName)
	}

	// the '_' (underscore) gives the loop the choice to iterate without the need to use the index
	for _, m := range members {
		fmt.Println(m.Language + " developed by " + m.FirstName + " " + m.LastName)
	}

}
