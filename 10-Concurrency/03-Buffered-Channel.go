package main

import (
	"fmt"
)

func main(){

	finished := make(chan bool, 2)

	go func() {
		OddNumbers(50)
		finished <- true
		finished <- true
		fmt.Println("Finished!")
	}()
	EvenNumbers(50)
	<- finished
}

func OddNumbers(stop int) {

	numbers := &[]int{}

	i:= 1

	for i <= stop{
		*numbers = append(*numbers, i)
		i = i + 2
	}
	fmt.Println("Odd:",*numbers)
}

func EvenNumbers(stop int) {

	numbers := &[]int{}

	i:= 2

	for i <= stop{
		*numbers = append(*numbers, i)
		i = i + 2
	}
	fmt.Println("Even:",*numbers)
}
