package main

import (
	"fmt"
)

type Member struct {
	firstname string
	lastname  string
	language  string
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
		fmt.Println("Member nº", i+1, "-", m.language+" developed by "+m.firstname+" "+m.lastname)
	}

	for i := range members {
		fmt.Println("Member nº", i+1, "-", members[i].language+" developed by "+members[i].firstname+" "+members[i].lastname)
	}

	// the '_' (underscore) gives the loop the choice to iterate without the need to use the index
	for _, m := range members {
		fmt.Println(m.language + " developed by " + m.firstname + " " + m.lastname)
	}

}
