package factories

import (
	"gohub-api/app/models/category"

	"github.com/bxcodec/faker/v3"
)

func MakeCategories(count int) []category.Category {

	var objs []category.Category

	// 设置唯一性，如 Category 模型的某个字段需要唯一，即可取消注释
	// faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {
		categoryModel := category.Category{
			Name:        faker.ChineseName(),
			Description: faker.ChineseName(),
		}
		objs = append(objs, categoryModel)
	}

	return objs
}