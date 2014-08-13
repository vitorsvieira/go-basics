package main

import (
"fmt"
)

func main() {

	// '&' is called the address-of operator
	// when placed before a variable it gives the memory address of that variable

	// '*' is a type modifier when placed before the type of the value
	// placing * before the pointer is called the dereference operator
	name := "Gopher"
	var ptrname *string = &name
	//name -> value
	//&name -> mem address
	fmt.Println(name, &name)
	//ptrname -> referenced memory address
	//&ptrname -> local memory address
	//*ptrname -> value
	fmt.Println(ptrname, &ptrname, *ptrname)

	fmt.Println("Print 1: ", name, ptrname)
	fmt.Println("Print 2: ", name, *ptrname)

	//change pointer value, reflecting on name variable
	*ptrname = "new name"

	fmt.Println("Print 3: ", name, *ptrname)

}
