package main

import (
	"github.com/gabo0802/Go-AngularTest/api/newsfeeder/httpd/handler"
	"github.com/gabo0802/Go-AngularTest/api/newsfeeder/platform/newsfeed"
	"github.com/gin-gonic/gin"
)

func main() {

	feed := newsfeed.New()
	r := gin.Default()

	//This is the magic part that links our api to the angular app
	api := r.Group("/api")
	{
		api.GET("/ping", handler.PingGet())
		api.GET("/newsfeed", handler.NewsFeedGet(feed))
		api.POST("/newsfeed", handler.NewsFeedPost(feed))
	}

	r.Run("0.0.0.0:5000")

}
