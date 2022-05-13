package user

import (
	"context"
	"errors"
	"gs-reader-lite/server/component"
	"gs-reader-lite/server/model"
	"gs-reader-lite/server/model/biz"
)

func GetUserInfo(ctx context.Context, uid string) (userInfo biz.UserInfo, err error) {
	userInfoModel := model.UserInfo{}
	err = component.GetDatabase().Model(&model.UserInfo{}).Where("uid", uid).Find(&userInfoModel).Error
	if err != nil {
		component.Logger().Error(ctx, err)
		return userInfo, errors.New("发生了一些问题")
	}
	return
}
