package handler

import (
	"net/http"

	"github.com/gabo0802/Go-AngularTest/newsfeeder/platform/newsfeed"
	"github.com/gin-gonic/gin"
)

type newsfeedPostRequest struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

// Does as a handler function that returns an annonymous function so we can pass
// dependencies to it.
func NewsFeedPost(feed newsfeed.Adder) gin.HandlerFunc {

	return func(c *gin.Context) {

		requestBody := newsfeedPostRequest{}
		c.Bind(&requestBody)

		item := newsfeed.Item{
			Title: requestBody.Title,
			Post:  requestBody.Post,
		}
		feed.Add(item)

		c.Status(http.StatusNoContent)
	}
}
