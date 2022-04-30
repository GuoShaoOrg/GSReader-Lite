package routers

import (
	"context"
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
	userGroup := v1.Group("/user")
	{
		userCtl := user.UsrCtl
		userGroup.POST("/register", userCtl.RegisterUser)
		userGroup.POST("/login", userCtl.Login)
	}

	authorized := v1.Group("/")
	authorized.Use(middlewear.AuthToken())
	{
		authorized.GET("/feed/item_by_channel_id", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "item_by_channel_id",
			})
		})
	}
}
