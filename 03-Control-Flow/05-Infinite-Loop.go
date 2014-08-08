package main

import (
	"fmt"
)

func main() {
	OddNumbers()
}

func OddNumbers(){
	i := 1

	for {
		fmt.Println(i)

		i = i + 2
	}
}
