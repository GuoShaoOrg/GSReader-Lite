package controller

import (
	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

func (ctl *BaseController) ValidateQuery(c *gin.Context, reqDataPointer interface{}) error {
	if err := c.ShouldBindQuery(reqDataPointer); err != nil {
		JsonExit(c, 1, err.Error())
		return err
	}
	return nil
}

func (ctl *BaseController) ValidateJson(c *gin.Context, reqDataPointer interface{}) error {
	if err := c.ShouldBindJSON(reqDataPointer); err != nil {
		JsonExit(c, 1, err.Error())
		return err
	}
	return nil
}