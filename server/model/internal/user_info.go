// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserInfo is the golang structure for table user_info.
type UserInfo struct {
	Uid        string      `orm:"uid,primary" json:"uid"`        //
	Password   string      `orm:"password"    json:"password"`   //
	Email      string      `orm:"email"       json:"email"`      //
	Mobile     string      `orm:"mobile"      json:"mobile"`     //
	Username   string      `orm:"username"    json:"username"`   //
	CreateDate *gtime.Time `orm:"create_date" json:"createDate"` //
	UpdateDate *gtime.Time `orm:"update_date" json:"updateDate"` //
}