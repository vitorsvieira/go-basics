package main

import (
	"fmt"
)

type Number int
type Numbers []Number

func main() {

	evenNumbers := EvenNumbers(50)
	oddNumbers := OddNumbers(50)

	chan1 := make(chan Number)
	chan2 := make(chan Number)

	go evenNumbers.ChannelNumbers(chan1)
	go oddNumbers.ChannelNumbers(chan2)

	for {
		select {
		case n1, ready1 := <-chan1:
			if ready1 {
				fmt.Println("Chan1:", n1)
			} else {
				return
			}
		case n2, ready2 := <-chan2:
			if ready2 {
				fmt.Println("Chan2:", n2)
			} else {
				return
			}
		}
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

func OddNumbers(stop Number) *Numbers {

	numbers := &Numbers{}

	var i Number = 1

	for i <= stop {
		*numbers = append(*numbers, i)
		i = i + 2
	}
	return numbers
}
