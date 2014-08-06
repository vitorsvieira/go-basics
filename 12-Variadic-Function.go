package main

import (
	"fmt"
)

type Plate struct {
	pasta     string
	vegetable string
}

func CookLunch(pasta string, vegetable ...string) string {
	//len(vegetable)
	return pasta + " with " + vegetable[0]

}

func CookDinner(pasta string, vegetable ...string) string {
	//len(vegetable)
	return pasta + " with " + vegetable[1]

}

func main() {
	plate := Plate{"Fettuccine", "Pumpkin"}

	fmt.Println("For Lunch: ", CookLunch(plate.pasta, plate.vegetable, "Tomato"))
	fmt.Println("For Dinner: ", CookDinner(plate.pasta, plate.vegetable, "Tomato"))
}
