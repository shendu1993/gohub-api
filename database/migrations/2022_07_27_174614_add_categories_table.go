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
        // Name     string `gorm:"type:varchar(255);not null;index"`
        // Put fields in here
        // FIXME()
        models.CommonTimestampsField
    }

    up := func(migrator gorm.Migrator, DB *sql.DB) {
        migrator.AutoMigrate(&Category{})
    }

    down := func(migrator gorm.Migrator, DB *sql.DB) {
        migrator.DropTable(&Category{})
    }

    migrate.Add("2022_07_27_174614_add_categories_table", up, down)
}