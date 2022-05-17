package model

import "github.com/gogf/gf/v2/os/gtime"

type RssFeedItemFTS struct {
	Id            string      `gorm:"column:id"  json:"id"`
	ChannelId     string      `gorm:"column:channel_id"  json:"channelId"`
	Title         string      `gorm:"column:title"       json:"title"`
	Description   string      `gorm:"column:description" json:"description"`
	Link          string      `gorm:"column:link"        json:"link"`
	Date          *gtime.Time `gorm:"column:date"        json:"date"`
	Author        string      `gorm:"column:author"      json:"author"`
	InputDate     *gtime.Time `gorm:"column:input_date"  json:"inputDate"`
	Thumbnail     string      `gorm:"column:thumbnail"   json:"thumbnail"`
	Content       string      `gorm:"column:content"     json:"content"`
	ContentSP     string      `gorm:"column:content_sp"     json:"content_sp"`
	TitleSP       string      `gorm:"column:title_sp"       json:"title_sp"`
	DescriptionSP string      `gorm:"column:description_sp" json:"description_sp"`
}

func (RssFeedItemFTS) TableName() string {
	return "rss_feed_item_fts"
}

var (
	RFFTSIWithoutContentFieldSql = "fts.id, fts.channel_id, fts.title, fts.description, fts.link, fts.date, fts.author, fts.input_date, fts.thumbnail"
)
