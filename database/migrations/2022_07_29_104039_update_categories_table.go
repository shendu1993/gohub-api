package migrations

import (
	"database/sql"
	"gohub-api/app/models"
	"gohub-api/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type Category struct {
		models.BaseModel
		Name        string `gorm:"type:varchar(255);not null;index"`
		Description string `gorm:"type:varchar(255);default:null"`
		Status      string `gorm:"type:tinyint(1);default:1"`
		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Category{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Category{})
	}

	migrate.Add("2022_07_29_104039_update_categories_table", up, down)
}
