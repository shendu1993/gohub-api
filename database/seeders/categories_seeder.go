package seeders

import (
	"fmt"
	"gohub-api/database/factories"
	"gohub-api/pkg/console"
	"gohub-api/pkg/logger"
	"gohub-api/pkg/seed"

	"gorm.io/gorm"
)

func init() {

	seed.Add("SeedCategoriesTable", func(db *gorm.DB) {
		//创建10条数据
		category := factories.MakeCategories(10)
		//批量创建用户（注意批创建不会条用模型钩子）

		result := db.Table("categories").Create(&category)

		//记录错误
		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		// 打印运行情况
		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))

	})

}