package pages

import "gs-reader-lite/server/api/controller"

type Controller struct {
	controller.BaseController
}

var PagesCtl = &Controller{}

type HomeReqData struct {
	UserId string `form:"userId" binding:"required"`
}

type FeedItemType struct {
	Id              string
	ChannelId       string
	Title           string
	Description     string
	Content         string
	Thumbnail       string
	Link            string
	Date            string
	Author          string
	InputDate       string
	RssLink         string
	ChannelImageUrl string
	ChannelTitle    string
	Marked          int
	Sub             int
}
