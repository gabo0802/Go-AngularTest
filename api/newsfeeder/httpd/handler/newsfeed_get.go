package handler

import (
	"net/http"

	"github.com/gabo0802/Go-AngularTest/api/newsfeeder/platform/newsfeed"
	"github.com/gin-gonic/gin"
)

// Does as a handler function that returns an annonymous function so we can pass
// dependencies to it.
func NewsFeedGet(feed newsfeed.Getter) gin.HandlerFunc {

	return func(c *gin.Context) {
		results := feed.GetAll()
		c.JSON(http.StatusOK, results)
	}
}
