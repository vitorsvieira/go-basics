package main

import (
	"fmt"
)

func main() {

	getOddNumbers := func (limit int) []int {
		var result []int
		i:= 1
		for i <= limit{
			result = append(result, i)
			i = i + 2
		}
		return result
	}

	getEvenNumbers := func (limit int) []int{
		var result []int
		i:= 2
		for i <= limit{
			result = append(result, i)
			i = i + 2
		}
		return result
	}

	fmt.Println(getOddNumbers(10))
	fmt.Println(getEvenNumbers(10))
}


