package feed

import (
	"context"
	"errors"
	"gs-reader-lite/server/component"
	"gs-reader-lite/server/lib"
	"gs-reader-lite/server/model"
	"gs-reader-lite/server/model/biz"
	"strconv"

	"github.com/gogf/gf/v2/encoding/ghash"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mmcdole/gofeed"
	"gorm.io/gorm"
)

func GetFeedChannelByTag(ctx context.Context, start, size int, name string) (feedList []biz.RssFeedChannelData) {

	if err := component.GetDatabase().Table("rss_feed_tag rft").
		Joins("left join rss_feed_channel rsc on rft.channel_id=rsc.id").
		Joins("left join user_sub_channel usc on usc.channel_id=rfc.id").
		Select("rfc.*, usc.status as sub").
		Where("rft.name", name).
		Limit(size).
		Offset(start).
		Find(&feedList); err != nil {
		component.Logger().Error(ctx, err)
	}

	return
}

func GetAllFeedChannelList(ctx context.Context) (feedList []biz.RssFeedChannelData) {
	feedChannelModel := make([]model.RssFeedChannel, 0)
	err := component.GetDatabase().Find(&feedChannelModel).Error
	if err != nil {
		component.Logger().Error(ctx, err)
		return
	}
	for _, v := range feedChannelModel {
		feed := biz.RssFeedChannelData{}
		feed.Id = v.Id
		feed.Title = v.Title
		feed.ChannelDesc = v.ChannelDesc
		feed.ImageUrl = v.ImageUrl
		feed.Link = v.Link
		feed.RssLink = v.RssLink
		feedList = append(feedList, feed)
	}
	return
}

func SubChannelByUserIdAndChannelId(ctx context.Context, userId, channelId string) error {

	subChannelModel := model.UserSubChannel{}
	err := component.GetDatabase().Table("user_sub_channel").Where("user_id = ? and channel_id = ?", userId, channelId).Find(&subChannelModel)

	if err == nil && subChannelModel.ChannelId != "" {
		if subChannelModel.Status == 1 {
			subChannelModel.Status = 0
		} else {
			subChannelModel.Status = 1
		}
		result := component.GetDatabase().Table("user_sub_channel").Updates(model.UserSubChannel{Status: subChannelModel.Status}).Where("user_id = ? and channel_id = ?", userId, channelId)
		return result.Error
	}

	nowTime := gtime.Now().Format("Y-m-d H:i:s.u")
	subChannelModel = model.UserSubChannel{
		UserId:    userId,
		ChannelId: channelId,
		InputTime: gtime.NewFromStr(nowTime),
		Status:    1,
	}

	result := component.GetDatabase().Create(subChannelModel)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetSubChannelListByUserId(ctx context.Context, userId string, start, size int) (feedList []biz.RssFeedChannelData) {
	if err := component.GetDatabase().Table("user_sub_channel usc").
		Select("rfc.*, usc.status as sub").
		Joins("inner join rss_feed_channel rfc on usc.channel_id=rfc.id").
		Where("usc.user_id", userId).
		Where("usc.status", 1).
		Limit(size).
		Offset(start).
		Find(&feedList); err != nil {
		component.Logger().Error(ctx, err)
	}

	return
}

func GetFeedChannelCatalogListByTag(ctx context.Context, tagName string, start, size int) (feedCatalogList []biz.RssFeedChannelCatalogData) {
	rssFeedList := make([]biz.RssFeedChannelData, 0)
	rssFeedItemList := make([]model.RssFeedItem, 0)
	if err := component.GetDatabase().Table("rss_feed_tag rft").
		Joins("inner join rss_feed_channel rfc on rft.channel_id=rfc.id").
		Where("rft.name", tagName).
		Where("rfc.link like ?", "https%").
		Group("rfc.id").
		Limit(size).
		Offset(start).
		Find(&rssFeedList); err != nil {
		component.Logger().Error(ctx, err)
		return
	}

	for _, rssFeed := range rssFeedList {
		rssFeedItemQueryList := make([]model.RssFeedItem, 0)
		if err := component.GetDatabase().Table("rss_feed_item rfi").
			Select(model.RFIWithoutContentFieldSql).
			Where("rfi.channel_id", rssFeed.Id).
			Order("rfi.input_date desc").
			Limit(4).
			Offset(0).
			Find(&rssFeedItemQueryList); err != nil {
			component.Logger().Error(ctx, err)
			return
		}
		rssFeedItemList = append(rssFeedItemList, rssFeedItemQueryList...)
	}

	for _, rssFeed := range rssFeedList {
		feedCatalog := biz.RssFeedChannelCatalogData{}
		feedCatalog.Id = rssFeed.Id
		feedCatalog.Title = rssFeed.Title
		feedCatalog.ChannelDesc = rssFeed.ChannelDesc
		feedCatalog.ImageUrl = rssFeed.ImageUrl
		feedCatalog.Link = rssFeed.Link
		feedCatalog.Sub = rssFeed.Sub
		feedCatalog.RssLink = rssFeed.RssLink
		for _, rssFeedItem := range rssFeedItemList {
			if rssFeedItem.ChannelId == rssFeed.Id {
				rssFeedItemData := biz.RssFeedItemData{}
				rssFeedItemData.Id = rssFeedItem.Id
				rssFeedItemData.ChannelId = rssFeedItem.ChannelId
				rssFeedItemData.Title = rssFeedItem.Title
				rssFeedItemData.Description = rssFeedItem.Description
				rssFeedItemData.Content = rssFeedItem.Content
				rssFeedItemData.Link = rssFeedItem.Link
				rssFeedItemData.Date = rssFeedItem.Date
				rssFeedItemData.Author = rssFeedItem.Author
				rssFeedItemData.InputDate = rssFeedItem.InputDate
				rssFeedItemData.Thumbnail = rssFeedItem.Thumbnail
				feedCatalog.ItemList = append(feedCatalog.ItemList, rssFeedItemData)
			}
		}
		feedCatalogList = append(feedCatalogList, feedCatalog)

	}

	return
}

func GetFeedChannelCatalogListByUserId(ctx context.Context, userId string, start, size int) (feedCatalogList []biz.RssFeedChannelCatalogData) {
	rssFeedList := make([]biz.RssFeedChannelData, 0)
	rssFeedItemList := make([]model.RssFeedItem, 0)
	if err := component.GetDatabase().Table("user_sub_channel usc").
		Select("rfc.*, usc.status as sub").
		Joins("inner join rss_feed_channel rfc on usc.channel_id=rfc.id").
		Where("usc.user_id", userId).
		Where("usc.status", 1).
		//Where("rfc.link like ?", "https%").
		Order("usc.input_time desc").
		Limit(size).
		Offset(start).
		Find(&rssFeedList); err != nil {
		component.Logger().Error(ctx, err)
		return
	}

	for _, rssFeed := range rssFeedList {
		rssFeedItemQueryList := make([]model.RssFeedItem, 0)
		if err := component.GetDatabase().Table("rss_feed_item rfi").
			Select(model.RFIWithoutContentFieldSql).
			Where("rfi.channel_id", rssFeed.Id).
			Order("rfi.input_date desc").
			Limit(4).
			Offset(0).
			Find(&rssFeedItemQueryList); err != nil {
			component.Logger().Error(ctx, err)
			return
		}
		rssFeedItemList = append(rssFeedItemList, rssFeedItemQueryList...)
	}

	for _, rssFeed := range rssFeedList {
		feedCatalog := biz.RssFeedChannelCatalogData{}
		feedCatalog.Id = rssFeed.Id
		feedCatalog.Title = rssFeed.Title
		feedCatalog.ChannelDesc = rssFeed.ChannelDesc
		feedCatalog.ImageUrl = rssFeed.ImageUrl
		feedCatalog.Link = rssFeed.Link
		feedCatalog.Sub = rssFeed.Sub
		feedCatalog.RssLink = rssFeed.RssLink
		for _, rssFeedItem := range rssFeedItemList {
			if rssFeedItem.ChannelId == rssFeed.Id {
				rssFeedItemData := biz.RssFeedItemData{}
				rssFeedItemData.Id = rssFeedItem.Id
				rssFeedItemData.ChannelId = rssFeedItem.ChannelId
				rssFeedItemData.Title = rssFeedItem.Title
				rssFeedItemData.Description = rssFeedItem.Description
				rssFeedItemData.Content = rssFeedItem.Content
				rssFeedItemData.Link = rssFeedItem.Link
				rssFeedItemData.Date = rssFeedItem.Date
				rssFeedItemData.Author = rssFeedItem.Author
				rssFeedItemData.InputDate = rssFeedItem.InputDate
				rssFeedItemData.Thumbnail = rssFeedItem.Thumbnail
				feedCatalog.ItemList = append(feedCatalog.ItemList, rssFeedItemData)
			}
		}
		feedCatalogList = append(feedCatalogList, feedCatalog)

	}

	return
}

func GetChannelInfoByChannelAndUserId(ctx context.Context, userId, channelId string) (feedInfo biz.RssFeedChannelData) {
	if err := component.GetDatabase().Table("rss_feed_channel rfc").
		Joins("left join user_sub_channel usc on usc.channel_id=rfc.id and usc.user_id="+"'"+userId+"'").
		Select("rfc.*, usc.status as sub").
		Where("rfc.id", channelId).
		Find(&feedInfo).Error; err != nil {
		component.Logger().Error(ctx, err)
		return
	}

	var count int64
	if result := component.GetDatabase().Table("rss_feed_item rfi").
		Where("rfi.channel_id=?", channelId).
		Count(&count); result.Error != nil {
		component.Logger().Error(ctx, result.Error)
	} else {
		feedInfo.Count = strconv.Itoa(int(count))
	}

	return
}

func AddFeedChannelByLink(ctx context.Context, userID, rssLink string) (err error) {
	var (
		rssFeedChannelMode model.RssFeedChannel
		userSubChannel     model.UserSubChannel
		userInfoMode       model.UserInfo
		rssResp            string
	)

	if result := component.GetDatabase().Table("user_info ui").Where("ui.uid", userID).Find(&userInfoMode); result.Error != nil {
		component.Logger().Error(ctx, err)
		err = errors.New("用户不存在")
		return err
	}

	if result := component.GetDatabase().Table("rss_feed_channel rfc").Where("rfc.rss_link", rssLink).Find(&rssFeedChannelMode); result.Error != nil {
		component.Logger().Error(ctx, result.Error)
		err = errors.New("获取RSS链接失败")
		return err
	} else if rssFeedChannelMode.Id != "" {
		if result := component.GetDatabase().Table("user_sub_channel usc").Where("usc.channel_id=? and usc.user_id=?", rssFeedChannelMode.Id, userID).Find(&userSubChannel); result.Error != nil {
			component.Logger().Error(ctx, result.Error)
			err = errors.New("发生了一些问题")
			return err
		} else if userSubChannel.Status == 1 {
			err = errors.New("已经订阅过了")
			return err
		}
		userSubChannel.ChannelId = rssFeedChannelMode.Id
		userSubChannel.UserId = userID
		userSubChannel.Status = 1
		userSubChannel.InputTime = gtime.Now()
		result = component.GetDatabase().Table("user_sub_channel").Create(&userSubChannel)
		if result.Error != nil {
			component.Logger().Error(ctx, result.Error)
			err = errors.New("添加订阅失败")
			return err
		} else {
			return
		}
	}

	if rssResp = component.GetContent(rssLink); rssResp == "" {
		err = errors.New("获取RSS内容失败")
		return err
	} else {
		var (
			goFeed *gofeed.Feed
			feedID string
		)
		goFeed, err = lib.ParseRSSFeed(rssResp)
		if err != nil {
			component.Logger().Error(ctx, err)
			err = errors.New("解析RSS链接失败")
			return err
		}
		rssFeedChannelMode.Title = goFeed.Title
		if goFeed.Image != nil {
			rssFeedChannelMode.ImageUrl = goFeed.Image.URL
		}
		rssFeedChannelMode.ChannelDesc = goFeed.Description
		rssFeedChannelMode.Link = goFeed.Link
		rssFeedChannelMode.RssLink = rssLink
		feedID = strconv.FormatUint(ghash.RS64([]byte(rssFeedChannelMode.Link+rssFeedChannelMode.Title)), 32)
		rssFeedChannelMode.Id = feedID

		userSubChannel.ChannelId = feedID
		userSubChannel.UserId = userID
		userSubChannel.Status = 1
		userSubChannel.InputTime = gtime.Now()
	}

	err = component.GetDatabase().Transaction(func(tx *gorm.DB) error {
		var (
			tranErr error
		)

		tranErr = tx.Create(&rssFeedChannelMode).Error
		if tranErr != nil {
			return tranErr
		}

		tranErr = tx.Create(&userSubChannel).Error
		if tranErr != nil {
			return tranErr
		}

		return tranErr
	})
	if err != nil {
		component.Logger().Error(ctx, "insert rss feed data failed : ", err)
	}
	return
}
