package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JsonExit(c *gin.Context, err int, msg string, data ...interface{}) {
	c.JSON(http.StatusOK, gin.H{"error": err, "data": data, "msg": msg})
}

func JsonExitWithStatus(c *gin.Context, statusCode int, err int, msg string, data ...interface{}) {
	c.JSON(statusCode, gin.H{"error": err, "data": data, "msg": msg})
}
