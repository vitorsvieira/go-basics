package main

import (
	"fmt"
)

func main() {

	//arrays have a defined size that cannot be changed during runtime
	//less flexible than slices but it is more efficient
	//the array size (the memory allocated) must be defined at the moment of declaration
	x := [5]int{1, 2, 3, 4, 5}
	fmt.Println(x)

	y := [3]int{1, 2, 3}
	fmt.Println(y)

	//array types can only be compared if they have the same type and length
	//running the comparison below will throw the following error
	//"invalid operation: x == y (mismatched types [5]int and [3]int)"
	//they have the same type but a different length
	/*if x == y {
		fmt.Println(true)
	}*/
}
