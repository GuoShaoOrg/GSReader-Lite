package jobs

import (
	"context"
	feedsvc "gs-reader-lite/server/api/service/feed"
	"gs-reader-lite/server/component"
	"gs-reader-lite/server/model/biz"
	"time"

	"github.com/gorilla/feeds"
	"github.com/mmcdole/gofeed"
)

func doSync(f func()) {
	go func() {
		var freshStartTime = time.Now()
		var refreshHoldTime = time.Minute * 40
		var ctx = context.Background()
		component.Logger().Info(ctx, "Start Feed sync job")
		for {
			freshStartTime = time.Now()
			f()
			if time.Now().Sub(freshStartTime) < refreshHoldTime {
				time.Sleep(time.Minute * 60)
			}
		}
	}()
}

func RegisterJob() {
	doSync(doNonAsyncRefreshFeed)
}

func doNonAsyncRefreshFeed() {
	var (
		routerLength    int
		ctx             context.Context
		feedChannelList []biz.RssFeedChannelData
	)

	feedChannelList = feedsvc.GetAllFeedChannelList(ctx)
	if len(feedChannelList) > 0 {
		routerLength = len(feedChannelList)
		for index, feedChannelInfo := range feedChannelList {
			var (
				resp      string
				err       error
				feed      *gofeed.Feed
				feedInfo  feeds.Feed
				feedItems []*feeds.Item
			)

			if resp = component.GetContent(feedChannelInfo.RssLink); resp == "" {
				component.Logger().Error(ctx, "Feed refresh cron job error ")
				continue
			}

			fp := gofeed.NewParser()
			feed, err = fp.ParseString(resp)
			if err != nil {
				component.Logger().Error(ctx, "Parse RSS response error : ", err)
				continue
			}

			feedInfo = feeds.Feed{
				Title:       feed.Title,
				Link:        &feeds.Link{Href: feed.Link},
				Description: feed.Description,
			}

			if feed.UpdatedParsed == nil {
				feedInfo.Updated = time.Now()
			} else {
				feedInfo.Updated = *feed.UpdatedParsed
			}

			if feed.PublishedParsed == nil {
				feedInfo.Created = time.Now()
			} else {
				feedInfo.Created = *feed.UpdatedParsed
			}

			if feed.Image != nil {
				feedInfo.Image = &feeds.Image{}
				feedInfo.Image.Title = feed.Image.Title
				feedInfo.Image.Url = feed.Image.URL
				feedInfo.Image.Link = feed.Image.URL
			}

			if len(feed.Authors) > 0 {
				feedInfo.Author = &feeds.Author{}
				feedInfo.Author.Name = feed.Authors[0].Name
				feedInfo.Author.Email = feed.Authors[0].Email
			}

			for _, itemV := range feed.Items {
				item := feeds.Item{
					Title:       itemV.Title,
					Link:        &feeds.Link{Href: itemV.Link},
					Description: itemV.Description,
					Content:     itemV.Content,
					Created:     *itemV.PublishedParsed,
				}

				if itemV.Author != nil {
					item.Author = &feeds.Author{}
					item.Author.Name = itemV.Author.Name
					item.Author.Email = itemV.Author.Email
				}

				if len(itemV.Enclosures) > 0 {
					item.Enclosure = &feeds.Enclosure{}
					item.Enclosure.Url = itemV.Enclosures[0].URL
					item.Enclosure.Type = itemV.Enclosures[0].Type
					item.Enclosure.Length = itemV.Enclosures[0].Length
				}

				feedItems = append(feedItems, &item)
			}

			feedInfo.Items = feedItems

			err = feedsvc.AddFeedChannelAndItem(ctx, feedInfo)
			if err != nil {
				component.Logger().Error(ctx, "Add feed channel and item error : ", err)
				continue
			}
			component.Logger().Infof(ctx, "Processed %d/%d feed refresh\n", (index + 1), routerLength)
		}
	}
}
