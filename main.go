package main

import (
	"fmt"
	"strings"

	"github.com/4dave/goapi/controllers"
	"github.com/4dave/goapi/initializers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectDB()
}

func main() {
	// mapit()
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/contact", controllers.ContactCreate)
	r.GET("/contacts", controllers.ContactList)
	r.GET("/contact/:id", controllers.ContactShow)
	r.PUT("/contact/:id", controllers.ContactUpdate)
	r.DELETE("/contact/:id", controllers.ContactDelete)
	r.Run()
}

func mapit() {
	address := "1463 Louisiana Ave S"
	// put + in place of spaces
	modified_address := strings.Replace(address, " ", "+", -1)

	url := "https://www.google.com/maps/place/" + modified_address

	fmt.Printf("URL: %s\n", url)
	// https://www.google.com/maps/search/?api=1&query=1200%20Pennsylvania%20Ave%20SE%2C%20Washington%2C%20District%20of%20Columbia%2C%2020003

}
