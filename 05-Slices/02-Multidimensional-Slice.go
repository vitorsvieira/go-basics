package main

import (
	"fmt"
)

func main() {

	x := [2][2]int{
		{2,4},
		{6,8},
	}

	y := [2][2]int{
		{5,10},
		{15,20},
	}

	fmt.Println(x, y)

	z := [2][2]int{
		{x[0][0] * y[0][0], x[0][1] * y[0][1]},
		{x[1][0] * y[1][0], x[1][1] * y[1][1]},
	}

	fmt.Println(z)
}
