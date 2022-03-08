package locale

import (
	"gohub-api/pkg/helpers"

	"github.com/gin-gonic/gin"
)

//翻译 key
func Translate(c *gin.Context, key string) string {
	//header 请求语言
	language := GetAcceptLanguage(c)
	//文件目录
	filePath := helpers.GetAppPath() + "\\locale\\" + language + ".json"
	//把json文件转化为Map
	messageMap := helpers.JSONToMap(filePath)
	//如果翻译为空，则把key 直接返回
	if helpers.Empty(messageMap[key]) {
		return key
	}
	return messageMap[key]
}

//获取Accept-Language
func GetAcceptLanguage(c *gin.Context) string {
	language := c.GetHeader("Accept-Language")
	return language
}
