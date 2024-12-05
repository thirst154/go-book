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

type Message struct {
	Title   string
	Content string
}

type Messages []Message

func GetPosts(c *gin.Context) {

	content := Messages{
		Message{Title: "Hello World", Content: "Hello Again"},
		Message{Title: "Hello Tom!", Content: "123123123123123123123123123123123123123123123123"},
		Message{Title: "Test Test", Content: "Hello"},
		Message{Title: "Hello World", Content: "Hello Again AND AgainAgainAgainAgainAgainAgainAgainAgain"},
	}

	c.HTML(200, "posts.tmpl", gin.H{"Messages": content})
}
