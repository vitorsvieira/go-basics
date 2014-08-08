package main

import (
	"fmt"
)

func main() {
	Run(8)
}

func Run(laps int) {

	for i := 0; i <= laps; i++ {
		switch i {
		case 0:
			fmt.Println("Starting to run!")
		case 1:
			fmt.Println("In", i, "lap you ran", i*850, " meters.")
		default:
			fmt.Println("In", i, "laps you ran", i*850, " meters.")
		}
	}
}
