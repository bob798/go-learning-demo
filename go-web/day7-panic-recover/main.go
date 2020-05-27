package main

import (
	"github.com/bob798/go-learning-demo/go-web/day7-panic-recover/gee"
	"net/http"
)


func main() {
	r := gee.New()

	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK,"Hello Bob")
	})

	r.Use(gee.Recovery())
	r.GET("/panic", func(c *gee.Context) {
		names := []string{"bob"}
		c.String(http.StatusOK,names[10])
	})

	r.Run(":9999")
}
