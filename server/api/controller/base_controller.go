package controller

import (
	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

func (ctl *BaseController) Validate(c *gin.Context, reqDataPointer interface{}) error {
	if err := c.ShouldBindJSON(reqDataPointer); err != nil {
		JsonExist(c, 1, err.Error())
		return err
	}
	return nil
}
