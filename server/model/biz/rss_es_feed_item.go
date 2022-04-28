package biz

import "github.com/gogf/gf/v2/os/gtime"

type RssFeedItemESData struct {
	Id              string
	ChannelId       string
	Title           string
	Description     string
	Content         string
	Thumbnail       string
	Link            string
	Date            *gtime.Time
	Author          string
	InputDate       *gtime.Time
	ChannelImageUrl string
	ChannelTitle    string
	ChannelLink     string
}
