package main

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	run()
}
func run() {
	r := gin.Default()
	r.Use(cors.Default())
	r.LoadHTMLGlob("./static/*.html")
	r.GET("/greet", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello stranger from dockerhub!!!!!!",
		})
	})
	r.GET("/", func(ctx *gin.Context) {
		data := make(map[string]int)
		data["pid"] = os.Getpid()
		ctx.HTML(http.StatusOK, "index.html", data)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
