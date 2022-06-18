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
		err = router.Run(":" + *port)
	}

	if err != nil {
		ctx := context.Background()
		component.Logger().Error(ctx, err.Error())
		panic(err)
	}
}

func initV1API(router *gin.Engine) {
	unAuthorized := router.Group("/")
	userCtl := user.UsrCtl
	unAuthorized.POST("/v1/api/user/register", userCtl.RegisterUser)
	unAuthorized.POST("/v1/api/user/login", userCtl.Login)

	authorized := router.Group("/")
	authorized.Use(middlewear.AuthToken())
	// feed channel
	authorized.GET("/v1/api/feed/channel/by_tag", feed.FeedCtl.GetFeedChannelByTag)
	authorized.GET("/v1/api/feed/channel/by_uid", feed.FeedCtl.GetFeedChannelInfoByChannelAndUserId)
	authorized.GET("/v1/api/feed/channel/catalogs/by_tag", feed.FeedCtl.GetFeedChannelCatalogListByTag)
	authorized.GET("/v1/api/feed/channel/sub/", feed.FeedCtl.GetSubFeedChannelByUserId)
	authorized.POST("/v1/api/feed/channel/sub/uid", feed.FeedCtl.SubChannelByUserIdAndChannelId)
	authorized.GET("/v1/api/feed/channel/catalogs/by_uid", feed.FeedCtl.GetFeedChannelCatalogListByUserId)
	authorized.POST("/v1/api/feed/link/uid", feed.FeedCtl.AddFeedChannelByLink)
	// feed item
	authorized.GET("/v1/api/feed/latest", feed.FeedCtl.GetLatestFeedItem)
	authorized.GET("/v1/api/feed/random", feed.FeedCtl.GetRandomFeedItem)
	authorized.GET("/v1/api/feed/search", feed.FeedCtl.SearchFeedItem)
	authorized.GET("/v1/api/feed/item/by_uid", feed.FeedCtl.GetFeedItemListByUserId)
	authorized.GET("/v1/api/feed/item/cid", feed.FeedCtl.GetFeedItemByChannelId)
	authorized.GET("/v1/api/feed/item/id", feed.FeedCtl.GetFeedItemByItemId)
	authorized.POST("/v1/api/feed/item/mark", feed.FeedCtl.MarkFeedItemByUserId)
	authorized.GET("/v1/api/feed/item/mark", feed.FeedCtl.GetMarkFeedItemListByUserId)
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
	pagesGroupWithCookie := router.Group("/")
	pagesGroupWithCookie.Use(middlewear.CookieToken())
	pageCtl := pages.PagesCtl

	pagesGroupWithCookie.GET("/view/", pageCtl.Index)
	pagesGroupWithCookie.GET("/view/add", pageCtl.AddChannel)
	pagesGroupWithCookie.GET("/view/mark", pageCtl.GetMarkedFeedItemPageTmpl)
	pagesGroupWithCookie.GET("/view/search/", pageCtl.GetSearchPageTmpl)

	pagesGroupWithCookie.GET("/view/feed/channel/info/:channelId/:userId", pageCtl.GetFeedChannelPageTmpl)
	pagesGroupWithCookie.GET("/view/feed/channel/items/", pageCtl.GetFeedChannelItemListTmpl)
	pagesGroupWithCookie.GET("/view/feed/all/item/list", pageCtl.UserAllFeedItemListTmpl)
	pagesGroupWithCookie.GET("/view/feed/sub_list", pageCtl.GetSubFeedChannelListTmpl)
	pagesGroupWithCookie.GET("/view/feed/search/result", pageCtl.GetSearchResultListTmpl)
	pagesGroupWithCookie.GET("/view/feed/items/mark", pageCtl.GetMarkedFeedItemListTmpl)

	pagesGroupWithoutCookie := router.Group("/")
	pagesGroupWithoutCookie.GET("/view/user/login", pageCtl.Login)
	pagesGroupWithoutCookie.GET("/view/user/register", pageCtl.Register)
	pagesGroupWithoutCookie.GET("/view/f/i/s/:id", pageCtl.GetFeedItemSharePageTmpl)
	pagesGroupWithoutCookie.GET("/view/error", pageCtl.Error)
	pagesGroupWithoutCookie.GET("/view/feed/item/detail/:id", pageCtl.GetFeedItemDetailPageTmpl)
}
