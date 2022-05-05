package lib

import "github.com/yanyiwu/gojieba"

var jiebaInstance *gojieba.Jieba

func GetJieBa() *gojieba.Jieba {
	if jiebaInstance == nil {
		jiebaInstance = gojieba.NewJieba()
	}
	return jiebaInstance
}