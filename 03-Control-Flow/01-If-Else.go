package main

import (
	"fmt"
)

func main() {
	OpenOneIde(true)
	OpenTwoIde(true)
	OpenAllIde(false)
}

func OpenOneIde(isGolang bool) {

	if isGolang {
		fmt.Println("Opening Sublime Text 3")
	}
}

func OpenTwoIde(isGolang bool) {

	if otherIde := " and IntelliJ!"; isGolang {
		fmt.Println("Opening Sublime Text 3" + otherIde)
	}
}

func OpenAllIde(isGolang bool) {

	if isGolang {
		fmt.Println("Opening all IDEs installed")
	} else {
		fmt.Println("Not coding now...")
	}
}
