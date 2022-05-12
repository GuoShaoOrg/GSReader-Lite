package pages

import (
	"context"
	ctlFeed "gs-reader-lite/server/api/controller/feed"
	"gs-reader-lite/server/api/service/feed"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctl *Controller) Home(req *gin.Context) {
	req.HTML(http.StatusOK, "index.html", gin.H{
		"username": "管理员",
	})
}

func (ctl *Controller) HomeContainerListTmpl(req *gin.Context) {
	var reqData *ctlFeed.ItemListByUserIdReqData
	if err := ctl.BaseController.ValidateQuery(req, &reqData); err != nil {
		return
	}
	ctx := context.Background()
	itemList := feed.GetFeedItemListByUserId(ctx, reqData.UserId, reqData.Start, reqData.Size)
	req.HTML(http.StatusOK, "home-container-list.html", gin.H{
		"items": itemList,
	})
}
