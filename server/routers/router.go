package routers

import (
	"context"
	"gs-reader-lite/server/api/controller/feed"
	"gs-reader-lite/server/api/controller/pages"
	"gs-reader-lite/server/api/controller/user"
	"gs-reader-lite/server/component"
	"gs-reader-lite/server/middlewear"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()
	initStaticRoute(router)
	initPages(router)
	initV1API(router)
	err := router.Run(":8083")
	if err != nil {
		ctx := context.Background()
		component.Logger().Error(ctx, err.Error())
		panic(err)
	}
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
		feedGroup.GET("/channel_by_tag", feed.FeedCtl.GetFeedChannelByTag)
		feedGroup.GET("/channel_info_by_id", feed.FeedCtl.GetFeedChannelInfoByChannelAndUserId)
		feedGroup.GET("/channel_catalog_list_by_tag", feed.FeedCtl.GetFeedChannelCatalogListByTag)
		feedGroup.GET("/channel_by_user_id", feed.FeedCtl.GetSubFeedChannelByUserId)
		feedGroup.POST("/sub_channel_by_user_id", feed.FeedCtl.SubChannelByUserIdAndChannelId)
		feedGroup.GET("/channel_catalog_list_by_user_id", feed.FeedCtl.GetFeedChannelCatalogListByUserId)
		feedGroup.POST("/link/uid", feed.FeedCtl.AddFeedChannelByLink)
		// feed item
		feedGroup.GET("/latest", feed.FeedCtl.GetLatestFeedItem)
		feedGroup.GET("/random", feed.FeedCtl.GetRandomFeedItem)
		feedGroup.GET("/search", feed.FeedCtl.SearchFeedItem)
		feedGroup.GET("/item_by_user_id", feed.FeedCtl.GetFeedItemListByUserId)
		feedGroup.GET("/item_by_channel_id", feed.FeedCtl.GetFeedItemByChannelId)
		feedGroup.GET("/item_by_item_id", feed.FeedCtl.GetFeedItemByItemId)
		feedGroup.POST("/mark_feed_item_by_user_id", feed.FeedCtl.MarkFeedItemByUserId)
		feedGroup.GET("/mark_feed_item_by_user_id", feed.FeedCtl.MarkFeedItemByUserId)
	}
}

func initStaticRoute(router *gin.Engine) {
	router.Static("/public", "./public")
}

func initPages(router *gin.Engine) {
	router.LoadHTMLGlob("templates/*")
	router.Use(middlewear.StaticRedirect())
	pagesGroup := router.Group("/view")
	{
		pageCtl := pages.PagesCtl
		pagesGroup.GET("/", pageCtl.Home)
	}
}
