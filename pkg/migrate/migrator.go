package migrate

import (
	"gohub-api/pkg/console"
	"gohub-api/pkg/database"
	"gohub-api/pkg/file"
	"io/ioutil"

	"gorm.io/gorm"
)

//Migrator 数据迁移操作类
type Migrator struct {
	Folder   string
	DB       *gorm.DB
	Migrator gorm.Migrator
}

//Migration 对应数据的migrations表里面的一条数据
type Migration struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement;"`
	Migration string `gorm:"type:varchar(255); not null;unique"`
	Batch     int
}

//NewMigrator 船舰 Migrator实例 ，用以执行迁移操作
func NewMigrator() *Migrator {
	//初始化必要属性
	migrator := &Migrator{
		Folder:   "database/migrations",
		DB:       database.DB,
		Migrator: database.DB.Migrator(),
	}
	//migrations 不存在的话就创建它
	migrator.createMigrationsTable()
	return migrator
}

//创建migrations 表
func (migrator *Migrator) createMigrationsTable() {
	migration := Migration{}
	//不存在才创建
	if !migrator.Migrator.HasTable(&migration) {
		err := migrator.Migrator.CreateTable(&migration)
		if err != nil {
			return
		}
	}
}

//从文件目录读取文件，去v奥正确的时间顺序
func (migrator *Migrator) readAllMigrationFiles() []MigrationFile {
	//读取 databases/migrations/ 目录下的所有文件
	//默认会按照文件名称进行排序
	files, err := ioutil.ReadDir(migrator.Folder)
	console.ExitIf(err)
	var migrateFiles []MigrationFile
	for _, f := range files {
		//去除文件的后缀 .go
		fileName := file.FileNameWithoutExtension(f.Name())
		//通过迁移文件的抿成获取【MigrationFile】对象
		mfile := getMigrationFile(fileName)

		//加个判断，确保迁移文件可用，再放进migrateFiles数组中
		if len(mfile.FileName) > 0 {
			migrationFiles = append(migrateFiles, mfile)
		}
	}
	//返回排序好的【migrateFiles】数组中
	return migrateFiles
}

//getBatch 获取当前这个批次的值
func (migrator *Migrator) getBatch() int {
	//batch 默认为1
	batch := 1
	//获取最后一条的迁移数据
	lastMigration := Migration{}
	migrator.DB.Order("id DESC").First(&lastMigration)
	//如果有值的话，加一
	if lastMigration.ID > 0 {
		batch = lastMigration.Batch + 1
	}
	return batch
}

//Up之宗所有未迁移过的表
func (migrator *Migrator) Up() {
	//读取所有迁移文件，确保上按照时间顺序排序
	migrator.readAllMigrationFiles()
	//获取当前批次的值
	batch := migrator.getBatch()

	//获取所有迁移数据
	migrations := []Migration{}
	migrator.DB.Find(&migrations)

	//可以通过此值来判断数据库是否已经是最新
	runed := false

	//对迁移文件进行遍历，如果灭有执行过，怎执行up 回调

	for _, mfile := range migrationFiles {
		//对比文件名，看是否已经运行过
		if mfile.isNotMigrated(migrations) {
			migrator.runUpMigration(mfile, batch)
			runed = true
		}
	}

	if !runed {
		console.Success("database is up to date.")
	}

}

//执行迁移，执行迁移up方法
func (migrator *Migrator) runUpMigration(mfile MigrationFile, batch int) {
	//执行up区块的SQL
	if mfile.Up != nil {
		//友好提示
		console.Warning("migrating " + mfile.FileName)
		//执行up方法
		mfile.Up(database.DB.Migrator(), database.SQLDB)
		//提示已经迁移了哪个文件
		console.Success("migrated " + mfile.FileName)
	}
	//入库
	err := migrator.DB.Create(&Migration{Migration: mfile.FileName, Batch: batch}).Error
	console.ExitIf(err)
}
