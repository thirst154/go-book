package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thirst154/go-book/controllers"
	"github.com/thirst154/go-book/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.LoadHTMLGlob("templates/**/*")

	pages := r.Group("/pages")

	pages.GET("/", controllers.GetIndex)
	r.GET("/", controllers.GetIndex)

	pages.GET("/login", controllers.GetLoginComponent)
	pages.GET("/create-account", controllers.GetCreateAccountComponent)
	pages.GET("/greeting", controllers.GetGreetingComponent)
	pages.GET("/posts", controllers.GetPosts)


	api := r.Group("/api")

	api.POST("/login", controllers.UserLogin)

	r.Run(":8080")
}
