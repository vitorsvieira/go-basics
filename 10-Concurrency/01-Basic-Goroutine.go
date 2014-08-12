package main

import (
	"fmt"
	"time"
)

func main() {

	go PrintEvenNumbers(16)
	PrintOddNumbers(19)
	//by using 'go' keyword before the function call
	//the even numbers will be printed by a separate thread
	//freeing the next function for execution

	//sleep will give the thread sufficient time to be initialized/executed
	time.Sleep(100 * time.Millisecond)
}

func PrintOddNumbers(stop int){

	i:= 1
	for i <= stop{

		fmt.Println("odd:", i)
		i = i + 2
	}
}

func PrintEvenNumbers(stop int){

	i:= 2
	for i <= stop{

		fmt.Println("even:", i)
		i = i + 2
	}
}
