package model

import "github.com/gogf/gf/v2/os/gtime"

type UserSubChannel struct {
	UserId    string     `gorm:"column:user_id"    json:"userId"`    
	ChannelId string     `gorm:"column:channel_id" json:"channelId"` 
	InputTime *gtime.Time `gorm:"column:input_time" json:"inputTime"` 
	Status    int        `gorm:"column:status"     json:"status"`    
}

func (UserSubChannel) TableName() string {
	return "user_sub_channel"
}