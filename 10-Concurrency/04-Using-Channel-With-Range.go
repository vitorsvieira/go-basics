package main

import (
	"fmt"
)

type Number int
type Numbers []Number

func main() {

	evenNumbers := EvenNumbers(50)

	c := make(chan Number)

	go evenNumbers.ChannelNumbers(c)

	for n := range c {
		fmt.Println(n)
	}
}

func (numbers Numbers) ChannelNumbers(c chan Number) {

	for _, n := range numbers {
		c <- n
	}
	close(c)
}

func EvenNumbers(stop Number) *Numbers {

	numbers := &Numbers{}

	var i Number = 2

	for i <= stop {
		*numbers = append(*numbers, i)
		i = i + 2
	}
	return numbers
}
