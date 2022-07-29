package topic

import (
	"gohub-api/app/models"
	"gohub-api/pkg/app"
	"gohub-api/pkg/database"
	"gohub-api/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (topic Topic) {
	database.DB.Where("id = ? AND status != ?", idstr, models.TopicStatusDeleted).First(&topic)
	return
}

func GetBy(field, value string) (topic Topic) {
	database.DB.Where("? = ? AND status != ?", field, value, models.TopicStatusDeleted).First(&topic)
	return
}

func All() (topics []Topic) {
	database.DB.Where("status!=?", models.TopicStatusDeleted).Find(&topics)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Topic{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (topics []Topic, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Topic{}),
		&topics,
		app.V1URL(database.TableName(&Topic{})),
		perPage,
	)
	return
}
