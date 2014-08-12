package main

import (
	"fmt"
	"time"
)

func main(){

	finished := make(chan bool)

	go func() {
		OddNumbers(27)
		finished <- true
	}()
	time.Sleep(1 * time.Millisecond)
}

func OddNumbers(stop int) {

	numbers := &[]int{}

	i:= 1

	for i <= stop{
		*numbers = append(*numbers, i)
		i = i + 2
	}
	fmt.Println(*numbers)
}
