package lib

import "github.com/mmcdole/gofeed"

func ParseRSSFeed(feedContent string) (feed *gofeed.Feed, err error) {
	feed, err = gofeed.NewParser().ParseString(feedContent)
	if err != nil {
		return nil, err
	}
	return feed, nil
	
}