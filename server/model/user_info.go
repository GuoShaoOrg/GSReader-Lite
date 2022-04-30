package model

import "github.com/gogf/gf/v2/os/gtime"

type UserInfo struct {
	Uid        string      `gorm:"column:uid;primaryKey" json:"uid"`
	Password   string      `gorm:"column:password"    json:"password"`
	Email      string      `gorm:"column:email"       json:"email"`
	Mobile     string      `gorm:"column:mobile"      json:"mobile"`
	Username   string      `gorm:"column:username"    json:"username"`
	CreateDate *gtime.Time `gorm:"column:create_date" json:"createDate"`
	UpdateDate *gtime.Time `gorm:"column:update_date" json:"updateDate"`
}

func (UserInfo) TableName() string {
  return "user_info"
}