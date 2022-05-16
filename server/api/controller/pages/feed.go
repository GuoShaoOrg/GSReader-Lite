package pages

import (
	"context"
	ctlFeed "gs-reader-lite/server/api/controller/feed"
	"gs-reader-lite/server/api/service/feed"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getCommonTemplateMap(tempMap gin.H) gin.H {
	commonMap := gin.H{
		"loadMoreBtn":   "点击加载更多",
		"title":         "锅烧阅读",
		"userSubHeader": "用户",
		"homeTitle":     "全部文章",
		"subTitle":      "已订阅源",
		"searchTitle":   "搜索文章",
		"favTitle":      "收藏文章",
		"addTitle":      "添加订阅",
	}

	for k, v := range tempMap {
		commonMap[k] = v
	}

	return commonMap
}

func (ctl *Controller) Index(req *gin.Context) {
	templateMap := gin.H{
		"feedDrawerTab": "home",
		"toolBarTitle":  "全部文章",
	}
	req.HTML(http.StatusOK, "index.html", getCommonTemplateMap(templateMap))
}

func (ctl *Controller) SubChannel(req *gin.Context) {
	templateMap := gin.H{
		"feedDrawerTab": "sub",
		"toolBarTitle":  "已订阅源",
	}
	req.HTML(http.StatusOK, "index.html", getCommonTemplateMap(templateMap))
}

func (ctl *Controller) AddChannel(req *gin.Context) {
	templateMap := gin.H{
		"feedDrawerTab": "add",
		"toolBarTitle":  "添加订阅",
		"RSSLink":       "RSS链接",
		"AddBtnText":    "添加订阅",
	}
	req.HTML(http.StatusOK, "index.html", getCommonTemplateMap(templateMap))
}

func (ctl *Controller) UserAllFeedItemListTmpl(req *gin.Context) {
	var reqData *ctlFeed.ItemListByUserIdReqData
	if err := ctl.BaseController.ValidateQuery(req, &reqData); err != nil {
		return
	}
	ctx := context.Background()
	itemList := feed.GetFeedItemListByUserId(ctx, reqData.UserId, reqData.Start, reqData.Size)
	var message string
	if len(itemList) == 0 {
		message = "没有更多的文章了，请订阅更多的频道"
	}
	req.HTML(http.StatusOK, "feed/feedItemList.html", gin.H{
		"items":   itemList,
		"message": message,
	})
}

func (ctl *Controller) GetSubFeedChannelListTmpl(req *gin.Context) {
	var reqData *ctlFeed.GetSubChannelByUserIdReqData
	if err := ctl.BaseController.ValidateQuery(req, &reqData); err != nil {
		return
	}
	if reqData.Size == 0 {
		reqData.Size = 10
	}
	subChannelList := feed.GetSubChannelListByUserId(context.Background(), reqData.UserID, reqData.Start, reqData.Size)
	var message string
	if len(subChannelList) == 0 {
		message = "您还没有订阅任何频道"
	}
	req.HTML(http.StatusOK, "feed/subedFeedList.html", gin.H{
		"subChannelList": subChannelList,
		"message":        message,
	})
}

func (ctl *Controller) GetFeedChannelPageTmpl(req *gin.Context) {
	channelId := req.Param("channelId")
	userId := req.Param("userId")
	channelInfo := feed.GetChannelInfoByChannelAndUserId(context.Background(), userId, channelId)
	req.HTML(http.StatusOK, "feed/channelPage.html", gin.H{
		"channelInfo":     channelInfo,
		"toolBarTitle":    channelInfo.Title,
		"loadMoreBtnText": "点击加载更多",
		"title":           "锅烧阅读",
	})
}

func (ctl *Controller) GetFeedChannelItemListTmpl(req *gin.Context) {
	var reqData *ctlFeed.ItemListByChannelIdReqData
	if err := ctl.BaseController.ValidateQuery(req, &reqData); err != nil {
		return
	}
	if reqData.Size == 0 {
		reqData.Size = 10
	}
	channleItemList := feed.GetFeedItemByChannelId(context.Background(), reqData.Start, reqData.Size, reqData.ChannelId, reqData.UserId)
	var message string
	if len(channleItemList) == 0 {
		message = "频道还没更多文章了"
	}
	req.HTML(http.StatusOK, "feed/feedItemList.html", gin.H{
		"items":   channleItemList,
		"message": message,
	})
}
