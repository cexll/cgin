package main

import (
	"net/http"

	"cgin"
)

func main() {
	r := cgin.New()
	r.GET("/", func(c *cgin.Context) {
		c.HTML(http.StatusOK, "<h1>Hello cGin</h1>")
	})

	r.GET("/hello", func(c *cgin.Context) {
		//
		c.String(http.StatusOK, "Hello %s, you`re at %s\n", c.Query("name"), c.Path)
	})
	r.POST("/login", func(c *cgin.Context) {
		c.JSON(http.StatusOK, cgin.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
