package routers

import (
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
	router.Run(":8080")
}

func initStaticWebRes(router *gin.Engine){
	router.StaticFS("/view", http.FS(static.Static))
}

func initV1API(router *gin.Engine) {
	v1 := router.Group("/v1/api")
	v1.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
