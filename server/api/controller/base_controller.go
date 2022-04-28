package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseController struct {
}

func (ctl *BaseController) Validate(c *gin.Context, reqDataPointer interface{}) {
	if err := c.ShouldBindJSON(reqDataPointer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
