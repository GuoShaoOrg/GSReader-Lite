package feed

import (
	"context"
	"gs-reader-lite/server/api/controller"
	"gs-reader-lite/server/api/service/feed"

	"github.com/gin-gonic/gin"
)

func (ctl *Controller) GetFeedItemByChannelId(req *gin.Context) {
	var reqData *ItemListByChannelIdReqData
	if err := ctl.BaseController.ValidateQuery(req, &reqData); err != nil {
		return
	}
	if reqData.Size == 0 {
		reqData.Size = 10
	}

	list := feed.GetFeedItemByChannelId(context.Background(), reqData.Start, reqData.Size, reqData.ChannelId, reqData.UserId)
	controller.JsonExit(req, 0, "success", list)
}

func (ctl *Controller) GetFeedItemListByUserId(req *gin.Context) {
	var reqData *ItemListByUserIdReqData
	if err := ctl.BaseController.ValidateQuery(req, &reqData); err != nil {
		return
	}
	if reqData.Size == 0 {
		reqData.Size = 10
	}

	itemList := feed.GetFeedItemListByUserId(context.Background(), reqData.UserId, reqData.Start, reqData.Size)
	if len(itemList) == 0 {
		latestList := feed.GetLatestFeedItem(context.Background(), reqData.UserId, reqData.Start, reqData.Size)
		controller.JsonExit(req, 0, "success", latestList)
	}
	controller.JsonExit(req, 0, "success", itemList)
}

func (ctl *Controller) GetLatestFeedItem(req *gin.Context) {
	var reqData *LatestFeedItemReqData
	if err := ctl.BaseController.ValidateQuery(req, &reqData); err != nil {
		return
	}
	if reqData.Size == 0 {
		reqData.Size = 10
	}
	latestList := feed.GetLatestFeedItem(context.Background(), "", reqData.Start, reqData.Size)
	controller.JsonExit(req, 0, "success", latestList)
}

func (ctl *Controller) MarkFeedItemByUserId(req *gin.Context) {
	var reqData *MarkFeedItemReqData
	if err := ctl.BaseController.ValidateJson(req, &reqData); err != nil {
		return
	}

	status, err := feed.MarkFeedItem(context.Background(), reqData.UserId, reqData.ItemId)
	msg := "发生了一些问题"
	if err != nil {
		controller.JsonExit(req, 1, msg)
	}
	if status == 1 {
		msg = "收藏成功"
	} else {
		msg = "取消收藏"
	}
	controller.JsonExit(req, 0, msg, status)
}

func (ctl *Controller) GetMarkFeedItemListByUserId(req *gin.Context) {
	var reqData *GetMarkFeedItemListReqData
	if err := ctl.BaseController.ValidateQuery(req, &reqData); err != nil {
		return
	}

	resultSet := feed.GetMarkedFeedItemListByUserId(context.Background(), reqData.UserId, reqData.Start, reqData.Size)
	controller.JsonExit(req, 0, "success", resultSet)
}

func (ctl *Controller) SearchFeedItem(req *gin.Context) {
	var reqData *SearchFeedItemReqData
	if err := ctl.BaseController.ValidateQuery(req, &reqData); err != nil {
		return
	}

	resultSet := feed.SearchFeedItem(context.Background(), reqData.UserId, reqData.Keyword, reqData.Start, reqData.Size)
	controller.JsonExit(req, 0, "success", resultSet)
}

func (ctl *Controller) GetRandomFeedItem(req *gin.Context) {
	var reqData *GetRandomFeedItemListByUserIdReqData
	if err := ctl.BaseController.ValidateQuery(req, &reqData); err != nil {
		return
	}

	resultSet := feed.GetRandomFeedItem(context.Background(), reqData.Start, reqData.Size, reqData.UserId)
	controller.JsonExit(req, 0, "success", resultSet)
}

func (ctl *Controller) GetFeedItemByItemId(req *gin.Context) {
	var reqData *SubItemByUserIdAndItemIdReqData
	if err := ctl.BaseController.ValidateQuery(req, &reqData); err != nil {
		return
	}

	itemInfo := feed.GetFeedItemByItemId(context.Background(), reqData.ItemId, reqData.UserID)
	controller.JsonExit(req, 0, "success", itemInfo)
}
