# lin

# Install 

```
go get -u github.com/go-ll/lin
```

# Run

```go
r := New()
r.GET("/", func(c *Context) {
    c.String(http.StatusOK, "Hello Lin!")
})
r.Run(":9090")
```
