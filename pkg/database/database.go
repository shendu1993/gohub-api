package database

import (
	"database/sql"
	"fmt"

	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

//DB object
var DB *gorm.DB
var SQLDB *sql.DB

//Connect db
func Connect(dbConfig gorm.Dialector, _logger gormLogger.Interface) {
	var err error
	DB, err = gorm.Open(dbConfig, &gorm.Config{
		Logger: _logger,
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	//获取底层的sqlDB
	SQLDB, err = DB.DB()
	if err != nil {
		fmt.Println(err.Error())
	}
}
