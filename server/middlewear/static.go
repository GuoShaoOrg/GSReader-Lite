package middlewear

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func StaticRedirect() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.RequestURI == "/" {
			c.Request.RequestURI = "/view" + c.Request.RequestURI
			c.Redirect(http.StatusMovedPermanently, c.Request.RequestURI)
		}
	}
}