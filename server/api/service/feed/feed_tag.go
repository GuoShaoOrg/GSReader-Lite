package feed

import (
	"context"
	"gs-reader-lite/server/component"
	"gs-reader-lite/server/model/biz"
)

func GetFeedTag(ctx context.Context, start, size int) (tagList []biz.RespFeedTagData) {

	if err := component.GetDatabase().Table("rss_feed_tag").
		Select("name, count(channel_id) as count ").
		Group("name").
		Order("count(channel_id) desc").
		Limit(size).
		Offset(start).
		Find(&tagList); err != nil {
		component.Logger().Error(ctx, err)
	}

	return
}
