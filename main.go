package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	run()
}
func run() {
	r := gin.Default()
	r.GET("/greet", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello stranger",
		})
	})
	r.GET("/", func(ctx *gin.Context) {
		ctx.File("./static/index.html")
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
