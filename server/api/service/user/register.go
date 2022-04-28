package user

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/guid"
	"gs-reader-lite/server/component"
	middleware "gs-reader-lite/server/middlewear"

	"gs-reader-lite/server/model"
	"gs-reader-lite/server/model/biz"
)

func RegisterUserByPassword(username, password, email, mobile string) (*biz.UserInfo, error) {
	if email == "" {
		result, _ := component.GetDatabase().Model("user_info").Where("mobile", mobile).All()
		if result.Size() != 0 {
			return nil, errors.New("该手机号已经注册")
		}
	} else if mobile == "" {
		result, _ := component.GetDatabase().Model("user_info").Where("email", email).All()
		if result.Size() != 0 {
			return nil, errors.New("该邮箱已经注册")
		}
	}
	cryptoPwd, _ := gmd5.Encrypt(password)
	createDate := gtime.Now()
	userInfoModel := model.UserInfo{
		Username:   username,
		Password:   cryptoPwd,
		Email:      email,
		Mobile:     mobile,
		CreateDate: createDate,
		UpdateDate: createDate,
		Uid:        guid.S(),
	}
	ctx := context.Background()
	_, err := component.GetDatabase().Insert(ctx, "user_info", userInfoModel)
	if err != nil {
		return nil, errors.New("发生了一些问题")
	}

	bizUseInfo := biz.UserInfo{
		Uid:        userInfoModel.Uid,
		Email:      userInfoModel.Email,
		Mobile:     userInfoModel.Mobile,
		Username:   userInfoModel.Username,
		CreateDate: userInfoModel.CreateDate.String(),
		UpdateDate: userInfoModel.UpdateDate.String(),
	}

	tokenMode := middleware.TokenModel{
		UserId:         bizUseInfo.Uid,
		UserName:       bizUseInfo.Username,
		Mobile:         bizUseInfo.Mobile,
		CreateDate:     bizUseInfo.CreateDate,
		UpdateDateTime: bizUseInfo.UpdateDate,
	}
	token, err := middleware.CreateToken(tokenMode)
	if err != nil {
		return nil, errors.New("发生了一些问题")
	}
	bizUseInfo.Token = token

	return &bizUseInfo, nil
}

func Login(password, email string, mobile int) (*biz.UserInfo, error) {
	cryptoPwd, _ := gmd5.Encrypt(password)
	userInfoModel := model.UserInfo{}

	if email == "" {
		err := component.GetDatabase().Model("user_info").Where("mobile", mobile).Where("password", cryptoPwd).Scan(&userInfoModel)
		if err != nil {
			return nil, err
		}
	} else if mobile == 0 {
		err := component.GetDatabase().Model("user_info").Where("email", email).Where("password", cryptoPwd).Scan(&userInfoModel)
		if err != nil {
			return nil, err
		}
	}

	bizUseInfo := biz.UserInfo{
		Uid:        userInfoModel.Uid,
		Email:      userInfoModel.Email,
		Mobile:     userInfoModel.Mobile,
		Username:   userInfoModel.Username,
		CreateDate: userInfoModel.CreateDate.String(),
		UpdateDate: userInfoModel.UpdateDate.String(),
	}

	tokenMode := middleware.TokenModel{
		UserId:         bizUseInfo.Uid,
		UserName:       bizUseInfo.Username,
		Mobile:         bizUseInfo.Mobile,
		CreateDate:     bizUseInfo.CreateDate,
		UpdateDateTime: bizUseInfo.UpdateDate,
	}
	token, err := middleware.CreateToken(tokenMode)
	if err != nil {
		return nil, errors.New("发生了一些问题")
	}
	bizUseInfo.Token = token

	return &bizUseInfo, nil

}
