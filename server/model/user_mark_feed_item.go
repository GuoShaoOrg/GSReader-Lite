package model

import "github.com/gogf/gf/v2/os/gtime"

type UserMarkFeedItem struct {
	UserId        string      `gorm:"column:user_id"         json:"userId"`        
	ChannelItemId string      `gorm:"column:channel_item_id" json:"channelItemId"` 
	InputTime     *gtime.Time `gorm:"column:input_time"      json:"inputTime"`     
	Status        int         `gorm:"column:status"          json:"status"`        
}

func (UserMarkFeedItem) TableName() string {
	return "user_mark_feed_item"
}
