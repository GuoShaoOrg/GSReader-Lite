package feed

import (
	"context"
	"gs-reader-lite/server/component"
	"gs-reader-lite/server/model"
	"gs-reader-lite/server/model/biz"
	"strconv"

	"github.com/gogf/gf/v2/encoding/ghash"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gorilla/feeds"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func AddFeedChannelAndItem(ctx context.Context, rssLink string, feed feeds.Feed) error {
	feedItemModeList := make([]model.RssFeedItem, 0)
	feedItemFTSModeList := make([]model.RssFeedItemFTS, 0)

	feedID := strconv.FormatUint(ghash.RS64([]byte(feed.Link.Href+feed.Title)), 32)
	feedChannelModel := model.RssFeedChannel{
		Id:          feedID,
		Title:       feed.Title,
		ChannelDesc: feed.Description,
		RssLink:     rssLink,
	}

	if feed.Image != nil {
		feedChannelModel.ImageUrl = feed.Image.Url
	}

	if feed.Link != nil {
		feedChannelModel.Link = feed.Link.Href
	}

	for _, item := range feed.Items {
		feedItem := model.RssFeedItem{
			ChannelId:   feedID,
			Title:       item.Title,
			Description: item.Description,
			Content:     item.Content,
			Date:        gtime.New(item.Created.String()),
			InputDate:   gtime.Now(),
		}
		if item.Link != nil {
			feedItem.Link = item.Link.Href
		}
		if item.Author != nil {
			feedItem.Author = item.Author.Name
		}
		uniString := feedItem.Link + feedItem.Title
		feedItemID := strconv.FormatUint(ghash.RS64([]byte(uniString)), 32)
		feedItem.Id = feedItemID
		feedItemModeList = append(feedItemModeList, feedItem)

		feedItemFTS := model.RssFeedItemFTS{
			ChannelId:   feedID,
			Title:       item.Title,
			Description: item.Description,
			Content:     item.Content,
			Date:        feedItem.InputDate,
			InputDate:   feedItem.InputDate,
		}
		if item.Link != nil {
			feedItemFTS.Link = item.Link.Href
		}
		if item.Author != nil {
			feedItemFTS.Author = item.Author.Name
		}
		feedItemFTS.Id = feedItemID
		// TODO Workaround for https://github.com/yanyiwu/gojieba/issues/81
		// jieba := lib.GetJieBa()
		// titleSPArr := jieba.CutForSearch(feedItemFTS.Title, true)
		// feedItemFTS.TitleSP = strings.Join(titleSPArr, " ")

		// contentSPArr := jieba.CutForSearch(feedItemFTS.Content, true)
		// feedItemFTS.ContentSP = strings.Join(contentSPArr, " ")

		// descriptionSPArr := jieba.CutForSearch(feedItemFTS.Description, true)
		// feedItemFTS.DescriptionSP = strings.Join(descriptionSPArr, " ")
		// TODO Workaround end

		feedItemFTSModeList = append(feedItemFTSModeList, feedItemFTS)
	}

	err := component.GetDatabase().Transaction(func(tx *gorm.DB) error {
		var (
			tranErr error
		)

		_ = tx.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&feedChannelModel).Error

		tranErr = tx.Create(&feedItemModeList).Error
		if tranErr != nil {
			return tranErr
		}

		tranErr = tx.Create(&feedItemFTSModeList).Error
		if tranErr != nil {
			return tranErr
		}

		return tranErr
	})

	if err != nil {
		component.Logger().Error(ctx, "insert rss feed item data failed : ", err)
	}

	return err
}

func GetFeedItemByChannelId(ctx context.Context, start, size int, channelId, userId string) (itemList []biz.RssFeedItemData) {

	if userId != "" {
		if err := component.GetDatabase().Table("rss_feed_item rfi").
			Select(model.RFIWithoutContentFieldSql+", rfc.rss_link as rssLink, rfc.title as channelTitle, rfc.image_url as channelImageUrl, umfi.status as marked, usc.status as sub").
			Joins("left join rss_feed_channel rfc on rfi.channel_id=rfc.id").
			Joins("left join user_sub_channel usc on usc.channel_id=rfi.channel_id and usc.user_id="+"'"+userId+"'").
			Joins("left join user_mark_feed_item umfi on umfi.channel_item_id=rfi.id and umfi.user_id="+"'"+userId+"'").
			Where("rfi.channel_id", channelId).
			Order("rfi.input_date desc").
			Limit(size).
			Offset(start).
			Scan(&itemList).Error; err != nil {
			component.Logger().Error(ctx, err)
		}
	} else {
		if err := component.GetDatabase().Table("rss_feed_item rfi").
			Select(model.RFIWithoutContentFieldSql+", rfc.rss_link as rssLink, rfc.title as channelTitle, rfc.image_url as channelImageUrl").
			Joins("left join rss_feed_channel rfc on rfi.channel_id=rfc.id").
			Where("rfi.channel_id", channelId).
			Order("rfi.input_date desc").
			Limit(size).
			Offset(start).
			Find(&itemList).Error; err != nil {
			component.Logger().Error(ctx, err)
		}
	}

	return
}

func GetFeedItemListByUserId(ctx context.Context, userId string, start, size int) (itemList []biz.RssFeedItemData) {
	if err := component.GetDatabase().Table("rss_feed_item rfi").
		Select(model.RFIWithoutContentFieldSql+", rfc.rss_link as rssLink, rfc.title as channelTitle, rfc.image_url as channelImageUrl, umfi.status as marked, usc.status as sub").
		Joins("inner join user_sub_channel usc on usc.channel_id=rfi.channel_id").
		Joins("left join user_mark_feed_item umfi on umfi.channel_item_id=rfi.id").
		Joins("inner join rss_feed_channel rfc on usc.channel_id=rfc.id").
		Where("usc.user_id = ? and usc.status = 1", userId).
		Order("rfi.input_date desc").
		Limit(size).
		Offset(start).
		Find(&itemList).Error; err != nil {
		component.Logger().Error(ctx, err)
	}

	return
}

func GetLatestFeedItem(ctx context.Context, userId string, start, size int) (itemList []biz.RssFeedItemData) {
	if userId != "" {
		if err := component.GetDatabase().Table("rss_feed_item rfi").
			Select(model.RFIWithoutContentFieldSql + ", rfc.rss_link as rssLink, rfc.title as channelTitle, rfc.image_url as channelImageUrl, umfi.status as marked").
			Joins("inner join rss_feed_channel rfc on rfi.channel_id=rfc.id").
			Joins("left join user_mark_feed_item umfi on umfi.channel_item_id=rfi.id and umfi.user_id=" + "'" + userId + "'").
			Group("rfc.id").
			Order("rfi.input_date desc").
			Limit(size).
			Offset(start).
			Find(&itemList).Error; err != nil {
			component.Logger().Error(ctx, err)
		}
	} else {
		if err := component.GetDatabase().Table("rss_feed_item rfi").
			Select(model.RFIWithoutContentFieldSql + ", rfc.rss_link as rssLink, rfc.title as channelTitle, rfc.image_url as channelImageUrl").
			Joins("inner join rss_feed_channel rfc on rfi.channel_id=rfc.id").
			Group("rfc.id").
			Order("rfi.input_date desc").
			Limit(size).
			Offset(start).
			Find(&itemList).Error; err != nil {
			component.Logger().Error(ctx, err)
		}
	}

	return
}

func GetRandomFeedItem(ctx context.Context, start, size int, userId string) (itemList []biz.RssFeedItemData) {
	threeDayBefore := gtime.Now().AddDate(0, 0, -3).Format("Y-m-d")
	if userId != "" {
		if err := component.GetDatabase().Table("rss_feed_item rfi").
			Select(model.RFIWithoutContentFieldSql+", rfc.rss_link as rssLink, rfc.title as channelTitle, rfc.image_url as channelImageUrl, umfi.status as marked, usc.status as sub").
			Joins("inner join rss_feed_channel rfc on rfi.channel_id=rfc.id").
			Joins("left join user_sub_channel usc on usc.channel_id=rfi.channel_id and usc.user_id="+"'"+userId+"'").
			Joins("left join user_mark_feed_item umfi on umfi.channel_item_id=rfi.id and umfi.user_id="+"'"+userId+"'").
			Order("RAND()").
			Where("rfi.input_date>=?", threeDayBefore).
			Limit(size).
			Offset(start).
			Find(&itemList).Error; err != nil {
			component.Logger().Error(ctx, err)
		}
	} else {
		if err := component.GetDatabase().Table("rss_feed_item rfi").
			Select(model.RFIWithoutContentFieldSql+", rfc.rss_link as rssLink, rfc.title as channelTitle, rfc.image_url as channelImageUrl").
			Joins("inner join rss_feed_channel rfc on rfi.channel_id=rfc.id").
			Order("RAND()").
			Where("rfi.input_date>=?", threeDayBefore).
			Limit(size).
			Offset(start).
			Find(&itemList).Error; err != nil {
			component.Logger().Error(ctx, err)
		}
	}

	return
}

func MarkFeedItem(ctx context.Context, userId, feedItemId string) (status int, err error) {
	markItem := model.UserMarkFeedItem{}
	result := component.GetDatabase().Table("user_mark_feed_item").Where("user_id = ? and channel_item_id = ?", userId, feedItemId).Find(&markItem)

	if err == nil && markItem.ChannelItemId != "" {
		if markItem.Status == 1 {
			markItem.Status = -1
		} else {
			markItem.Status = 1
		}
		err = component.GetDatabase().Table("user_mark_feed_item").Where("user_id = ? and channel_item_id = ?", userId, feedItemId).Updates(model.UserMarkFeedItem{Status: markItem.Status}).Error
		return markItem.Status, err
	}

	nowTime := gtime.Now().Format("Y-m-d H:i:s.u")
	markItem = model.UserMarkFeedItem{
		UserId:        userId,
		ChannelItemId: feedItemId,
		InputTime:     gtime.NewFromStr(nowTime),
		Status:        1,
	}

	result = component.GetDatabase().Create(&markItem)

	return markItem.Status, result.Error
}

func GetMarkedFeedItemListByUserId(ctx context.Context, userId string, start, size int) (itemList []biz.RssFeedItemData) {
	if err := component.GetDatabase().Table("rss_feed_item rfi").
		Select(model.RFIWithoutContentFieldSql + ", rfc.rss_link as rssLink, rfc.title as channelTitle, rfc.image_url as channelImageUrl, umfi.status as marked, usc.status as sub").
		Joins("inner join user_mark_feed_item umfi on umfi.channel_item_id=rfi.id and umfi.user_id=" + "'" + userId + "'").
		Joins("left join user_sub_channel usc on usc.channel_id=rfi.channel_id and usc.user_id=" + "'" + userId + "'").
		Joins("inner join rss_feed_channel rfc on rfi.channel_id=rfc.id").
		Order("umfi.input_time desc").
		Where("umfi.status = 1").
		Limit(size).
		Offset(start).
		Find(&itemList).Error; err != nil {
		component.Logger().Error(ctx, err)
	}

	return
}

func GetFeedItemByItemId(ctx context.Context, itemId, userId string) (item biz.RssFeedItemData) {

	if userId != "" {
		if err := component.GetDatabase().Table("rss_feed_item rfi").
			Joins("left join rss_feed_channel rfc on rfi.channel_id=rfc.id").
			Joins("left join user_sub_channel usc on usc.channel_id=rfi.channel_id and usc.user_id="+"'"+userId+"'").
			Joins("left join user_mark_feed_item umfi on umfi.channel_item_id=rfi.id and umfi.user_id="+"'"+userId+"'").
			Select("rfi.*, rfc.rss_link as rssLink, rfc.title as channelTitle, rfc.image_url as channelImageUrl, umfi.status as marked, usc.status as sub").
			Where("rfi.id", itemId).
			Find(&item).Error; err != nil {
			component.Logger().Error(ctx, err)
		}
	} else {
		if err := component.GetDatabase().Table("rss_feed_item rfi").
			Joins("left join rss_feed_channel rfc on rfi.channel_id=rfc.id").
			Select("rfi.*, rfc.rss_link as rssLink, rfc.title as channelTitle, rfc.image_url as channelImageUrl").
			Where("rfi.id", itemId).
			Find(&item).Error; err != nil {
			component.Logger().Error(ctx, err)
		}
	}

	return
}

func SearchFeedItem(ctx context.Context, userId, keyword string, start, size int) (items []biz.RssFeedItemData) {

	if size == 0 {
		size = 10
	}

	// TODO Workaround for https://github.com/yanyiwu/gojieba/issues/81
	// keywordArray := lib.GetJieBa().CutForSearch(keyword, true)
	// queryKeyword := strings.Join(keywordArray, " ")
	// queryString := "(SELECT id FROM rss_feed_item_fts fts WHERE description_sp MATCH ? OR content_sp MATCH ? OR title_sp MATCH ? LIMIT ? OFFSET ?)"
	// TODO Workaround end
	queryString := "(SELECT id FROM rss_feed_item_fts fts WHERE description_sp MATCH ? OR content_sp MATCH ? OR title_sp MATCH ? LIMIT ? OFFSET ?)"

	if err := component.GetDatabase().Table("rss_feed_item rfi").
		Select(model.RFIWithoutContentFieldSql+", rfc.rss_link as rssLink, rfc.title as channelTitle, rfc.image_url as channelImageUrl, umfi.status as marked, usc.status as sub").
		Joins("inner join user_sub_channel usc on usc.channel_id=rfi.channel_id").
		Joins("left join user_mark_feed_item umfi on umfi.channel_item_id=rfi.id").
		Joins("inner join rss_feed_channel rfc on usc.channel_id=rfc.id").
		Where("usc.user_id = ?", userId).
		Where("rfi.id in "+queryString, keyword, keyword, keyword, size, start).
		Order("rfi.input_date desc").
		Limit(size).
		Offset(start).
		Find(&items).Error; err != nil {
		component.Logger().Error(ctx, err)
	}
	return
}
