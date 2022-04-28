package user

import (
	"gs-reader-lite/server/api/controller"
)

type Controller struct {
	controller.BaseController
}

var UsrCtl = &Controller{}

type RegisterReqData struct {
	Username       string `p:"username" v:"required|length:2,16#请输入昵称|昵称长度必须在2到16位"`
	Password       string `p:"password" v:"required|length:8,16#请输入密码|密码长度必须在8到16位"`
	PasswordVerify string `p:"passwordVerify" v:"required|length:8,16|same:password#请输入密码|密码长度必须在8到16位|两次密码不一致"`
	Email          string `p:"email" v:"required-without:mobile|email"`
	Mobile         string `p:"mobile" v:"required-without:email"`
}

type LoginReqData struct {
	Username string
	Password string `p:"password" v:"required#请输入密码"`
	Email    string `p:"email" v:"required-without:mobile|email#请输入账号|请输入正确的邮箱格式"`
	Mobile   int    `p:"mobile" v:"required-without:email#请输入正确的手机号"`
}
