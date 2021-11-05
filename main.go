package main

import (
	"log"
	"net/http"
	"time"

	"cgin"
)

func onlyForV2() cgin.HandlerFunc {
	return func(c *cgin.Context) {
		t := time.Now()
		c.Fail(500, "Internal Server Error")
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := cgin.New()

	r.GET("/", func(c *cgin.Context) {
		c.HTML(http.StatusOK, "<h1>Hello cGin</h1>")
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *cgin.Context) {
			c.HTML(http.StatusOK, "<h1>Hello cGin </h1>")
		})

		v1.GET("/hello", func(c *cgin.Context) {
			c.String(http.StatusOK, "Hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	v2.Use(onlyForV2())
	{
		v2.GET("/hello/:name", func(c *cgin.Context) {
			c.String(http.StatusOK, "Hello %s, you're at %s\n", c.Param("name"), c.Path)
		})

		//v2.POST("/login", func(c *cgin.Context) {
		//	c.JSON(http.StatusOK, cgin.H{
		//		"username": c.PostForm("username"),
		//		"password": c.PostForm("password"),
		//	})
		//})
	}
	//r.GET("/hello", func(c *cgin.Context) {
	//	//
	//	c.String(http.StatusOK, "Hello %s, you`re at %s\n", c.Query("name"), c.Path)
	//})
	//
	//r.GET("/hello/:name", func(c *cgin.Context) {
	//	c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	//})
	//
	//r.GET("/assets/*filepath", func(c *cgin.Context) {
	//	c.JSON(http.StatusOK, cgin.H{"filepath": c.Param("filepath")})
	//})

	//r.POST("/login", func(c *cgin.Context) {
	//	c.JSON(http.StatusOK, cgin.H{
	//		"username": c.PostForm("username"),
	//		"password": c.PostForm("password"),
	//	})
	//})

	r.Run(":9999")
}
