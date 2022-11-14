package main

import (
	"net/http"

	gsgreeter "github.com/George-Spanos/gs_greeter"
	"github.com/gin-gonic/gin"
)

func main() {
	run()
}
func run() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": gsgreeter.Greet("George"),
		})
	})
	r.GET("/", func(ctx *gin.Context) {
		ctx.File("./static/index.html")
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
