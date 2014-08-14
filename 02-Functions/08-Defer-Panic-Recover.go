package main

import "fmt"

func main() {

	//the defer statement schedules a function to be called after the function completes
	defer FortyToSixty()  // <- this defer schedules the function to be executed last

	defer TwentyToForty() // <- try to schedule at last position but has a function already there. This one will be executed first

	defer func() {        // <- this is now the first one\
		message := recover()
		fmt.Println(message)
	}()

	OneToTwenty() 		  // <- executed as the normal flow

	DivideByZero(5,0) 	  // <- throws a panic to be recovered by the deferred func
}

func OneToTwenty(){

	numbers := &[]int{}

	for i:= 0; i<= 20; i++{
		*numbers = append(*numbers, i)
	}
	fmt.Println(*numbers)
}

func TwentyToForty(){

	numbers := &[]int{}

	for i:= 20; i<= 40; i++{
		*numbers = append(*numbers, i)
	}
	fmt.Println(*numbers)
}

func FortyToSixty(){

	numbers := &[]int{}

	for i:= 40; i<= 60; i++{
		*numbers = append(*numbers, i)
	}
	fmt.Println(*numbers)
}

func SixtyToEighty(){

	numbers := &[]int{}

	for i:= 60; i<= 80; i++{
		*numbers = append(*numbers, i)
	}
	fmt.Println(*numbers)
}

func DivideByZero(x, y int) int{
	if x > 0 && y >0 {
		return x / y
	}
	panic(fmt.Sprintf("Error!!! Invalid numbers %d and %d", x, y))
}
