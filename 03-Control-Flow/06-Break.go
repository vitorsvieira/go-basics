package main

import (
	"fmt"
)

func main() {
	OddNumbers()
}

func OddNumbers() {
	i := 1

	for {
		if i >= 25 {
			break
		}
		fmt.Println(i)
		i = i + 2
	}// 1 3 5 7 9 11 13 15 17 19 21 23
}
