package main

import (
	"net/http"

	"cgin"
)

func main() {
	r := cgin.Default()

	r.GET("/", func(c *cgin.Context) {
		c.String(http.StatusOK, "hello Cexll\n")
	})

	r.GET("/panic", func(c *cgin.Context) {
		names := []string{"cexll"}
		c.String(http.StatusOK, names[100])
	})
	r.Run(":9999")
}
