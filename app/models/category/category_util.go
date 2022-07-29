package category

import (
	"gohub-api/app/models"
	"gohub-api/pkg/app"
	"gohub-api/pkg/database"
	"gohub-api/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (category Category) {
	database.DB.Where("id = ? AND status = ?", idstr, models.CategoryStatusNormal).First(&category)
	return
}

func GetBy(field, value string) (category Category) {
	database.DB.Where("? = ? AND status = ?", field, value, models.CategoryStatusNormal).First(&category)
	return
}

func All() (categories []Category) {
	database.DB.Where("status", models.CategoryStatusNormal).Find(&categories)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Category{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (categories []Category, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Category{}),
		&categories,
		app.V1URL(database.TableName(&Category{})),
		perPage,
	)
	return
}
