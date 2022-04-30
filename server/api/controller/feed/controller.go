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
	Start int
	Size  int
}

type ChannelReqData struct {
	Start int
	Size  int
	Name  string `binding:"required"`
}

type ItemListByChannelIdReqData struct {
	Start     int
	Size      int
	ChannelId string `binding:"required"`
	UserId    string
}

type ItemListByUserIdReqData struct {
	Start  int
	Size   int
	UserId string `binding:"required"`
}

type SubChannelByUserIdAndChannelIdReqData struct {
	UserID    string `binding:"required"`
	ChannelId string `binding:"required"`
}

type GetSubChannelByUserIdReqData struct {
	UserID string `binding:"required"`
	Start  int
	Size   int
}

type GetChannelCatalogListByTagReqData struct {
	TagName string `binding:"required"`
	Start   int
	Size    int
}

type LatestFeedItemReqData struct {
	Start int
	Size  int
}

type MarkFeedItemReqData struct {
	UserId string `binding:"required"`
	ItemId string `binding:"required"`
}

type GetMarkFeedItemListReqData struct {
	UserId string `binding:"required"`
	Start  int
	Size   int
}

type SearchFeedItemReqData struct {
	Keyword string `binding:"required"`
	Start   int
	Size    int
}

type GetChannelCatalogListByUserIdReqData struct {
	UserId string `binding:"required"`
	Start  int
	Size   int
}

type GetRandomFeedItemListByUserIdReqData struct {
	UserId string
	Start  int
	Size   int
}

type GetChannelInfoByChannelUserIdReqData struct {
	UserID    string
	ChannelId string `binding:"required"`
}

type SubItemByUserIdAndItemIdReqData struct {
	UserID string
	ItemId string `binding:"required"`
}

type AddFeedChannelByLinkReqData struct {
	UserID string `binding:"required"`
	Link   string `binding:"required"`
}
