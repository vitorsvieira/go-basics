package main

import (
	"fmt"
)

func main() {
	OddNumbers(20)
}

func OddNumbers(loops int) {
	i := 0

	for {
		if i >= loops {
			break
		}
		if i%2 == 0 {
			i++
			continue
		}

		fmt.Println(i)
		i++
	}// 1 3 5 7 9 11 13 15 17 19
}
