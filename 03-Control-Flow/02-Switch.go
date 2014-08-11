package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Program Name: ", GetProgram("doc"))
	fmt.Println("Program Name: ", GetProgram("xlsx"))
	fmt.Println("Program Name: ", GetProgram2("txt"))
	fmt.Println("Program Name: ", GetProgram2(""))

	fmt.Println("Type Name: ", SwitchByType(1234))
	fmt.Println("Type Name: ", SwitchByType("Go"))
	fmt.Println("Type Name: ", SwitchByType(0.5))
	fmt.Println("Type Name: ", SwitchByType(10i))
}

func GetProgram(extension string) (program string) {

	switch extension {
	case "pdf":
		program = "Adobe Reader"
	case "xls", "xlsx":
		program = "Microsoft Excel"
	case "doc", "docx":
		program = "Microsoft Word"
	default:
		program = "Go"
	}

	return
}

func GetProgram2(extension string) (program string) {

	switch {
	case extension == "pdf":
		program = "Adobe Reader"
	case extension == "xls", extension == "xlsx":
		program = "Microsoft Excel"
	case extension == "doc", extension == "docx":
		program = "Microsoft Word"
	case len(extension) == 0:
		program = "No extension"
	default:
		program = "Go"
	}

	return
}

//Go to folder 09-Interface for basics and examples about Interface
func SwitchByType(x interface{}) (typeName string) {

	switch t := x.(type) {
	case int:
		typeName = reflect.TypeOf(t).String()
	case string:
		typeName = reflect.TypeOf(t).String()
	case float64:
		typeName = reflect.TypeOf(t).String()
	default:
		typeName = "Unknow Type"
	}

	return
}
