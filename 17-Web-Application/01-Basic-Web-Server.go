package main

import (
	"fmt"
	"log"
	"net/http"
)

func IndexHandler(res http.ResponseWriter, req *http.Request) {

	fmt.Fprintln(res, "<h1>Home Page</h1><h3>From a Go Web Server</h3>")
}

func AboutHandler(res http.ResponseWriter, req *http.Request) {

	fmt.Fprintln(res, "<h1>About</h1><h3>From a Go Web Server</h3>")
}

func ContactHandler(res http.ResponseWriter, req *http.Request) {

	fmt.Fprintln(res, "<h1>Contact</h1><h3>From a Go Web Server</h3>")
}

func UserHandler(res http.ResponseWriter, req *http.Request) {

	// .../user?name={name}&city={city}
	name := req.FormValue("name")
	city := req.FormValue("city")

	fmt.Fprintf(res, "<h1>%s, welcome to %s!</h1>", name, city)
}

func main() {

	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/about", AboutHandler)
	http.HandleFunc("/contact", ContactHandler)
	http.HandleFunc("/user", UserHandler)
	http.HandleFunc("/hello", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "<h1>Hello</h1><h3>From a Go Web Server</h3>")
	})

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal("ERROR :-> ListenAndServe: ", err)
	}
}
