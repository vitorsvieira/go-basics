package main

import (
	"fmt"
	"log"
	"net/http"
)

func IndexHandler(res http.ResponseWriter, req *http.Request) {

	fmt.Fprintln(res, "<h1>Home Page</h1>")

}

func AboutHandler(res http.ResponseWriter, req *http.Request) {

	fmt.Fprintln(res, "<h1>About</h1>")

}

func ContactHandler(res http.ResponseWriter, req *http.Request) {

	fmt.Fprintln(res, "<h1>Contact</h1>")

}

func main() {

	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/about", AboutHandler)
	http.HandleFunc("/contact", ContactHandler)

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal("ERROR :-> ListenAndServe: ", err)
	}
}
