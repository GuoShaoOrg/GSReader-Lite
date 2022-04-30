package model

import "github.com/gogf/gf/v2/os/gtime"

type RssFeedTag struct {
	Name      string      `gorm:"column:name" json:"name"`
	ChannelId string      `gorm:"column:channel_id"  json:"channelId"`
	Title     string      `gorm:"column:title"       json:"title"`
	Date      *gtime.Time `gorm:"column:date"        json:"date"`
}

func (RssFeedTag) TableName() string {
  return "rss_feed_tag"
}