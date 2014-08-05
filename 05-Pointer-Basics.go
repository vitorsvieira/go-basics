package main

import (
	"fmt"
)

func main() {

	var name string = "Gopher"
	var pointername *string = &name

	fmt.Println("Print 1: ", name, pointername)
	fmt.Println("Print 2: ", name, *pointername)

	*pointername = "new name"

	fmt.Println("Print 3: ", name, *pointername)
}
