package pages

import "gs-reader-lite/server/api/controller"

type Controller struct {
	controller.BaseController
}

var PagesCtl = &Controller{}

type HomeReqData struct {
	UserId string `form:"userId" binding:"required"`
}
