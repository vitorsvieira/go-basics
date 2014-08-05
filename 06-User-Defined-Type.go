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

}
