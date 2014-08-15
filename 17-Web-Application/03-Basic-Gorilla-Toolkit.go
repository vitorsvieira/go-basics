package main

//run the command below before execute this file
//go get github.com/gorilla/mux
import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func IndexHandler(res http.ResponseWriter, req *http.Request) {

	fmt.Fprintln(res, "<h1>Home Page</h1><h3>Using Gorilla Toolkit</h3>")
}

func AboutHandler(res http.ResponseWriter, req *http.Request) {

	fmt.Fprintln(res, "<h1>About</h1><h3>Using Gorilla Toolkit</h3>")
}

func ContactHandler(res http.ResponseWriter, req *http.Request) {

	fmt.Fprintln(res, "<h1>Contact</h1><h3>Using Gorilla Toolkit</h3>")
}

func UserHandler(res http.ResponseWriter, req *http.Request) {

	// ->/user/{name}/{city}
	params := mux.Vars(req)

	fmt.Fprintf(res, "<h1>%s, welcome to %s!</h1><h3>Using Gorilla Toolkit</h3>", params["name"], params["city"])
}

func main() {

	gorilla := mux.NewRouter()

	gorilla.HandleFunc("/", IndexHandler)
	gorilla.HandleFunc("/about", AboutHandler)
	gorilla.HandleFunc("/contact", ContactHandler)
	gorilla.HandleFunc("/user/{name}/{city}", UserHandler)
	gorilla.HandleFunc("/hello", func(res http.ResponseWriter, req *http.Request) {
			fmt.Fprintf(res, "<h1>Hello</h1><h3>From Gorilla Mux</h3>")
	})

	http.ListenAndServe(":3000", gorilla)

}
