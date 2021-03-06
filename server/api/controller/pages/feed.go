package pages

import (
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

func (ctl *Controller) Error(req *gin.Context) {
	msg := req.Query("msg")
	title := req.Query("title")
	req.HTML(http.StatusOK, "base/error.html", gin.H{
		"toolBarTitle": title,
		"errorMsg":     msg,
	})
}

func (ctl *Controller) AddChannel(req *gin.Context) {
	templateMap := gin.H{
		"feedDrawerTab": "add",
		"toolBarTitle":  "添加订阅",
		"RSSLink":       "RSS链接",
		"AddBtnText":    "添加订阅",
	}
	req.HTML(http.StatusOK, "page/addChannelPage.html", getCommonTemplateMap(templateMap))
}

func (ctl *Controller) UserAllFeedItemListTmpl(req *gin.Context) {
	var reqData *ctlFeed.ItemListByUserIdReqData
	if err := ctl.BaseController.ValidateQuery(req, &reqData); err != nil {
		return
	}
	itemList := feed.GetFeedItemListByUserId(req.Request.Context(), reqData.UserId, reqData.Start, reqData.Size)
	var message string
	if len(itemList) == 0 {
		message = "没有更多的文章了，请订阅更多的频道"
	}

	items := make([]FeedItemType, 0)
	for i := 0; i < len(itemList); i++ {
		item := FeedItemType{
			Id:              itemList[i].Id,
			ChannelId:       itemList[i].ChannelId,
			Title:           itemList[i].Title,
			Description:     itemList[i].Description,
			Link:            itemList[i].Link,
			Content:         itemList[i].Content,
			Thumbnail:       itemList[i].Thumbnail,
			Date:            "发布于：" + itemList[i].Date.Format("Y-m-d H:i"),
			Author:          "作者：" + itemList[i].Author,
			InputDate:       "发布于：" + itemList[i].InputDate.Format("Y-m-d H:i"),
			RssLink:         itemList[i].RssLink,
			ChannelImageUrl: itemList[i].ChannelImageUrl,
			ChannelTitle:    itemList[i].ChannelTitle,
			Marked:          itemList[i].Marked,
			Sub:             itemList[i].Sub,
		}
		items = append(items, item)
	}

	req.HTML(http.StatusOK, "component/feedItemList.html", gin.H{
		"items":      items,
		"message":    message,
		"showMarked": true,
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
	subChannelList := feed.GetSubChannelListByUserId(req.Request.Context(), reqData.UserID, reqData.Start, reqData.Size)
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
	channelInfo := feed.GetChannelInfoByChannelAndUserId(req.Request.Context(), userId, channelId)
	req.HTML(http.StatusOK, "page/channelPage.html", gin.H{
		"channelInfo":     channelInfo,
		"toolBarTitle":    channelInfo.Title,
		"loadMoreBtnText": "正在加载...",
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
	itemList := feed.GetFeedItemByChannelId(req.Request.Context(), reqData.Start, reqData.Size, reqData.ChannelId, reqData.UserId)
	var message string
	if len(itemList) == 0 {
		message = "频道没有更多文章了"
	}

	items := make([]FeedItemType, 0)
	for i := 0; i < len(itemList); i++ {
		item := FeedItemType{
			Id:              itemList[i].Id,
			ChannelId:       itemList[i].ChannelId,
			Title:           itemList[i].Title,
			Description:     itemList[i].Description,
			Link:            itemList[i].Link,
			Content:         itemList[i].Content,
			Thumbnail:       itemList[i].Thumbnail,
			Date:            "发布于：" + itemList[i].Date.Format("Y-m-d H:i"),
			Author:          "作者：" + itemList[i].Author,
			InputDate:       "发布于：" + itemList[i].InputDate.Format("Y-m-d H:i"),
			RssLink:         itemList[i].RssLink,
			ChannelImageUrl: itemList[i].ChannelImageUrl,
			ChannelTitle:    itemList[i].ChannelTitle,
			Marked:          itemList[i].Marked,
			Sub:             itemList[i].Sub,
		}
		items = append(items, item)
	}

	req.HTML(http.StatusOK, "component/feedItemList.html", gin.H{
		"items":      items,
		"message":    message,
		"showMarked": true,
	})
}

func (ctl *Controller) GetSearchPageTmpl(req *gin.Context) {
	templateMap := gin.H{
		"feedDrawerTab":   "search",
		"toolBarTitle":    "搜索",
		"loadMoreBtnText": "正在加载...",
		"title":           "锅烧阅读",
	}
	req.HTML(http.StatusOK, "page/search.html", getCommonTemplateMap(templateMap))
}

func (ctl *Controller) GetSearchResultListTmpl(req *gin.Context) {
	var reqData *ctlFeed.SearchFeedItemReqData
	if err := ctl.BaseController.ValidateQuery(req, &reqData); err != nil {
		return
	}
	itemList := feed.SearchFeedItem(req.Request.Context(), reqData.UserId, reqData.Keyword, reqData.Start, reqData.Size)
	var message string
	if len(itemList) == 0 {
		message = "没有更多文章了"
	}

	items := make([]FeedItemType, 0)
	for i := 0; i < len(itemList); i++ {
		item := FeedItemType{
			Id:              itemList[i].Id,
			ChannelId:       itemList[i].ChannelId,
			Title:           itemList[i].Title,
			Description:     itemList[i].Description,
			Link:            itemList[i].Link,
			Content:         itemList[i].Content,
			Thumbnail:       itemList[i].Thumbnail,
			Date:            "发布于：" + itemList[i].Date.Format("Y-m-d H:i"),
			Author:          "作者：" + itemList[i].Author,
			InputDate:       "发布于：" + itemList[i].InputDate.Format("Y-m-d H:i"),
			RssLink:         itemList[i].RssLink,
			ChannelImageUrl: itemList[i].ChannelImageUrl,
			ChannelTitle:    itemList[i].ChannelTitle,
			Marked:          itemList[i].Marked,
			Sub:             itemList[i].Sub,
		}
		items = append(items, item)
	}

	req.HTML(http.StatusOK, "component/feedItemList.html", gin.H{
		"items":           items,
		"toolBarTitle":    "搜索",
		"loadMoreBtnText": "正在加载...",
		"title":           "锅烧阅读",
		"showMarked":      true,
		"message":         message,
	})
}

func (ctl *Controller) GetMarkedFeedItemPageTmpl(req *gin.Context) {
	templateMap := gin.H{
		"toolBarTitle":    "收藏",
		"feedDrawerTab":   "mark",
		"loadMoreBtnText": "正在加载...",
		"title":           "锅烧阅读",
	}
	req.HTML(http.StatusOK, "page/markedItemPage.html", getCommonTemplateMap(templateMap))
}

func (ctl *Controller) GetMarkedFeedItemListTmpl(req *gin.Context) {
	var reqData *ctlFeed.GetMarkFeedItemListReqData
	if err := ctl.BaseController.ValidateQuery(req, &reqData); err != nil {
		return
	}
	itemList := feed.GetMarkedFeedItemListByUserId(req.Request.Context(), reqData.UserId, reqData.Start, reqData.Size)
	var message string
	if len(itemList) == 0 {
		message = "没有更多文章了"
	}

	items := make([]FeedItemType, 0)
	for i := 0; i < len(itemList); i++ {
		item := FeedItemType{
			Id:              itemList[i].Id,
			ChannelId:       itemList[i].ChannelId,
			Title:           itemList[i].Title,
			Description:     itemList[i].Description,
			Link:            itemList[i].Link,
			Content:         itemList[i].Content,
			Thumbnail:       itemList[i].Thumbnail,
			Date:            "发布于：" + itemList[i].Date.Format("Y-m-d H:i"),
			Author:          "作者：" + itemList[i].Author,
			InputDate:       "发布于：" + itemList[i].InputDate.Format("Y-m-d H:i"),
			RssLink:         itemList[i].RssLink,
			ChannelImageUrl: itemList[i].ChannelImageUrl,
			ChannelTitle:    itemList[i].ChannelTitle,
			Marked:          itemList[i].Marked,
			Sub:             itemList[i].Sub,
		}
		items = append(items, item)
	}

	req.HTML(http.StatusOK, "component/feedItemList.html", gin.H{
		"items":           itemList,
		"toolBarTitle":    "收藏",
		"loadMoreBtnText": "正在加载...",
		"title":           "锅烧阅读",
		"showMarked":      false,
		"message":         message,
	})
}

func (ctl *Controller) GetFeedItemSharePageTmpl(req *gin.Context) {
	id := req.Param("id")
	uid := req.Param("uid")
	sharedItemInfo := feed.GetFeedItemByItemId(req.Request.Context(), id, uid)
	feedChannelInfo := feed.GetChannelInfoByChannelAndUserId(req.Request.Context(), uid, sharedItemInfo.ChannelId)
	feedItemList := feed.GetFeedItemByChannelId(req.Request.Context(), 0, 10, sharedItemInfo.ChannelId, uid)

	templateMap := gin.H{
		"title":           sharedItemInfo.Title,
		"toolBarTitle":    "锅烧阅读",
		"sharedItemInfo":  sharedItemInfo,
		"feedChannelInfo": feedChannelInfo,
		"items":           feedItemList,
		"countMsg":        "共" + feedChannelInfo.Count + "篇文章",
		"dateText":        "发布于：" + sharedItemInfo.Date.Format("Y-m-d H:i"),
		"authorText":      "作者：" + sharedItemInfo.Author,
		"homeLinkText":    "前往主页浏览更多",
	}
	req.HTML(http.StatusOK, "page/sharedItemPage.html", getCommonTemplateMap(templateMap))
}

func (ctl *Controller) GetFeedItemDetailPageTmpl(req *gin.Context) {
	id := req.Param("id")
	uid := req.Param("uid")
	itemDetailInfo := feed.GetFeedItemByItemId(req.Request.Context(), id, uid)

	templateMap := gin.H{
		"title":        itemDetailInfo.Title,
		"toolBarTitle": "锅烧阅读",
		"item":         itemDetailInfo,
		"dateText":     "时间：" + itemDetailInfo.Date.Format("Y-m-d H:i"),
		"authorText":   "作者：" + itemDetailInfo.Author,
		"linkText":     "浏览原文",
	}
	req.HTML(http.StatusOK, "page/feedItemDetail.html", getCommonTemplateMap(templateMap))
}
