package component

import (
	"context"
	"github.com/gogf/gf/v2/net/gclient"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

func GetHttpClient() (client *gclient.Client) {

	client = g.Client()
	client.SetTimeout(time.Second * 10)

	return
}

func GetContent(link string) (resp string) {
	var (
		client *gclient.Client
	)
	ctx := context.Background()
	client = GetHttpClient()
	resp = client.SetHeaderMap(getHeaders()).GetContent(ctx, link)

	return
}

func getHeaders() map[string]string {
	headers := make(map[string]string)
	headers["accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
	headers["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36 Edg/84.0.522.63"
	return headers
}
