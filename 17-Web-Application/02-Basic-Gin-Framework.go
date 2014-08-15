package main

//run the command below before execute this file
//go get github.com/gin-gonic/gin
import (
	"github.com/gin-gonic/gin"
)

func IndexHandler(g *gin.Context) {

	g.String(200, "Home Page by Gin Framework")
}

func AboutHandler(g *gin.Context) {

	g.String(200, "About by Gin Framework")
}

func ContactHandler(g *gin.Context) {

	g.String(200, "Contact by Gin Framework")
}

func UserHandler(g *gin.Context) {

	// ->/user/{name}/{city}
	name := g.Params.ByName("name")
	city := g.Params.ByName("city")
	g.String(200, "%s, welcome to %s!", name, city)
}

func main() {

	g := gin.Default()

	g.GET("/", IndexHandler)
	g.GET("/about", AboutHandler)
	g.GET("/contact", ContactHandler)
	g.GET("/user/:name/:city", UserHandler)

	g.GET("/hello", func(g *gin.Context) {
		g.String(200, "Hello from Gin Framework")
	})

	g.Run(":3000")
}
