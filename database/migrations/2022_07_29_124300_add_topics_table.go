package migrations

import (
	"database/sql"
	"gohub-api/app/models"
	"gohub-api/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type User struct {
		models.BaseModel
	}
	type Category struct {
		models.BaseModel
	}
	type Topic struct {
		models.BaseModel
		Title      string `gorm:"type:varchar(255);not null;index"`
		Body       string `gorm:"type:longtext;not null"`
		UserID     string `gorm:"type:bigint;not null;index"`
		CategoryID string `gorm:"type:bigint;not null;index"`
		Status     int    `gorm:"type:tinyint(1);not null; default:1"`
		// 会创建 user_id 和 category_id 外键的约束
		User     User
		Category Category
		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Topic{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Topic{})
	}

	migrate.Add("2022_07_29_124300_add_topics_table", up, down)
}
