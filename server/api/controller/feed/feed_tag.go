package feed

import (
	"context"
	"gs-reader-lite/server/api/controller"
	"gs-reader-lite/server/api/service/feed"

	"github.com/gin-gonic/gin"
)

func (ctl *Controller) GetFeedTag(req *gin.Context) {
	var reqData *TagReqData
	if err := ctl.BaseController.Validate(req, &reqData); err != nil {
		return
	}
	if reqData.Size == 0 {
		reqData.Size = 10
	}
	tagList := feed.GetFeedTag(context.Background(), reqData.Start, reqData.Size)
	controller.JsonExit(req, 0, "success", tagList)
}
