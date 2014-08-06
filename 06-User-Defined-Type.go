package main

import (
	"fmt"
)

type Developer struct {
	name string
	ide  string
}

func main() {

	var user1 = Developer{"Gopher", "Sublime Text"}

	fmt.Println(user1.name, user1.ide)

	var user2 = Developer{name: "Golang", ide: "LiteIDE"}

	fmt.Println(user2.name, user2.ide)

	var user3 = Developer{}
	user3.name = "Go"
	user3.ide = "Vim"

	fmt.Println(user3.name, user3.ide)

	var user4 = Developer{}
	user4.ide = "Intellij"

	fmt.Println(user4.name, user4.ide)
}
