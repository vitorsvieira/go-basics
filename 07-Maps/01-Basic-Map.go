package main

import (
	"fmt"
)

func main() {

	m := CreateMap()

	fmt.Println("FindMapValue: ", FindMapValue(m, "Bob"))
	fmt.Println("FindMapValue: ", FindMapValue(m, "John"))

	fmt.Println("FindMapValue2: ", FindMapValue(m, "John"))
	fmt.Println("FindMapValue2: ", FindMapValue(m, "Bob"))

	result, i := AddMapValue(m, "NewBob", "NewIDE")
	fmt.Println("Added: ", result, "Total: ", i)

	fmt.Println("Update: ", UpdateMapValue(m, "Bob", "Vim"))
	fmt.Println("Update: ", UpdateMapValue(m, "John", "Vim"))

	fmt.Println("Delete-> Map Length Now: ", DeleteMapValue(m, "Bob"))

	fmt.Println("Bob Exist? ", NameExist(m, "Bob"))
	fmt.Println("John Exist? ", NameExist(m, "John"))
	fmt.Println("NewBob Exist? ", NameExist(m, "NewBob"))

}

func Run() {
	fmt.Println("Run!")
}

func CreateMap() map[string]string {

	var m map[string]string
	m = make(map[string]string)

	m["John"] = "LiteIDE"
	m["Bob"] = "IntelliJ"

	return m
}

func CreateMap2() map[string]string {

	//shorthand map declaration
	m := map[string]string{
		"John": "LiteIDE",
		"Bob":  "IntelliJ",
	}
	return m
}

func FindMapValue(m map[string]string, name string) (ide string) {

	return m[name]
}

func AddMapValue(m map[string]string, name, newIde string) (map[string]string, int) {

	m[name] = newIde

	return m, len(m)
}

func UpdateMapValue(m map[string]string, name, newIde string) map[string]string {

	m[name] = newIde

	return m
}

func DeleteMapValue(m map[string]string, name string) (mapLength int) {

	delete(m, name)

	return len(m)
}

func NameExist(m map[string]string, name string) (exists bool) {

	if _, exists = m[name]; exists {
		return exists
	}
	return
}
