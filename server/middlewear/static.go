package middlewear

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)


func StaticRedirect() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.RequestURI == "/" || strings.HasPrefix(c.Request.RequestURI, "/js") || strings.HasPrefix(c.Request.RequestURI, "/css") {
			c.Request.RequestURI = "/view" + c.Request.RequestURI
			c.Redirect(http.StatusMovedPermanently, c.Request.RequestURI)
		}
	}
}