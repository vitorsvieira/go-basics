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

	//2 is the length and 8 is the capacity
	//the length stores the current available slots
	//most of the time capacity will not be provided
	var slice1 []int
	slice1 = make([]int, 2, 8)
	slice1[0] = 1
	slice1[1] = 2

	var slice2 []int
	slice2 = make([]int, 4)
	slice2[0] = 1
	slice2[1] = 2
	slice2[2] = 3
	slice2[3] = 4

	slice3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	members := []Member{
		{"Rob", "Pike", "Go"},
		{"Andrew", "Gerrand", "Go"},
		{"Robert", "Griesemer", "Go"},
		{"Brad", "Fitzpatrick", "Go"},
	}

	fmt.Println("Total: ", len(slice1), "Values: ", slice1)
	fmt.Println("Total: ", len(slice2), "Values: ", slice2)
	fmt.Println("Total: ", len(slice3), "Values: ", slice3)
	fmt.Println("Total: ", len(members), "Values: ", members)

	SliceTheSlice(members)

	AppendToSlice(members)

	DeleteSlice(members)

}

func SliceTheSlice(members []Member) {

	fmt.Println("First Member: ", members[0:1])
	fmt.Println("Second Member: ", members[1:2])
	fmt.Println("First Two Members: ", members[0:2])
	fmt.Println("First Two Members: ", members[:2])
	fmt.Println("Last Two Members: ", members[2:4])
	fmt.Println("Last Two Members: ", members[len(members)-2:len(members)])
	fmt.Println("Last Two Members: ", members[len(members)-2:])
	fmt.Println("Last One Members: ", members[len(members)-1:len(members)])
	fmt.Println("Last One Members: ", members[len(members)-1:])

	fmt.Println("All Members: ", members[0:])
	fmt.Println("All Members: ", members[0:len(members)])
	fmt.Println("All Members: ", members[:len(members)])
	fmt.Println("All Members: ", members[len(members)-len(members):len(members)])
}

func AppendToSlice(members []Member) {

	newSlice := append(members, Member{"Go", "Lang", "Go"})
	fmt.Println("Appended: ", newSlice)

	clonedSlice := append(members, members...)
	fmt.Println("Total: ", len(clonedSlice), "Cloned: ", clonedSlice)
}

func DeleteSlice(members []Member) {

	deleteFirstMember := append(members[1:])
	fmt.Println("Delete First Member:", deleteFirstMember)

	deleteLastMember := append(members[:len(members)-1])
	fmt.Println("Delete Last Member:", deleteLastMember)

	deleteFirstAndLast := append(members[1:2], members[2:3]...)
	// the '...' basically expands the 'members' to the second append argument
	// ignoring the index modification made at the first append slot
	fmt.Println("Delete First and Last: ", deleteFirstAndLast)

}
