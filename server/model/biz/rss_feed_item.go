package biz

import "github.com/gogf/gf/v2/os/gtime"

type RssFeedItemData struct {
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
	RssLink         string `gorm:"column:rssLink"`
	ChannelImageUrl string `gorm:"column:channelImageUrl"`
	ChannelTitle    string `gorm:"column:channelTitle"`
	Marked          int
	Sub             int
}
