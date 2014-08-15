package main

//run the command below before execute this file
//go get github.com/go-martini/martini
import (
	"fmt"
	"github.com/go-martini/martini"
	"net/http"
)

func IndexHandler(res http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(res, "<h1>Home Page</h1><h3>Using Martini Framework</h3>")
}

func AboutHandler(res http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(res, "<h1>About</h1><h3>Using Martini Framework</h3>")
}

func ContactHandler(res http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(res, "<h1>Contact</h1><h3>Using Martini Framework</h3>")
}

func UserHandler(params martini.Params) string {

	// ->/user/{name}/{city}
	return fmt.Sprintf("<h1>%s, welcome to %s!</h1>", params["name"], params["city"])
}

func main() {

	m := martini.Classic()

	m.Get("/", IndexHandler)
	m.Get("/about", AboutHandler)
	m.Get("/contact", ContactHandler)
	m.Get("/user/:name/:city", UserHandler)
	m.Get("/hello", func() string {
		return fmt.Sprintf("<h1>Hello</h1><h2>from Martini Framework</h3>")
	})

	// Run() listens on os.GetEnv("PORT") or 3000 by default.
	m.Run()
}
