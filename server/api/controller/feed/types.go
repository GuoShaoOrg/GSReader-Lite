package feed

import (
	"gs-reader-lite/server/api/controller"
)

type Controller struct {
	controller.BaseController
}

var FeedCtl = &Controller{}

type RouterInfoData struct {
	Route string
	Port  string
}

type TagReqData struct {
	Start int `form:"start"`
	Size  int `form:"size"`
}

type ChannelReqData struct {
	Start int    `form:"start"`
	Size  int    `form:"size"`
	Name  string `form:"name" binding:"required"`
}

type ItemListByChannelIdReqData struct {
	Start     int    `form:"start"`
	Size      int    `form:"size"`
	ChannelId string `form:"channelId" binding:"required"`
	UserId    string `form:"userId"`
}

type ItemListByUserIdReqData struct {
	Start  int    `form:"start"`
	Size   int    `form:"size"`
	UserId string `form:"userId" binding:"required"`
}

type SubChannelByUserIdAndChannelIdReqData struct {
	UserID    string `json:"userId" binding:"required"`
	ChannelId string `json:"channelId" binding:"required"`
}

type GetSubChannelByUserIdReqData struct {
	UserID string `form:"userId" binding:"required"`
	Start  int    `form:"start"`
	Size   int    `form:"size"`
}

type GetChannelCatalogListByTagReqData struct {
	TagName string `binding:"required"`
	Start   int    `form:"start"`
	Size    int    `form:"size"`
}

type LatestFeedItemReqData struct {
	Start int `form:"start"`
	Size  int `form:"size"`
}

type MarkFeedItemReqData struct {
	UserId string `json:"userId" binding:"required"`
	ItemId string `json:"itemId" binding:"required"`
}

type GetMarkFeedItemListReqData struct {
	UserId string `binding:"required"`
	Start  int    `form:"start"`
	Size   int    `form:"size"`
}

type SearchFeedItemReqData struct {
	Keyword string `form:"keyword" binding:"required"`
	Start   int    `form:"start"`
	Size    int    `form:"size"`
}

type GetChannelCatalogListByUserIdReqData struct {
	UserId string `binding:"required"`
	Start  int    `form:"start"`
	Size   int    `form:"size"`
}

type GetRandomFeedItemListByUserIdReqData struct {
	UserId string `form:"userId" binding:"required"`
	Start  int    `form:"start"`
	Size   int    `form:"size"`
}

type GetChannelInfoByChannelUserIdReqData struct {
	UserID    string `form:"userId" binding:"required"`
	ChannelId string `binding:"required"`
}

type SubItemByUserIdAndItemIdReqData struct {
	UserID string `form:"userId" binding:"required"`
	ItemId string `binding:"required"`
}

type AddFeedChannelByLinkReqData struct {
	UserID string `json:"userId" binding:"required"`
	Link   string `json:"link" binding:"required"`
}

type AddFeedItemReqData struct {
	Id          string `json:"id"`
	ChannelId   string `json:"channelId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Author      string `json:"author"`
	Thumbnail   string `json:"thumbnail"`
	Content     string `json:"content"`
}
