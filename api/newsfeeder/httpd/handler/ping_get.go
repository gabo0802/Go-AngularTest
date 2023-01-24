package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Does as a handler function that returns an annonymous function so we can pass
// dependencies to it.
func PingGet() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"hello": "Found me",
		})
	}
}
