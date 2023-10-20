package main

import (
	"github.com/4dave/goapi/controllers"
	"github.com/4dave/goapi/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()
	r.POST("/contact", controllers.ContactCreate)
	r.GET("/contacts", controllers.ContactList)
	r.GET("/contact/:id", controllers.ContactShow)
	r.PUT("/contact/:id", controllers.ContactUpdate)
	r.DELETE("/contact/:id", controllers.ContactDelete)
	r.Run()
}
