package main

import (
	"fmt"
)

func main() {

	var x int

	for {
		num, err := fmt.Scanf("%d", &x)

		if err == nil {
			fmt.Println(num, err)
			break
		}
	}

	fmt.Println("x = ", x)
}
