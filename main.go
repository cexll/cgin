package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"cgin"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func onlyForV2() cgin.HandlerFunc {
	return func(c *cgin.Context) {
		t := time.Now()
		c.Fail(500, "Internal Server Error")
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := cgin.New()
	r.Use(cgin.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	stu1 := &student{Name: "Cexll", Age: 20}
	stu2 := &student{Name: "TheShy", Age: 22}
	r.GET("/", func(c *cgin.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})

	r.GET("/students", func(c *cgin.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", cgin.H{
			"title":  "cgin",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.GET("/date", func(c *cgin.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", cgin.H{
			"title": "cgin",
			"now":   time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC),
		})
	})
	r.Run(":9999")

	//r.GET("/", func(c *cgin.Context) {
	//	c.HTML(http.StatusOK, "<h1>Hello cGin</h1>")
	//})
	//
	//v1 := r.Group("/v1")
	//{
	//	v1.GET("/", func(c *cgin.Context) {
	//		c.HTML(http.StatusOK, "<h1>Hello cGin </h1>")
	//	})
	//
	//	v1.GET("/hello", func(c *cgin.Context) {
	//		c.String(http.StatusOK, "Hello %s, you're at %s\n", c.Query("name"), c.Path)
	//	})
	//}
	//v2 := r.Group("/v2")
	//v2.Use(onlyForV2())
	//{
	//	v2.GET("/hello/:name", func(c *cgin.Context) {
	//		c.String(http.StatusOK, "Hello %s, you're at %s\n", c.Param("name"), c.Path)
	//	})
	//	r.Static("/assets", "./static")

	//v2.POST("/login", func(c *cgin.Context) {
	//	c.JSON(http.StatusOK, cgin.H{
	//		"username": c.PostForm("username"),
	//		"password": c.PostForm("password"),
	//	})
	//})
	//}
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

}
