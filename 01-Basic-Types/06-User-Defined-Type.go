package main

import (
	"fmt"
)

type Developer struct {
	Name string
	Ide  string
}

func main() {

	var user1 = Developer{"Gopher", "Sublime Text"}

	fmt.Println(user1.Name, user1.Ide)

	var user2 = Developer{Name: "Golang", Ide: "LiteIDE"}

	fmt.Println(user2.Name, user2.Ide)

	var user3 = Developer{}
	user3.Name = "Go"
	user3.Ide = "Vim"

	fmt.Println(user3.Name, user3.Ide)

	var user4 = Developer{}
	user4.Ide = "Intellij"

	fmt.Println(user4.Name, user4.Ide)
}
