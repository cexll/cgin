package lin

import (
	"net/http"
	"testing"
)

func TestStart(t *testing.T) {
	r := New()
	r.GET("/", func(c *Context) {
		c.String(http.StatusOK, "Hello Lin!")
	})
	r.GET("/hello", func(c *Context) {
		c.String(http.StatusOK, "Hello %s, you' re at %s\n!", c.Query("name"), c.Path)	
	})
	v1 := r.Group("/v1")
	v1.GET("/", func (c *Context) {
		c.String(http.StatusOK, "hello")
	})
	r.Run(":9090")
}
