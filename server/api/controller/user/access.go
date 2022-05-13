package user

import (
	"gs-reader-lite/server/api/controller"
	"gs-reader-lite/server/api/service/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (ctl *Controller) RegisterUser(req *gin.Context) {
	var reqData *RegisterReqData
	if err := ctl.BaseController.ValidateJson(req, &reqData); err != nil {
		return
	}

	userInfo, err := user.RegisterUserByPassword(reqData.Username, reqData.Password, reqData.Email, strconv.Itoa(reqData.Mobile))
	if err != nil {
		controller.JsonExit(req, 1, err.Error())
	} else {
		controller.JsonExit(req, 0, "success", userInfo)
	}

}

func (ctl *Controller) Login(req *gin.Context) {
	var reqData *LoginReqData
	if err := ctl.BaseController.ValidateJson(req, &reqData); err != nil {
		return
	}

	userInfo, err := user.Login(reqData.Password, reqData.Email, reqData.Mobile)
	if err != nil {
		controller.JsonExit(req, 1, "账号或密码不正确")
	} else {
		controller.JsonExit(req, 0, "success", userInfo)
	}
}
