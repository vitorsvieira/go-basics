package main

import (
	"fmt"
)

func main() {
	Knit(5)
}

func Knit(minutes int) {

	i := 0
	for i <= minutes {
		switch i {
		case 0:
			fmt.Println("Starting to knit!")
		case 1:
			fmt.Println("In", i, "minute you knitted", i*2, "inches or your new scarf!")
		default:
			fmt.Println("In", i, "minutes you knitted", i*2, "inches or your new scarf!")
		}
		i++
	}
}
