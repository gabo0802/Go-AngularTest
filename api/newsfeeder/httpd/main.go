package main

//http://localhost:8080/ping

import (
	"github.com/gabo0802/Go-AngularTest/api/newsfeeder/httpd/handler"
	"github.com/gabo0802/Go-AngularTest/api/newsfeeder/platform/newsfeed"
	"github.com/gin-gonic/gin"
)

func main() {

	feed := newsfeed.New()
	r := gin.Default()

	r.GET("/ping", handler.PingGet())
	r.GET("/newsfeed", handler.NewsFeedGet(feed))
	r.POST("/newsfeed", handler.NewsFeedPost(feed))

	r.Run()

}
