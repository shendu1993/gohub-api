package migrations

import (
	"database/sql"
	"gohub-api/app/models"
	"gohub-api/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type Link struct {
		models.BaseModel
		Name   string `gorm:"type:varchar(255);not null"`
		URL    string `gorm:"type:varchar(255);default:null"`
		Status int    `gorm:"type:tinyint(1);default:1"`
		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Link{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Link{})
	}

	migrate.Add("2022_08_01_173309_add_links_table", up, down)
}
