package controllers

import "github.com/gin-gonic/gin"

func GetIndex(c *gin.Context) {
	c.HTML(200, "index.tmpl", gin.H{
		"message": "Hello, World!",
		"Content": "greeting.tmpl"})

}

func GetLoginComponent(c *gin.Context) {
	c.HTML(200, "login.tmpl", gin.H{})
}

func GetGreetingComponent(c *gin.Context) {
	c.HTML(200, "greeting.tmpl", gin.H{})
}

func GetCreateAccountComponent(c *gin.Context) {
	c.HTML(200, "create-account.tmpl", gin.H{})
}
