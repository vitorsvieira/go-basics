package main

import (
	"fmt"
)

const NAME = "Golang"
const SHORTNAME = "Go"

const (
	NICKNAME   = "Gopher"
	TEXTEDITOR = "Sublime Text"
	IDE        = "Intellij"
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

	fmt.Println(NAME, SHORTNAME)

	fmt.Println(NICKNAME)
	fmt.Println(TEXTEDITOR)
	fmt.Println(IDE)

	fmt.Println(X, Y, Z)
	fmt.Println(I, J, K)
}
