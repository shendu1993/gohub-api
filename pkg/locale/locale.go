package locale

import (
	"gohub-api/pkg/config"
	"gohub-api/pkg/helpers"
	"gohub-api/pkg/str"
	"strings"

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
	//如果从Accept-Language获取不到，设置默认语言
	if helpers.Empty(language) {
		return config.GetString("app.language")
	}
	//把支持的语言分割为数组
	SupportLanguageArr := strings.Split(config.GetString("app.support_language"), ",")
	//如果不在支持的语言内，则返回默认语言
	if str.InArray(language, SupportLanguageArr) {
		return language
	}
	return config.GetString("app.language")
}
