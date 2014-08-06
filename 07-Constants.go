package main

import (
	"fmt"
)

const name = "Golang"
const shortname = "Go"

const (
	nickname   = "Gopher"
	texteditor = "Sublime Text"
	ide        = "Intellij"
)
const (
	//iota is a reserved word / keyword that represents untyped integer constants
	X = iota
	Y = iota
	Z = iota
	I
	J
	K
)

func main() {

	fmt.Println(name, shortname)

	fmt.Println(nickname)
	fmt.Println(texteditor)
	fmt.Println(ide)

	fmt.Println(X, Y, Z)
	fmt.Println(I, J, K)
}
