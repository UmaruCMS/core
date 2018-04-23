package http

import (
	"github.com/gin-gonic/gin"
)

func defaultFunc(defaultGetResp interface{}) func(*gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, defaultGetResp)
	}
}

// Serve starts a DEMO http server.
func Serve() {
	r := gin.Default()
	for _, route := range Routes {
		r.GET(route.URL, defaultFunc(route.DefaultGetResp))
	}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
