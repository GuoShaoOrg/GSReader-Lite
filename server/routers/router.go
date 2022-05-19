package routers

import (
	"context"
	"flag"
	"gs-reader-lite/public"
	"gs-reader-lite/server/api/controller/feed"
	"gs-reader-lite/server/api/controller/pages"
	"gs-reader-lite/server/api/controller/user"
	"gs-reader-lite/server/component"
	"gs-reader-lite/server/middlewear"
	"gs-reader-lite/templates"
	"html/template"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()
	initPages(router)
	initV1API(router)
	var err error
	if os.Getenv("env") == "dev" {
		err = router.Run(":8083")
	} else {
		port := flag.String("port", "80", "listen port")
		flag.Parse()
		gin.SetMode(gin.ReleaseMode)
		err = router.Run(":" + *port)
	}

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

func initPages(router *gin.Engine) {
	if os.Getenv("env") == "dev" {
		router.LoadHTMLGlob("templates/**/*")
		router.Static("/public", "./public")
	} else {
		tmplFS := templates.Templates
		templ := template.Must(template.New("").ParseFS(tmplFS, "*/*.html"))
		router.SetHTMLTemplate(templ)
		publicFS := public.Public
		router.StaticFS("/public", http.FS(publicFS))
	}

	router.Use(middlewear.StaticRedirect())
	pagesGroup := router.Group("/view")
	pageCtl := pages.PagesCtl
	{
		pagesGroup.GET("/", middlewear.CookieToken(), pageCtl.Index)
		pagesGroup.GET("/add", middlewear.CookieToken(), pageCtl.AddChannel)
		pagesGroup.GET("/search/", middlewear.CookieToken(), pageCtl.GetSearchPageTmpl)

		pagesGroup.GET("/user/login", pageCtl.Login)
		pagesGroup.GET("/user/register", pageCtl.Register)
		pagesGroup.GET("/error", pageCtl.Error)

		pagesGroup.GET("/feed/channel/info/:channelId/:userId", middlewear.CookieToken(), pageCtl.GetFeedChannelPageTmpl)
		pagesGroup.GET("/feed/channel/items/", middlewear.CookieToken(), pageCtl.GetFeedChannelItemListTmpl)
		pagesGroup.GET("/feed/all/item/list", middlewear.CookieToken(), pageCtl.UserAllFeedItemListTmpl)
		pagesGroup.GET("/feed/sub_list", middlewear.CookieToken(), pageCtl.GetSubFeedChannelListTmpl)
		pagesGroup.GET("/feed/search/result", middlewear.CookieToken(), pageCtl.GetSearchResultListTmpl)

	}
}
