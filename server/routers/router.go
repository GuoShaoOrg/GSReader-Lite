package routers

import (
	"context"
	"gs-reader-lite/server/api/controller/feed"
	"gs-reader-lite/server/api/controller/user"
	"gs-reader-lite/server/component"
	"gs-reader-lite/server/middlewear"
	"gs-reader-lite/web/static"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()
	router.Use(middlewear.StaticRedirect())
	initStaticWebRes(router)
	initV1API(router)
	err := router.Run(":8080")
	if err != nil {
		ctx := context.Background()
		component.Logger().Error(ctx, err.Error())
		panic(err)
	}
}

func initStaticWebRes(router *gin.Engine) {
	router.StaticFS("/view", http.FS(static.Static))
}

func initV1API(router *gin.Engine) {
	v1 := router.Group("/v1/api")
	authorized := v1.Group("/")
	authorized.Use(middlewear.AuthToken())

	userGroup := v1.Group("/user")
	{
		userCtl := user.UsrCtl
		userGroup.POST("/register", userCtl.RegisterUser)
		userGroup.POST("/login", userCtl.Login)
	}

	feedGroup := authorized.Group("/feed")
	{
		// feed channel
		feedGroup.GET("/tag", feed.FeedCtl.GetFeedTag)
		feedGroup.GET("/channel_by_tag", feed.FeedCtl.GetFeedChannelByTag)
		feedGroup.GET("/channel_info_by_id", feed.FeedCtl.GetFeedChannelInfoByChannelAndUserId)
		feedGroup.GET("/channel_catalog_list_by_tag", feed.FeedCtl.GetFeedChannelCatalogListByTag)
		feedGroup.GET("/channel_by_user_id", feed.FeedCtl.GetSubFeedChannelByUserId)
		feedGroup.POST("/sub_channel_by_user_id", feed.FeedCtl.SubChannelByUserIdAndChannelId)
		feedGroup.GET("/channel_catalog_list_by_user_id", feed.FeedCtl.GetFeedChannelCatalogListByUserId)
		// feed item
		feedGroup.GET("/latest", feed.FeedCtl.GetLatestFeedItem)
		feedGroup.GET("/random", feed.FeedCtl.GetRandomFeedItem)
		feedGroup.GET("/search", feed.FeedCtl.SearchFeedItem)
		feedGroup.GET("/item_by_user_id", feed.FeedCtl.GetFeedItemListByUserId)
		feedGroup.GET("/item_by_channel_id", feed.FeedCtl.GetFeedItemByChannelId)
		feedGroup.GET("/item_by_item_id", feed.FeedCtl.GetFeedItemByItemId)
		feedGroup.POST("/mark_feed_item_by_user_id", feed.FeedCtl.MarkFeedItemByUserId)
		feedGroup.GET("/mark_feed_item_by_user_id", feed.FeedCtl.MarkFeedItemByUserId)
		feedGroup.GET("/link/uid", feed.FeedCtl.AddFeedChannelByLink)
	}
}
