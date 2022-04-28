package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func JsonExist(c *gin.Context, err int, msg string, data ...interface{}) {
	c.JSON(http.StatusBadRequest, gin.H{"error": err, "data": data, "msg": msg})
}

func JsonExistWithStatus(c *gin.Context, statusCode int, err int, msg string, data ...interface{}) {
	c.JSON(statusCode, gin.H{"error": err, "data": data, "msg": msg})
}
