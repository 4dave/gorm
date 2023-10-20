package main

import (
	"github.com/4dave/goapi/initializers"
	"github.com/4dave/goapi/models"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Contact{})
}
