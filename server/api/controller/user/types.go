package user

import (
	"gs-reader-lite/server/api/controller"
)

type Controller struct {
	controller.BaseController
}

var UsrCtl = &Controller{}

type RegisterReqData struct {
	Username       string `json:"username"`
	Password       string `json:"password" binding:"required,min=8,max=16"`
	PasswordVerify string `json:"passwordVerify" binding:"required,min=8,max=16,eqfield=Password"`
	Email          string `json:"email" binding:"required_without=Mobile"`
	Mobile         int    `json:"mobile" binding:"required_without=Email"`
}

type LoginReqData struct {
	Username string
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required_without=Mobile"`
	Mobile   int    `json:"mobile" binding:"required_without=Email"`
}
