package http

import "github.com/gin-gonic/gin"

// Route is Gin's route which contains url regexp.
type Route struct {
	URL            string
	DefaultGetResp interface{}
}

// Routes is the route list.
var Routes = []Route{
	{URL: "/posts", DefaultGetResp: gin.H{
		"albmu_id": "123",
	}},
}
