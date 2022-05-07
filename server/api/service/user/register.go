package user

import (
	"context"
	"errors"
	"gs-reader-lite/server/component"
	middleware "gs-reader-lite/server/middlewear"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/guid"

	"gs-reader-lite/server/model"
	"gs-reader-lite/server/model/biz"
)

func RegisterUserByPassword(username, password, email, mobile string) (*biz.UserInfo, error) {
	ctx := context.Background()
	checkUserInfoModel := model.UserInfo{}
	if email == "" {
		result := component.GetDatabase().Model(&model.UserInfo{}).Where("mobile", mobile).Find(&checkUserInfoModel)
		if result.Error != nil {
			component.Logger().Error(ctx, result.Error)
			return nil, errors.New("发生了一些问题")
		} else if checkUserInfoModel.Uid != "" {
			return nil, errors.New("该手机号已经注册")
		} else {
			component.Logger().Error(ctx, result.Error)
		}
	} else if mobile == "" {
		result := component.GetDatabase().Model(&model.UserInfo{}).Where("email", email).Find(&checkUserInfoModel)
		if result.Error != nil {
			component.Logger().Error(ctx, result.Error)
			return nil, errors.New("发生了一些问题")
		} else if checkUserInfoModel.Uid != "" {
			return nil, errors.New("该邮箱已经注册")
		} else {
			component.Logger().Error(ctx, result.Error)
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
	result := component.GetDatabase().Create(&userInfoModel)
	if result.Error != nil {
		component.Logger().Error(ctx, result.Error)
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
		component.Logger().Error(ctx, err.Error())
		return nil, errors.New("发生了一些问题")
	}
	bizUseInfo.Token = token

	return &bizUseInfo, nil
}

func Login(password, email string, mobile int) (*biz.UserInfo, error) {
	cryptoPwd, _ := gmd5.Encrypt(password)
	userInfoModel := model.UserInfo{}
	ctx := context.Background()

	if email == "" {
		result := component.GetDatabase().Model(&model.UserInfo{}).Where("mobile", mobile).Where("password", cryptoPwd).Scan(&userInfoModel)
		if result.Error != nil {
			component.Logger().Error(ctx, result.Error)
			return nil, result.Error
		}
	} else if mobile == 0 {
		result := component.GetDatabase().Model(&model.UserInfo{}).Where("email", email).Where("password", cryptoPwd).Scan(&userInfoModel)
		if result.Error != nil {
			component.Logger().Error(ctx, result.Error)
			return nil, result.Error
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
		component.Logger().Error(ctx, err.Error())
		return nil, errors.New("发生了一些问题")
	}
	bizUseInfo.Token = token

	return &bizUseInfo, nil

}
