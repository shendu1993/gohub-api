package link

import (
	"gohub-api/pkg/app"
	"gohub-api/pkg/cache"
	"gohub-api/pkg/database"
	"gohub-api/pkg/helpers"
	"gohub-api/pkg/paginator"
	"time"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (link Link) {
	database.DB.Where("id", idstr).First(&link)
	return
}

func GetBy(field, value string) (link Link) {
	database.DB.Where("? = ?", field, value).First(&link)
	return
}

func All() (links []Link) {
	database.DB.Find(&links)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Link{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (links []Link, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Link{}),
		&links,
		app.V1URL(database.TableName(&Link{})),
		perPage,
	)
	return
}

// getByLimit 限制返回数量
func GetByLimit(limit int) (link []Link) {
	database.DB.Order("updated_at Desc").Limit(limit).Find(&link)
	return
}

//links:all
func AllCached() (links []Link) {
	//设置缓存key
	cacheKey := "links:all"
	//设置过期时间
	expireTime := 120 * time.Minute
	//取数据
	cache.GetObject(cacheKey, &links)
	//如果缓存数据为空
	if helpers.Empty(links) {
		//检查数据库
		links = All()
		if helpers.Empty(links) {
			return links
		}
		// 设置缓存
		cache.Set(cacheKey, links, expireTime)
	}
	return
}
