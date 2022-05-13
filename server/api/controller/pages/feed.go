package pages

import (
	"context"
	ctlFeed "gs-reader-lite/server/api/controller/feed"
	"gs-reader-lite/server/api/service/feed"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctl *Controller) Home(req *gin.Context) {
	req.HTML(http.StatusOK, "index.html", gin.H{})
}

func (ctl *Controller) FeedItemListTmpl(req *gin.Context) {
	var reqData *ctlFeed.ItemListByUserIdReqData
	if err := ctl.BaseController.ValidateQuery(req, &reqData); err != nil {
		return
	}
	ctx := context.Background()
	itemList := feed.GetFeedItemListByUserId(ctx, reqData.UserId, reqData.Start, reqData.Size)
	var message string
	if len(itemList) == 0 {
		message = "您还没有订阅任何文章"
	}
	req.HTML(http.StatusOK, "feed-item-list.html", gin.H{
		"items":   itemList,
		"message": message,
	})
}

func (ctl *Controller) AddFeedChannelTmpl(req *gin.Context) {
	req.HTML(http.StatusOK, "add-feed.html", gin.H{})
}
