package feed

import (
	"context"
	"gs-reader-lite/server/api/controller"
	"gs-reader-lite/server/api/service/feed"

	"github.com/gin-gonic/gin"
)

func (ctl *Controller) GetFeedChannelByTag(req *gin.Context) {
	var reqData *ChannelReqData
	if err := ctl.BaseController.ValidateQuery(req, &reqData); err != nil {
		return
	}
	if reqData.Size == 0 {
		reqData.Size = 10
	}
	tagList := feed.GetFeedChannelByTag(context.Background(), reqData.Start, reqData.Size, reqData.Name)
	controller.JsonExit(req, 0, "success", tagList)
}

func (ctl *Controller) SubChannelByUserIdAndChannelId(req *gin.Context) {
	var reqData *SubChannelByUserIdAndChannelIdReqData
	if err := ctl.BaseController.ValidateJson(req, &reqData); err != nil {
		return
	}

	err := feed.SubChannelByUserIdAndChannelId(context.Background(), reqData.UserID, reqData.ChannelId)

	if err != nil {
		controller.JsonExit(req, 1, "failed", "sub channel with user id failed")
	} else {
		controller.JsonExit(req, 0, "success", nil)
	}
}

func (ctl *Controller) GetSubFeedChannelByUserId(req *gin.Context) {
	var reqData *GetSubChannelByUserIdReqData
	if err := ctl.BaseController.ValidateQuery(req, &reqData); err != nil {
		return
	}
	if reqData.Size == 0 {
		reqData.Size = 10
	}
	tagList := feed.GetSubChannelListByUserId(context.Background(), reqData.UserID, reqData.Start, reqData.Size)
	controller.JsonExit(req, 0, "success", tagList)
}

func (ctl *Controller) GetFeedChannelCatalogListByTag(req *gin.Context) {
	var reqData *GetChannelCatalogListByTagReqData
	if err := ctl.BaseController.ValidateQuery(req, &reqData); err != nil {
		return
	}
	if reqData.Size == 0 {
		reqData.Size = 10
	}
	tagList := feed.GetFeedChannelCatalogListByTag(context.Background(), reqData.TagName, reqData.Start, reqData.Size)
	controller.JsonExit(req, 0, "success", tagList)
}

func (ctl *Controller) GetFeedChannelCatalogListByUserId(req *gin.Context) {
	var reqData *GetChannelCatalogListByUserIdReqData
	if err := ctl.BaseController.ValidateQuery(req, &reqData); err != nil {
		return
	}
	if reqData.Size == 0 {
		reqData.Size = 10
	}
	tagList := feed.GetFeedChannelCatalogListByUserId(context.Background(), reqData.UserId, reqData.Start, reqData.Size)
	controller.JsonExit(req, 0, "success", tagList)
}

func (ctl *Controller) GetFeedChannelInfoByChannelAndUserId(req *gin.Context) {
	var reqData *GetChannelInfoByChannelUserIdReqData
	if err := ctl.BaseController.ValidateQuery(req, &reqData); err != nil {
		return
	}

	tagList := feed.GetChannelInfoByChannelAndUserId(context.Background(), reqData.UserID, reqData.ChannelId)
	controller.JsonExit(req, 0, "success", tagList)
}

func (ctl *Controller) AddFeedChannelByLink(req *gin.Context) {
	var reqData *AddFeedChannelByLinkReqData
	if err := ctl.BaseController.ValidateJson(req, &reqData); err != nil {
		return
	}

	var err error
	err = feed.AddFeedChannelByLink(context.Background(), reqData.UserID, reqData.Link)
	if err == nil {
		controller.JsonExit(req, 0, "正在添加...", "")
	} else {
		controller.JsonExit(req, 1, err.Error())
	}
}
